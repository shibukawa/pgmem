package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_spgExtractNodeLabels(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = int32(base.Ui32(v7)>>(uint(int32(3))%32)) & int32(_a_F_spgExtractNodeLabels_0)
	v13 = l1 + int32(8)
	v16 = v13 + int32(base.Ui32(v7)>>(uint(int32(16))%32))
	v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+6)))
	if v17 < v3 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L12
	} else {
		goto L26
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L23
	}
L3:
	;
	return v91
L4:
	;
	if v11 == int32(0) {
		v91 = v3
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v40 = F_palloc_mul(m, int32(8), v11)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v24 = v16
	v27 = v3
	goto L8
L8:
	;
	v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+6)))
	if int32(0) <= v28 {
		goto L2
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v35 = v27 + int32(1)
	if v35 != v11 {
		v24 = v24 + v28&int32(_a_F_spgExtractNodeLabels_0)
		v27 = v35
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	return int32(0)
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v44&int32(_a_F_spgExtractNodeLabels_1) == int32(0) {
		v91 = v40
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v55 = v13 + int32(base.Ui32(v44)>>(uint(int32(16))%32))
	v56 = int32(0)
	goto L15
L15:
	;
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+6)))
	if v59 < int32(0) {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v91 = v40
	goto L3
L17:
	;
	v63 = v55 + int32(8)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
	if v67 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40+v56<<(uint(int32(3))%32)))) = v72
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+6)))
	v75 = int32(_a_F_spgExtractNodeLabels_0)
	v79 = v56 + int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v79) < base.Ui32(int32(base.Ui32(v80)>>(uint(int32(3))%32))&v75) {
		v55 = v55 + v74&v75
		v56 = v79
		goto L15
	} else {
		goto L22
	}
L19:
	;
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
	v72 = v70
	goto L18
L20:
	;
	goto L21
L21:
	;
	v72 = base.I64_extend_i32_u(v63)
	goto L18
L22:
	;
	goto L16
L23:
	;
	F_errmsg_internal(m, int32(_a_F_spgExtractNodeLabels_2), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_spgExtractNodeLabels_3), int32(1172), int32(_a_F_spgExtractNodeLabels_4))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	F_errmsg_internal(m, int32(_a_F_spgExtractNodeLabels_2), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_spgExtractNodeLabels_3), int32(1183), int32(_a_F_spgExtractNodeLabels_4))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_spgFormDeadTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v3 = l2
	v4 = l3
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1&int32(3) | int32(64)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
	v14 = v12 & int32(_a_F_spgFormDeadTuple_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)) = uint16(v14)
	if l1 == int32(1) {
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+10)) = uint16(v4)
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+8)) = uint16(v3)
		v21 = int32(base.Ui32(v3) >> (uint(int32(16)) % 32))
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+6)) = uint16(v21)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v23
		return v6
	} else {
		v26 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+10)) = uint16(v26)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+6)) = int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v26
		return v6
	}
}
func F_spg_box_quad_choose(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 float64
	_ = v19
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v23 int32
	_ = v23
	var v26 float64
	_ = v26
	var v27 float64
	_ = v27
	var v29 int32
	_ = v29
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v6)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(1)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+20)))
	if v13 == int32(0) {
		v18 = base.I32_wrap_i64(v9)
		v19 = *(*float64)(unsafe.Add(mBase, uint32(v18)+16))
		v20 = base.I32_wrap_i64(v7)
		v21 = *(*float64)(unsafe.Add(mBase, uint32(v20)+16))
		if base.F64_gt(v19, v21) != 0 {
			v23 = int32(8)
		} else {
			v23 = int32(0)
		}
		v26 = *(*float64)(unsafe.Add(mBase, uint32(v18)))
		v27 = *(*float64)(unsafe.Add(mBase, uint32(v20)))
		if base.F64_gt(v26, v27) != 0 {
			v29 = v23 | int32(4)
		} else {
			v29 = v23
		}
		v32 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
		v33 = *(*float64)(unsafe.Add(mBase, uint32(v20)+24))
		if base.F64_gt(v32, v33) != 0 {
			v35 = v29 | int32(2)
		} else {
			v35 = v29
		}
		v36 = *(*float64)(unsafe.Add(mBase, uint32(v18)+8))
		v37 = *(*float64)(unsafe.Add(mBase, uint32(v20)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v35 | base.F64_gt(v36, v37)
	} else {
	}
	return int64(0)
}
func F_spg_box_quad_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 float64
	_ = v66
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v75 float64
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 float64
	_ = v127
	var v130 float64
	_ = v130
	var v133 float64
	_ = v133
	var v136 float64
	_ = v136
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int64
	_ = v180
	var v181 int32
	_ = v181
	var v182 float64
	_ = v182
	var v183 float64
	_ = v183
	var v184 float64
	_ = v184
	var v185 float64
	_ = v185
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v188 float64
	_ = v188
	var v189 float64
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v23 = F_palloc_mul(m, int32(8), v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int64(0)
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		v29 = F_palloc_mul(m, int32(8), v28)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int64(0)
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v33 = F_palloc_mul(m, int32(8), v32)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int64(0)
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v37 = F_palloc_mul(m, int32(8), v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int64(0)
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
					if int32(0) < v39 {
						v50 = int32(0)
						for {
							v61 = v50 << (uint(int32(3)) % 32)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v63+v61)))
							v66 = *(*float64)(unsafe.Add(mBase, uint32(v65)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v23+v61))) = v66
							v69 = *(*float64)(unsafe.Add(mBase, uint32(v65)))
							*(*float64)(unsafe.Add(mBase, uint32(v61+v29))) = v69
							v72 = *(*float64)(unsafe.Add(mBase, uint32(v65)+24))
							*(*float64)(unsafe.Add(mBase, uint32(v61+v33))) = v72
							v75 = *(*float64)(unsafe.Add(mBase, uint32(v65)+8))
							*(*float64)(unsafe.Add(mBase, uint32(v61+v37))) = v75
							v78 = v50 + int32(1)
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							if v78 < v79 {
								v50 = v78
								continue
							} else {
								break
							}
							break
						}
						v81 = v79
					} else {
						v81 = v39
					}
					F_pg_qsort(m, v23, v81, int32(8), int32(1442))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int64(0)
					} else {
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
						F_pg_qsort(m, v29, v103, int32(8), int32(1442))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int64(0)
						} else {
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							F_pg_qsort(m, v33, v108, int32(8), int32(1442))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int64(0)
							} else {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
								F_pg_qsort(m, v37, v113, int32(8), int32(1442))
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return int64(0)
								} else {
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
									v120 = F_palloc(m, int32(32))
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int64(0)
									} else {
										v123 = base.I32_div_s(v118, int32(2))
										v125 = v123 << (uint(int32(3)) % 32)
										v127 = *(*float64)(unsafe.Add(mBase, uint32(v23+v125)))
										*(*float64)(unsafe.Add(mBase, uint32(v120)+16)) = v127
										v130 = *(*float64)(unsafe.Add(mBase, uint32(v29+v125)))
										*(*float64)(unsafe.Add(mBase, uint32(v120))) = v130
										v133 = *(*float64)(unsafe.Add(mBase, uint32(v125+v33)))
										*(*float64)(unsafe.Add(mBase, uint32(v120)+24)) = v133
										v136 = *(*float64)(unsafe.Add(mBase, uint32(v125+v37)))
										*(*float64)(unsafe.Add(mBase, uint32(v120)+8)) = v136
										*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = int64(16)
										*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = base.I64_extend_i32_u(v120)
										v142 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v142)
										v145 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
										v146 = F_palloc_mul(m, int32(4), v145)
										mBase = m.M
										v147 = m.ExcPending
										if v147 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v146
											v150 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
											v151 = F_palloc_mul(m, int32(8), v150)
											mBase = m.M
											v152 = m.ExcPending
											if v152 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v151
												v154 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
												if int32(0) < v154 {
													v162 = int32(0)
													for {
														v177 = v162 << (uint(int32(3)) % 32)
														v178 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
														v180 = *(*int64)(unsafe.Add(mBase, uint32(v177+v178)))
														v181 = base.I32_wrap_i64(v180)
														v182 = *(*float64)(unsafe.Add(mBase, uint32(v181)+8))
														v183 = *(*float64)(unsafe.Add(mBase, uint32(v181)+24))
														v184 = *(*float64)(unsafe.Add(mBase, uint32(v181)+16))
														v185 = *(*float64)(unsafe.Add(mBase, uint32(v181)))
														v186 = *(*float64)(unsafe.Add(mBase, uint32(v120)+8))
														v187 = *(*float64)(unsafe.Add(mBase, uint32(v120)+24))
														v188 = *(*float64)(unsafe.Add(mBase, uint32(v120)+16))
														v189 = *(*float64)(unsafe.Add(mBase, uint32(v120)))
														v190 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
														*(*int64)(unsafe.Add(mBase, uint32(v190+v177))) = v180 & int64(4294967295)
														v195 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
														if base.F64_gt(v184, v188) != 0 {
															v203 = int32(8)
														} else {
															v203 = int32(0)
														}
														if base.F64_gt(v185, v189) != 0 {
															v207 = v203 | int32(4)
														} else {
															v207 = v203
														}
														if base.F64_gt(v183, v187) != 0 {
															v211 = v207 | int32(2)
														} else {
															v211 = v207
														}
														*(*int32)(unsafe.Add(mBase, uint32(v195+v162<<(uint(int32(2))%32)))) = base.F64_gt(v182, v186) | v211
														v215 = v162 + int32(1)
														v216 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
														if v215 < v216 {
															v162 = v215
															continue
														} else {
															break
														}
														break
													}
												} else {
												}
												return int64(0)
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
func F_spg_desc(m *base.Module, l0 int32, l1 int32) {
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
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
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
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	v8 = m.G0
	v10 = v8 - int32(160)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+48)))
	switch int32(base.Ui32(v14-int32(16)) >> (uint(int32(4)) % 32)) {
	case 0:
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+6)))
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v20
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = v19
		F_appendStringInfo(m, l0, int32(_a_F_spg_desc_0), v10)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v30 == int32(1) {
				F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_1))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
					if v36 != int32(1) {
						m.G0 = v10 + int32(160)
						return
					} else {
						F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_2))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							m.G0 = v10 + int32(160)
							return
						}
					}
				}
			} else {
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
				if v36 != int32(1) {
					m.G0 = v10 + int32(160)
					return
				} else {
					F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_2))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						m.G0 = v10 + int32(160)
						return
					}
				}
			}
		}
	case 1:
		v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
		v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+6)))
		v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v44
		*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v43
		*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v42
		F_appendStringInfo(m, l0, int32(_a_F_spg_desc_3), v10+int32(16))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return
		} else {
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
			if v53 == int32(1) {
				F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_1))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
					if v59 == int32(1) {
						F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_4))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)))
							if v65 != int32(1) {
								m.G0 = v10 + int32(160)
								return
							} else {
								F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_2))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									m.G0 = v10 + int32(160)
									return
								}
							}
						}
					} else {
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)))
						if v65 != int32(1) {
							m.G0 = v10 + int32(160)
							return
						} else {
							F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_2))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								m.G0 = v10 + int32(160)
								return
							}
						}
					}
				}
			} else {
				v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
				if v59 == int32(1) {
					F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_4))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)))
						if v65 != int32(1) {
							m.G0 = v10 + int32(160)
							return
						} else {
							F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_2))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								m.G0 = v10 + int32(160)
								return
							}
						}
					}
				} else {
					v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)))
					if v65 != int32(1) {
						m.G0 = v10 + int32(160)
						return
					} else {
						F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_2))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							m.G0 = v10 + int32(160)
							return
						}
					}
				}
			}
		}
	case 2:
		v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
		v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
		v73 = int32(*(*int8)(unsafe.Add(mBase, uint32(v13)+5)))
		v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+6)))
		v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v75
		*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v74
		*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v73
		*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v72
		*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v71
		F_appendStringInfo(m, l0, int32(_a_F_spg_desc_5), v10+int32(32))
		mBase = m.M
		v85 = m.ExcPending
		if v85 != 0 {
			return
		} else {
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)))
			if v86 != int32(1) {
				m.G0 = v10 + int32(160)
				return
			} else {
				F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_1))
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return
				} else {
					m.G0 = v10 + int32(160)
					return
				}
			}
		}
	case 3:
		v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
		v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v93
		*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v92
		F_appendStringInfo(m, l0, int32(_a_F_spg_desc_6), v10-int32(-64))
		mBase = m.M
		v100 = m.ExcPending
		if v100 != 0 {
			return
		} else {
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)))
			if v101 == int32(1) {
				F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_1))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return
				} else {
					v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+5)))
					if v107 != int32(1) {
						m.G0 = v10 + int32(160)
						return
					} else {
						F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_7))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return
						} else {
							m.G0 = v10 + int32(160)
							return
						}
					}
				}
			} else {
				v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+5)))
				if v107 != int32(1) {
					m.G0 = v10 + int32(160)
					return
				} else {
					F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_7))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return
					} else {
						m.G0 = v10 + int32(160)
						return
					}
				}
			}
		}
	case 4:
		v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
		v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
		v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)))
		v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+14)))
		v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+16)))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v117
		*(*int32)(unsafe.Add(mBase, uint32(v10)+92)) = v116
		*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = v115
		*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v114
		*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v113
		F_appendStringInfo(m, l0, int32(_a_F_spg_desc_8), v10+int32(80))
		mBase = m.M
		v127 = m.ExcPending
		if v127 != 0 {
			return
		} else {
			v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)))
			if v128 == int32(1) {
				F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_9))
				mBase = m.M
				v133 = m.ExcPending
				if v133 != 0 {
					return
				} else {
					v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
					if v134 == int32(1) {
						F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_2))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
							if v140 != int32(1) {
								m.G0 = v10 + int32(160)
								return
							} else {
								F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_10))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return
								} else {
									m.G0 = v10 + int32(160)
									return
								}
							}
						}
					} else {
						v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
						if v140 != int32(1) {
							m.G0 = v10 + int32(160)
							return
						} else {
							F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_10))
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return
							} else {
								m.G0 = v10 + int32(160)
								return
							}
						}
					}
				}
			} else {
				v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
				if v134 == int32(1) {
					F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_2))
					mBase = m.M
					v139 = m.ExcPending
					if v139 != 0 {
						return
					} else {
						v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
						if v140 != int32(1) {
							m.G0 = v10 + int32(160)
							return
						} else {
							F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_10))
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return
							} else {
								m.G0 = v10 + int32(160)
								return
							}
						}
					}
				} else {
					v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
					if v140 != int32(1) {
						m.G0 = v10 + int32(160)
						return
					} else {
						F_appendStringInfoString(m, l0, int32(_a_F_spg_desc_10))
						mBase = m.M
						v145 = m.ExcPending
						if v145 != 0 {
							return
						} else {
							m.G0 = v10 + int32(160)
							return
						}
					}
				}
			}
		}
	case 5:
		v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
		v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
		v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
		v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+6)))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+124)) = v149
		*(*int32)(unsafe.Add(mBase, uint32(v10)+120)) = v148
		*(*int32)(unsafe.Add(mBase, uint32(v10)+116)) = v147
		*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v146
		F_appendStringInfo(m, l0, int32(_a_F_spg_desc_11), v10+int32(112))
		mBase = m.M
		v158 = m.ExcPending
		if v158 != 0 {
			return
		} else {
			m.G0 = v10 + int32(160)
			return
		}
	case 6:
		v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+128)) = v159
		F_appendStringInfo(m, l0, int32(_a_F_spg_desc_12), v10+int32(128))
		mBase = m.M
		v165 = m.ExcPending
		if v165 != 0 {
			return
		} else {
			m.G0 = v10 + int32(160)
			return
		}
	case 7:
		v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)))
		v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
		v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
		v169 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+152)) = v169
		*(*int32)(unsafe.Add(mBase, uint32(v10)+148)) = v168
		*(*int32)(unsafe.Add(mBase, uint32(v10)+144)) = v167
		if v166 != 0 {
			v175 = int32(84)
		} else {
			v175 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+156)) = v175
		F_appendStringInfo(m, l0, int32(_a_F_spg_desc_13), v10+int32(144))
		mBase = m.M
		v181 = m.ExcPending
		if v181 != 0 {
			return
		} else {
			m.G0 = v10 + int32(160)
			return
		}
	default:
		m.G0 = v10 + int32(160)
		return
	}
}
func F_spg_quad_inner_consistent(m *base.Module, l0 int32) int64 {
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
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v122 int64
	_ = v122
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int64
	_ = v143
	var v145 int64
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v159 int64
	_ = v159
	var v160 int32
	_ = v160
	var v166 int64
	_ = v166
	var v167 int32
	_ = v167
	var v173 int64
	_ = v173
	var v174 int32
	_ = v174
	var v180 int64
	_ = v180
	var v181 int32
	_ = v181
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 float64
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v198 int64
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 float64
	_ = v202
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int64
	_ = v333
	var v335 int64
	_ = v335
	var v337 int64
	_ = v337
	var v339 int64
	_ = v339
	var v341 float64
	_ = v341
	var v343 float64
	_ = v343
	var v345 float64
	_ = v345
	var v347 float64
	_ = v347
	var v349 int64
	_ = v349
	var v351 int64
	_ = v351
	var v353 int64
	_ = v353
	var v355 int64
	_ = v355
	var v357 float64
	_ = v357
	var v359 float64
	_ = v359
	var v361 float64
	_ = v361
	var v363 float64
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)+40))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v24 <= v2 {
		v53 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+37)))
	if v54 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v29 = F_palloc_mul(m, int32(4), v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int64(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v29
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v36 = F_palloc_mul(m, int32(4), v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v36
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	if v39 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v42 = int64(-4503599627370496)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v42
	v46 = int64(9218868437227405312)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v46
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v46
	v53 = v16 + int32(-32)
	goto L1
L7:
	;
	goto L8
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v53 = v52
	goto L1
L9:
	;
	m.G0 = v18 - int32(-64)
	return int64(0)
L10:
	;
	v270 = int32(4)
	v272 = F_palloc_mul(m, v270, v270)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L3
	} else {
		goto L57
	}
L11:
	;
	v122 = v22 & int64(4294967295)
	v129 = v2
	v131 = int32(30)
	goto L26
L12:
	;
	v57 = base.I32_wrap_i64(v22)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if int32(0) < v58 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v62
	v65 = F_palloc_mul(m, int32(4), v62)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L3
	} else {
		goto L16
	}
L15:
	;
	v262 = int32(30)
	goto L10
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v65
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	if v68 <= int32(0) {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v76 = v2
	goto L18
L18:
	;
	v87 = v76 << (uint(int32(2)) % 32)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v87+v88))) = v76
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if int32(0) < v91 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L9
L20:
	;
	v94 = int32(_a_F_spg_quad_inner_consistent_0)
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0])) = v97
	v99 = F_box_copy(m, v53)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L3
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v118 = v76 + int32(1)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	if v118 < v119 {
		v76 = v118
		goto L18
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0])) = v95
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v103+v87))) = v99
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v110 = F_spg_key_orderbys_distances(m, base.I64_extend_i32_u(v99), int32(0), v108, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v112+v87))) = v110
	goto L22
L25:
	;
	goto L19
L26:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v142 = v139 + v129*int32(56)
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v142)+48))
	v145 = v143 & int64(4294967295)
	v146 = base.I32_wrap_i64(v143)
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+6)))
	switch v147 - int32(1) {
	case 0:
		goto L36
	default:
		goto L31
	case 4:
		goto L35
	case 5:
		goto L30
	case 7:
		goto L32
	case 9, 28:
		goto L34
	case 10, 29:
		goto L33
	}
L27:
	;
	v262 = v246
	goto L10
L28:
	;
	v252 = v129 + int32(1)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v252 < v253 {
		v129 = v252
		v131 = v246
		goto L26
	} else {
		goto L56
	}
L29:
	;
	v244 = v243 & v131
	if v244 != 0 {
		v246 = v244
		goto L28
	} else {
		goto L55
	}
L30:
	;
	v236 = F_getQuadrant(m, v57, v146)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L3
	} else {
		goto L54
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L3
	} else {
		goto L51
	}
L32:
	;
	v180 = F_DirectFunctionCall2Coll(m, int32(262), int32(0), v145, v122)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L3
	} else {
		goto L45
	}
L33:
	;
	v173 = F_DirectFunctionCall2Coll(m, int32(260), int32(0), v122, v145)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L3
	} else {
		goto L43
	}
L34:
	;
	v166 = F_DirectFunctionCall2Coll(m, int32(256), int32(0), v122, v145)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L3
	} else {
		goto L41
	}
L35:
	;
	v159 = F_DirectFunctionCall2Coll(m, int32(261), int32(0), v122, v145)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L3
	} else {
		goto L39
	}
L36:
	;
	v152 = F_DirectFunctionCall2Coll(m, int32(258), int32(0), v122, v145)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	if v152 == int64(0) {
		v246 = v131
		goto L28
	} else {
		goto L38
	}
L38:
	;
	v243 = int32(24)
	goto L29
L39:
	;
	if v159 == int64(0) {
		v246 = v131
		goto L28
	} else {
		goto L40
	}
L40:
	;
	v243 = int32(6)
	goto L29
L41:
	;
	if v166 == int64(0) {
		v246 = v131
		goto L28
	} else {
		goto L42
	}
L42:
	;
	v243 = int32(12)
	goto L29
L43:
	;
	if v173 == int64(0) {
		v246 = v131
		goto L28
	} else {
		goto L44
	}
L44:
	;
	v243 = int32(18)
	goto L29
L45:
	;
	if v180 != int64(0) {
		v246 = v131
		goto L28
	} else {
		goto L46
	}
L46:
	;
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v146)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v184
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v146)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v186
	v189 = v16 + int32(-48)
	v190 = F_getQuadrant(m, v57, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	v192 = *(*float64)(unsafe.Add(mBase, uint32(v146)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v192
	v194 = F_getQuadrant(m, v57, v189)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v146)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v196
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v146)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v198
	v200 = F_getQuadrant(m, v57, v189)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v202 = *(*float64)(unsafe.Add(mBase, uint32(v146)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+16)) = v202
	v204 = int32(1)
	v213 = F_getQuadrant(m, v57, v189)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v243 = v204<<(uint(v194)%32) | v204<<(uint(v190)%32) | v204<<(uint(v200)%32) | v204<<(uint(v213)%32)
	goto L29
L51:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221+v129*int32(56))+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v225
	F_errmsg_internal(m, int32(_a_F_spg_quad_inner_consistent_1), v18)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_spg_quad_inner_consistent_2), int32(365), int32(_a_F_spg_quad_inner_consistent_3))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	v243 = int32(1) << (uint(v236) % 32)
	goto L29
L55:
	;
	v262 = int32(0)
	goto L10
L56:
	;
	goto L27
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v272
	v275 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v272))) = v275
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v277)+4)) = v275
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+8)) = v275
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v283)+12)) = v275
	v286 = int32(4)
	v288 = F_palloc_mul(m, v286, v286)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v290 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v290
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v288
	v294 = v53 + int32(16)
	v298 = v290
	v302 = int32(1)
	goto L59
L59:
	;
	if int32(base.Ui32(v262)>>(uint(v302)%32))&int32(1) != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L9
L61:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v320 = v302 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v315+v298<<(uint(int32(2))%32)))) = v320
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if int32(0) < v322 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v391 = v298
	goto L63
L63:
	;
	v395 = v302 + int32(1)
	if v395 != int32(5) {
		v298 = v391
		v302 = v395
		goto L59
	} else {
		goto L74
	}
L64:
	;
	v325 = int32(_a_F_spg_quad_inner_consistent_0)
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0]))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0])) = v328
	v331 = F_palloc(m, int32(32))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L3
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v389 = v387 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v389
	v391 = v389
	goto L63
L67:
	;
	switch v320 {
	case 0:
		goto L72
	case 1:
		goto L71
	case 2:
		goto L70
	case 3:
		goto L69
	default:
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0])) = v326
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v367+v368<<(uint(int32(2))%32)))) = v331
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v377 = F_spg_key_orderbys_distances(m, base.I64_extend_i32_u(v331), int32(0), v375, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L3
	} else {
		goto L73
	}
L69:
	;
	v357 = *(*float64)(unsafe.Add(mBase, uint32(v57)))
	*(*float64)(unsafe.Add(mBase, uint32(v331))) = v357
	v359 = *(*float64)(unsafe.Add(mBase, uint32(v53)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v331)+8)) = v359
	v361 = *(*float64)(unsafe.Add(mBase, uint32(v53)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v331)+16)) = v361
	v363 = *(*float64)(unsafe.Add(mBase, uint32(v57)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v331)+24)) = v363
	goto L68
L70:
	;
	v349 = *(*int64)(unsafe.Add(mBase, uint32(v57)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v331)+8)) = v349
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
	*(*int64)(unsafe.Add(mBase, uint32(v331))) = v351
	v353 = *(*int64)(unsafe.Add(mBase, uint32(v294)))
	*(*int64)(unsafe.Add(mBase, uint32(v331)+16)) = v353
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v294)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v331)+24)) = v355
	goto L68
L71:
	;
	v341 = *(*float64)(unsafe.Add(mBase, uint32(v53)))
	*(*float64)(unsafe.Add(mBase, uint32(v331))) = v341
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v57)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v331)+8)) = v343
	v345 = *(*float64)(unsafe.Add(mBase, uint32(v57)))
	*(*float64)(unsafe.Add(mBase, uint32(v331)+16)) = v345
	v347 = *(*float64)(unsafe.Add(mBase, uint32(v53)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v331)+24)) = v347
	goto L68
L72:
	;
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v331)+8)) = v333
	v335 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
	*(*int64)(unsafe.Add(mBase, uint32(v331))) = v335
	v337 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
	*(*int64)(unsafe.Add(mBase, uint32(v331)+16)) = v337
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v57)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v331)+24)) = v339
	goto L68
L73:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v379+v380<<(uint(int32(2))%32)))) = v377
	goto L66
L74:
	;
	goto L60
}
func F_spg_quad_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v18 int32
	_ = v18
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v26 int32
	_ = v26
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int64
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_palloc0(m, int32(16))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v17 = *(*float64)(unsafe.Add(mBase, uint32(v13)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v18 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v13))) = base.F64_div(v57, base.F64_convert_i32_s(v54))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = base.F64_div(v56, base.F64_convert_i32_s(v62))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(4)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = base.I64_extend_i32_u(v13)
	v70 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v70)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v74 = F_palloc_mul(m, int32(4), v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L10
	}
L4:
	;
	v21 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
	v54 = v18
	v56 = v21
	v57 = v17
	goto L3
L5:
	;
	goto L6
L6:
	;
	v22 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
	v26 = int32(0)
	v29 = v22
	v30 = v17
	goto L7
L7:
	;
	v33 = v26 << (uint(int32(3)) % 32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33+v34)))
	v37 = *(*float64)(unsafe.Add(mBase, uint32(v36)))
	v38 = base.F64_add(v37, v30)
	*(*float64)(unsafe.Add(mBase, uint32(v13))) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40+v33)))
	v43 = *(*float64)(unsafe.Add(mBase, uint32(v42)+8))
	v44 = base.F64_add(v43, v29)
	*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = v44
	v47 = v26 + int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v47 < v48 {
		v26 = v47
		v29 = v44
		v30 = v38
		goto L7
	} else {
		goto L9
	}
L8:
	;
	v54 = v48
	v56 = v44
	v57 = v38
	goto L3
L9:
	;
	goto L8
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v74
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v79 = F_palloc_mul(m, int32(8), v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v79
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if int32(0) < v82 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v89 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	return int64(0)
L15:
	;
	v96 = v89 << (uint(int32(3)) % 32)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v96+v97)))
	v101 = F_getQuadrant(m, v13, base.I32_wrap_i64(v99))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v103+v96))) = v99 & int64(4294967295)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v112 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v108+v89<<(uint(int32(2))%32)))) = v101 - v112
	v116 = v89 + v112
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v116 < v117 {
		v89 = v116
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
}
func F_spg_range_quad_inner_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int64
	_ = v230
	var v232 int32
	_ = v232
	var v236 int64
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
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
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v611 int64
	_ = v611
	var v614 int64
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(224)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+37)))
	if v23 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v19 + int32(224)
	return int64(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v26
	v29 = F_palloc_mul(m, int32(4), v26)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+38)))
	if v63 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	return int64(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v29
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if v34 <= int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v38 = int32(0)
	goto L8
L8:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v54+v38<<(uint(int32(2))%32)))) = v38
	v60 = v38 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if v60 < v61 {
		v38 = v60
		goto L8
	} else {
		goto L10
	}
L9:
	;
	goto L1
L10:
	;
	goto L9
L11:
	;
	v581 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v581
	v583 = int32(_a_F_spg_range_quad_inner_consistent_0)
	v584 = *(*int32)(unsafe.Add(mBase, _c_F_spg_range_quad_inner_consistent[0]))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_spg_range_quad_inner_consistent[0])) = v586
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if v581 < v588 {
		goto L154
	} else {
		goto L155
	}
L12:
	;
	v66 = int32(6)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v67 <= int32(0) {
		v154 = v66
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v178 = F_pg_detoast_datum(m, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L5
	} else {
		goto L44
	}
L15:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v172 = F_palloc_mul(m, int32(4), v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L43
	}
L16:
	;
	v70 = v66
	v78 = v2
	goto L17
L17:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v89 = v86 + v78*int32(56)
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+6)))
	if v90 == int32(16) {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v154 = int32(0)
	goto L15
L19:
	;
	goto L18
L20:
	;
	v148 = v78 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v148 < v149 {
		v70 = v145
		v78 = v148
		goto L17
	} else {
		goto L42
	}
L21:
	;
	if v142 == int32(0) {
		goto L19
	} else {
		goto L41
	}
L22:
	;
	v142 = v70 & int32(4)
	goto L21
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)+48))
	v94 = F_pg_detoast_datum(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v94+int32(base.Ui32(v96)>>(uint(int32(2))%32))-int32(1)))))
	goto L25
L25:
	;
	if base.Ui32(int32(6)) <= base.Ui32(v90-int32(1)) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L38
	}
L27:
	;
	v142 = v70 & int32(2)
	goto L21
L28:
	;
	if v102&int32(1) == int32(0) {
		goto L22
	} else {
		goto L37
	}
L29:
	;
	if v102&int32(1) != 0 {
		goto L27
	} else {
		goto L36
	}
L30:
	;
	if v102&int32(1) == int32(0) {
		goto L22
	} else {
		goto L35
	}
L31:
	;
	switch v90 - int32(7) {
	case 0:
		goto L30
	case 1:
		goto L29
	default:
		goto L26
	case 11:
		goto L28
	}
L32:
	;
	goto L33
L33:
	;
	if v102&int32(1) != 0 {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	goto L22
L35:
	;
	v145 = v70
	goto L20
L36:
	;
	v145 = v70
	goto L20
L37:
	;
	goto L27
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v90
	F_errmsg_internal(m, int32(_a_F_spg_range_quad_inner_consistent_1), v19+int32(16))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_spg_range_quad_inner_consistent_2), int32(403), int32(_a_F_spg_range_quad_inner_consistent_3))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
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
	v145 = v142
	goto L20
L42:
	;
	v154 = v145
	goto L15
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v172
	v565 = v154
	v567 = v22 + int32(48)
	v569 = v2
	goto L11
L44:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v181 = F_range_get_typcache(m, l0, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	F_range_deserialize(m, v181, v178, v19+int32(128), v19+int32(112), v19+int32(111))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v191 <= int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v196 = F_palloc_mul(m, int32(4), v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v203 = int32(62)
	v211 = v2
	v216 = v2
	goto L52
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v196
	v565 = int32(62)
	v567 = v22 + int32(48)
	v569 = v2
	goto L11
L51:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	v561 = F_palloc_mul(m, int32(4), v560)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L5
	} else {
		goto L153
	}
L52:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v222 = v219 + v211*int32(56)
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+6)))
	if v223 == int32(16) {
		goto L69
	} else {
		goto L70
	}
L53:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v541 = F_palloc_mul(m, int32(4), v540)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L5
	} else {
		goto L151
	}
L54:
	;
	v536 = v211 + int32(1)
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v536 < v537 {
		v203 = v526
		v211 = v536
		v216 = v532
		goto L52
	} else {
		goto L150
	}
L55:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v519 = F_palloc_mul(m, int32(4), v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L5
	} else {
		goto L148
	}
L56:
	;
	if v507 != 0 {
		v526 = v507
		v532 = v504
		goto L54
	} else {
		goto L147
	}
L57:
	;
	if v467 != 0 {
		goto L129
	} else {
		goto L130
	}
L58:
	;
	v456 = v446 & int32(56)
	v459 = F_range_cmp_bounds(m, v181, v19+int32(128), v451)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L5
	} else {
		goto L119
	}
L59:
	;
	if v441 == int32(0) {
		v466 = v443
		v467 = v437
		v468 = v438
		v470 = v440
		v472 = v442
		goto L57
	} else {
		goto L118
	}
L60:
	;
	v431 = F_range_cmp_bounds(m, v181, v19+int32(128), v423)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L5
	} else {
		goto L114
	}
L61:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+71)))
	if v415 != 0 {
		v514 = v414
		goto L55
	} else {
		goto L112
	}
L62:
	;
	v408 = v400
	v409 = v401
	v410 = v406
	v411 = v403
	v412 = int32(0)
	v413 = v404
	v414 = v405
	goto L61
L63:
	;
	v398 = int32(0)
	v400 = v393
	v401 = v398
	v403 = v396
	v404 = v252
	v405 = v397
	v406 = v398
	goto L62
L64:
	;
	v400 = v203
	v401 = v252
	v403 = v252
	v404 = v252
	v405 = v216
	v406 = v19 + int32(72)
	goto L62
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L5
	} else {
		goto L109
	}
L66:
	;
	v339 = v19 + int32(208)
	v341 = v19 + int32(192)
	F_range_deserialize(m, v181, v178, v339, v341, v19+int32(191))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L5
	} else {
		goto L96
	}
L67:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+71)))
	if v328 != int32(1) {
		goto L93
	} else {
		goto L94
	}
L68:
	;
	v326 = int32(0)
	v446 = v203 & int32(30)
	v447 = v19 + int32(72)
	v448 = v326
	v450 = v326
	v451 = v19 + int32(88)
	v452 = v216
	goto L58
L69:
	;
	v226 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+98)) = uint8(v226)
	v228 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+96)) = uint16(v228)
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v222)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+88)) = v230
	v232 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+82)) = uint8(v232)
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+80)) = uint16(v228)
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v222)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+71)) = uint8(v232)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v236
	goto L68
L70:
	;
	goto L71
L71:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v222)+48))
	v241 = F_pg_detoast_datum(m, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	v244 = v19 + int32(88)
	F_range_deserialize(m, v181, v241, v244, v19+int32(72), v19+int32(71))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	v251 = int32(1)
	v252 = int32(0)
	switch v223 - v251 {
	case 0:
		v408 = v203
		v409 = v252
		v410 = v244
		v411 = v252
		v412 = v251
		v413 = v252
		v414 = v216
		goto L61
	case 1:
		goto L64
	case 2:
		goto L78
	case 3:
		goto L77
	case 4:
		goto L76
	case 5:
		goto L75
	case 6:
		goto L74
	case 7:
		goto L67
	default:
		goto L65
	case 17:
		goto L66
	}
L74:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+71)))
	if v312 != 0 {
		v526 = v203
		v532 = v216
		goto L54
	} else {
		goto L92
	}
L75:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+71)))
	if v267 != 0 {
		v393 = v203
		v396 = v252
		v397 = v216
		goto L63
	} else {
		goto L79
	}
L76:
	;
	v408 = v203
	v409 = v252
	v410 = int32(0)
	v411 = v19 + int32(72)
	v412 = v251
	v413 = v252
	v414 = v216
	goto L61
L77:
	;
	v393 = v203
	v396 = v19 + int32(88)
	v397 = v216
	goto L63
L78:
	;
	v400 = v203
	v401 = v19 + int32(88)
	v403 = v252
	v404 = v19 + int32(72)
	v405 = v216
	v406 = int32(0)
	goto L62
L79:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v272 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v274 = v19 + int32(48)
	v276 = v19 + int32(32)
	F_range_deserialize(m, v181, v272, v274, v276, v19+int32(31))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L5
	} else {
		goto L83
	}
L81:
	;
	v282 = v252
	v284 = int32(0)
	goto L82
L82:
	;
	v285 = F_adjacent_inner_consistent(m, v181, v19+int32(88), v19+int32(112), v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L5
	} else {
		goto L84
	}
L83:
	;
	v282 = v274
	v284 = v276
	goto L82
L84:
	;
	v292 = F_adjacent_inner_consistent(m, v181, v19+int32(72), v19+int32(128), v282)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	if int32(0) < v292 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v300 = int32(6)
	goto L88
L87:
	;
	v300 = v292 >> (uint(int32(31)) % 32) & int32(24)
	goto L88
L88:
	;
	if int32(0) < v285 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v308 = int32(18)
	goto L91
L90:
	;
	v308 = v285 >> (uint(int32(31)) % 32) & int32(12)
	goto L91
L91:
	;
	v393 = v203 & (v300 | v308)
	v396 = v252
	v397 = int32(1)
	goto L63
L92:
	;
	goto L68
L93:
	;
	v420 = v203
	v421 = v252
	v422 = v19 + int32(72)
	v423 = v19 + int32(88)
	v424 = int32(0)
	v425 = v252
	v426 = v216
	goto L60
L94:
	;
	goto L95
L95:
	;
	v504 = v216
	v507 = v203 & int32(32)
	goto L56
L96:
	;
	v347 = v19 + int32(168)
	v349 = v19 + int32(152)
	F_range_deserialize(m, v181, v241, v347, v349, v19+int32(151))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+151)))
	if v356 != 0 {
		v375 = int32(5)
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v504 = v216
	v507 = v203 & (int32(1) << (uint(v375) % 32))
	goto L56
L99:
	;
	v357 = F_range_cmp_bounds(m, v181, v347, v339)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	v361 = F_range_cmp_bounds(m, v181, v349, v341)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	if int32(0) <= v361 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v365 = int32(1)
	goto L104
L103:
	;
	v365 = int32(2)
	goto L104
L104:
	;
	if int32(0) <= v357 {
		v375 = v365
		goto L98
	} else {
		goto L105
	}
L105:
	;
	if int32(0) <= v361 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v372 = int32(4)
	goto L108
L107:
	;
	v372 = int32(3)
	goto L108
L108:
	;
	v375 = v372
	goto L98
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v223
	F_errmsg_internal(m, int32(_a_F_spg_range_quad_inner_consistent_1), v19)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_spg_range_quad_inner_consistent_2), int32(653), int32(_a_F_spg_range_quad_inner_consistent_3))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	v417 = v408 & int32(30)
	if v411 == int32(0) {
		v437 = v409
		v438 = v410
		v440 = v412
		v441 = v413
		v442 = v414
		v443 = v417
		goto L59
	} else {
		goto L113
	}
L113:
	;
	v420 = v417
	v421 = v409
	v422 = v410
	v423 = v411
	v424 = v412
	v425 = v413
	v426 = v414
	goto L60
L114:
	;
	if v431 <= int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v435 = v420 & int32(38)
	goto L117
L116:
	;
	v435 = v420
	goto L117
L117:
	;
	v437 = v421
	v438 = v422
	v440 = v424
	v441 = v425
	v442 = v426
	v443 = v435
	goto L59
L118:
	;
	v446 = v443
	v447 = v437
	v448 = v438
	v450 = v440
	v451 = v441
	v452 = v442
	goto L58
L119:
	;
	if v459 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v461 = v446
	goto L122
L121:
	;
	v461 = v456
	goto L122
L122:
	;
	if v450 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v462 = v461
	goto L125
L124:
	;
	v462 = v446
	goto L125
L125:
	;
	if int32(0) < v459 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v465 = v456
	goto L128
L127:
	;
	v465 = v462
	goto L128
L128:
	;
	v466 = v465
	v467 = v447
	v468 = v448
	v470 = v450
	v472 = v452
	goto L57
L129:
	;
	v479 = F_range_cmp_bounds(m, v181, v19+int32(112), v467)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L5
	} else {
		goto L132
	}
L130:
	;
	v484 = v466
	goto L131
L131:
	;
	if v468 == int32(0) {
		v504 = v472
		v507 = v484
		goto L56
	} else {
		goto L136
	}
L132:
	;
	if v479 <= int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v483 = v466 & int32(50)
	goto L135
L134:
	;
	v483 = v466
	goto L135
L135:
	;
	v484 = v483
	goto L131
L136:
	;
	v488 = v484 & int32(44)
	v491 = F_range_cmp_bounds(m, v181, v19+int32(112), v468)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L5
	} else {
		goto L137
	}
L137:
	;
	if v491 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v493 = v484
	goto L140
L139:
	;
	v493 = v488
	goto L140
L140:
	;
	if v470 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v494 = v493
	goto L143
L142:
	;
	v494 = v484
	goto L143
L143:
	;
	if int32(0) < v491 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v497 = v488
	goto L146
L145:
	;
	v497 = v494
	goto L146
L146:
	;
	v504 = v472
	v507 = v497
	goto L56
L147:
	;
	v514 = v504
	goto L55
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v519
	v523 = v22 + int32(48)
	v524 = int32(0)
	if v514 != 0 {
		v549 = v524
		v550 = v523
		goto L51
	} else {
		goto L149
	}
L149:
	;
	v565 = v524
	v567 = v523
	v569 = v524
	goto L11
L150:
	;
	goto L53
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v541
	v545 = v22 + int32(48)
	v546 = int32(0)
	if v532 == v546 {
		v565 = v526
		v567 = v545
		v569 = v546
		goto L11
	} else {
		goto L152
	}
L152:
	;
	v549 = v526
	v550 = v545
	goto L51
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v561
	v565 = v549
	v567 = v550
	v569 = int32(1)
	goto L11
L154:
	;
	v597 = v588
	v600 = int32(1)
	goto L157
L155:
	;
	goto L156
L156:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spg_range_quad_inner_consistent[0])) = v584
	goto L1
L157:
	;
	if int32(base.Ui32(v565)>>(uint(v600)%32))&int32(1) != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	goto L156
L159:
	;
	if v569 != 0 {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	v636 = v597
	goto L161
L161:
	;
	v639 = v600 + int32(1)
	if v639 <= v636 {
		v597 = v636
		v600 = v639
		goto L157
	} else {
		goto L166
	}
L162:
	;
	v611 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
	v614 = F_datumCopy(m, v611, int32(0), int32(-1))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L5
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v628 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v623+v624<<(uint(int32(2))%32)))) = v600 - v628
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v631 + v628
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	v636 = v635
	goto L161
L165:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*uint32)(unsafe.Add(mBase, uint32(v616+v617<<(uint(int32(2))%32)))) = uint32(v614)
	goto L164
L166:
	;
	goto L158
}
func F_spg_text_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int64
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v371 int32
	_ = v371
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if base.B2i32(v51 < int32(2))|base.B2i32(v50 <= int32(0)) != 0 {
		v194 = v50
		goto L14
	} else {
		goto L15
	}
L2:
	;
	return int64(0)
L3:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v21 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v27 == int32(18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v38 = int32(1)
	if v21&v38 != 0 {
		v50 = int32(base.Ui32(v21)>>(uint(v38)%32)) - v38
		goto L1
	} else {
		goto L13
	}
L7:
	;
	v30 = int32(16)
	goto L9
L8:
	;
	v30 = int32(0)
	goto L9
L9:
	;
	if base.Ui32((v27-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = int32(4)
	goto L12
L11:
	;
	v37 = v30
	goto L12
L12:
	;
	v50 = v37
	goto L1
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L14:
	;
	v204 = int32(3964)
	if v204 <= v194 {
		goto L61
	} else {
		goto L62
	}
L15:
	;
	v60 = v50
	v67 = int32(1)
	goto L16
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+v67<<(uint(int32(3))%32))))
	v75 = F_pg_detoast_datum_packed(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L18
	}
L17:
	;
	v194 = v185
	goto L14
L18:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v79 = int32(1)
	v80 = v78 & v79
	v82 = v77 & v79
	v83 = int32(0)
	if v78 == v79 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if v176 < v60 {
		goto L56
	} else {
		goto L57
	}
L20:
	;
	if v77 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L21:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v89 == int32(18) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v100 = int32(1)
	if v80 != 0 {
		v110 = int32(base.Ui32(v78)>>(uint(v100)%32)) - v100
		goto L20
	} else {
		goto L30
	}
L24:
	;
	v92 = int32(16)
	goto L26
L25:
	;
	v92 = int32(0)
	goto L26
L26:
	;
	if base.Ui32((v89-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v99 = int32(4)
	goto L29
L28:
	;
	v99 = v92
	goto L29
L29:
	;
	v110 = v99
	goto L20
L30:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v110 = int32(base.Ui32(v104)>>(uint(int32(2))%32)) - int32(4)
	goto L20
L31:
	;
	if v110 < v137 {
		goto L42
	} else {
		goto L43
	}
L32:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	if v116 == int32(18) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v127 = int32(1)
	if v82 != 0 {
		v137 = int32(base.Ui32(v77)>>(uint(v127)%32)) - v127
		goto L31
	} else {
		goto L41
	}
L35:
	;
	v119 = int32(16)
	goto L37
L36:
	;
	v119 = int32(0)
	goto L37
L37:
	;
	if base.Ui32((v116-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v126 = int32(4)
	goto L40
L39:
	;
	v126 = v119
	goto L40
L40:
	;
	v137 = v126
	goto L31
L41:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v137 = int32(base.Ui32(v131)>>(uint(int32(2))%32)) - int32(4)
	goto L31
L42:
	;
	v139 = v110
	goto L44
L43:
	;
	v139 = v137
	goto L44
L44:
	;
	if v139 <= int32(0) {
		v176 = v83
		goto L19
	} else {
		goto L45
	}
L45:
	;
	if v82 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v144 = int32(1)
	goto L48
L47:
	;
	v144 = int32(4)
	goto L48
L48:
	;
	if v80 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v148 = int32(1)
	goto L51
L50:
	;
	v148 = int32(4)
	goto L51
L51:
	;
	v150 = v75 + v144
	v153 = v17 + v148
	v154 = v83
	goto L52
L52:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	if v162 != v163 {
		v176 = v154
		goto L19
	} else {
		goto L54
	}
L53:
	;
	v176 = v139
	goto L19
L54:
	;
	v165 = int32(1)
	v170 = v154 + v165
	if v170 != v139 {
		v150 = v150 + v165
		v153 = v153 + v165
		v154 = v170
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v185 = v176
	goto L58
L57:
	;
	v185 = v60
	goto L58
L58:
	;
	v187 = v67 + int32(1)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v188 <= v187 {
		v194 = v185
		goto L14
	} else {
		goto L59
	}
L59:
	;
	if int32(0) < v185 {
		v60 = v185
		v67 = v187
		goto L16
	} else {
		goto L60
	}
L60:
	;
	goto L17
L61:
	;
	v207 = v204
	goto L63
L62:
	;
	v207 = v194
	goto L63
L63:
	;
	if v194 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v250 = F_palloc_mul(m, int32(16), v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L2
	} else {
		goto L79
	}
L65:
	;
	v210 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v210)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v212 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v212)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v216 = v207 + int32(4)
	v217 = F_palloc(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v220 = v207 + int32(1)
	if base.Ui32(v220) <= base.Ui32(int32(127)) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v207 != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v223 = int32(1)
	v226 = v220<<(uint(v223)%32) | v223
	*(*uint8)(unsafe.Add(mBase, uint32(v217))) = uint8(v226)
	v233 = v223
	goto L69
L71:
	;
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = v216 << (uint(int32(2)) % 32)
	v233 = int32(4)
	goto L69
L73:
	;
	v235 = int32(1)
	if v214&v235 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = base.I64_extend_i32_u(v217)
	goto L64
L76:
	;
	v239 = v235
	goto L78
L77:
	;
	v239 = int32(4)
	goto L78
L78:
	;
	base.MemoryCopy(m, v217+v233, v17+v239, v207)
	goto L75
L79:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if int32(0) < v252 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v256 = int32(0)
	goto L83
L81:
	;
	v332 = v252
	goto L82
L82:
	;
	F_pg_qsort(m, v250, v332, int32(16), int32(267))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L2
	} else {
		goto L104
	}
L83:
	;
	v269 = v256 << (uint(int32(3)) % 32)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v269+v270)))
	v273 = F_pg_detoast_datum_packed(m, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L2
	} else {
		goto L86
	}
L84:
	;
	v332 = v327
	goto L82
L85:
	;
	if base.Ui32(v207) < base.Ui32(v304) {
		goto L97
	} else {
		goto L98
	}
L86:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	if v275 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+1)))
	if v281 == int32(18) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v292 = int32(1)
	if v275&v292 != 0 {
		v304 = int32(base.Ui32(v275)>>(uint(v292)%32)) - v292
		goto L85
	} else {
		goto L96
	}
L90:
	;
	v284 = int32(16)
	goto L92
L91:
	;
	v284 = int32(0)
	goto L92
L92:
	;
	if base.Ui32((v281-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v291 = int32(4)
	goto L95
L94:
	;
	v291 = v284
	goto L95
L95:
	;
	v304 = v291
	goto L85
L96:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v304 = int32(base.Ui32(v298)>>(uint(int32(2))%32)) - int32(4)
	goto L85
L97:
	;
	v307 = int32(1)
	if v275&v307 != 0 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v315 = int32(_a_F_spg_text_picksplit_0)
	goto L99
L99:
	;
	v318 = v250 + v256<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v318)+8)) = v256
	*(*uint16)(unsafe.Add(mBase, uint32(v318)+12)) = uint16(v315)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v321+v269)))
	*(*int64)(unsafe.Add(mBase, uint32(v318))) = v323
	v326 = v256 + int32(1)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v326 < v327 {
		v256 = v326
		goto L83
	} else {
		goto L103
	}
L100:
	;
	v311 = v307
	goto L102
L101:
	;
	v311 = int32(4)
	goto L102
L102:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273+v311+v207))))
	v315 = v314
	goto L99
L103:
	;
	goto L84
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(0)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v349 = F_palloc_mul(m, int32(8), v348)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v349
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v354 = F_palloc_mul(m, int32(4), v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L2
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v354
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v359 = F_palloc_mul(m, int32(8), v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v359
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if int32(0) < v362 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v371 = int32(0)
	goto L111
L109:
	;
	goto L110
L110:
	;
	return int64(0)
L111:
	;
	v382 = v250 + v371<<(uint(int32(4))%32)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v384 = F_pg_detoast_datum_packed(m, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L2
	} else {
		goto L113
	}
L112:
	;
	goto L110
L113:
	;
	v386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v382)+12)))
	if v371 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v403 != int32(1) {
		goto L124
	} else {
		goto L125
	}
L115:
	;
	v389 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v382-int32(4)))))
	if v389 == v386 {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v391+v392<<(uint(int32(3))%32)))) = base.I64_extend16_s(base.I64_extend_i32_u(v386))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v399 + int32(1)
	goto L114
L118:
	;
	goto L117
L119:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v491+v492<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v487)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v504 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v498+v499<<(uint(int32(2))%32)))) = v503 - v504
	v508 = v371 + v504
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v508 < v509 {
		v371 = v508
		goto L111
	} else {
		goto L151
	}
L120:
	;
	v463 = v460 + (v207 ^ int32(-1))
	v465 = v463 + int32(4)
	v466 = F_palloc(m, v465)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L2
	} else {
		goto L144
	}
L121:
	;
	if v407 != 0 {
		goto L141
	} else {
		goto L142
	}
L122:
	;
	v448 = F_palloc(m, int32(4))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L2
	} else {
		goto L140
	}
L123:
	;
	v439 = int32(1)
	v442 = int32(base.Ui32(v403)>>(uint(v439)%32)) - v439
	if base.Ui32(v207) < base.Ui32(v442) {
		goto L121
	} else {
		goto L139
	}
L124:
	;
	v407 = v403 & int32(1)
	if v407 != 0 {
		goto L123
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+1)))
	if v424 == int32(18) {
		goto L132
	} else {
		goto L133
	}
L127:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	v412 = int32(base.Ui32(v408)>>(uint(int32(2))%32)) - int32(4)
	if base.Ui32(v412) <= base.Ui32(v207) {
		goto L122
	} else {
		goto L128
	}
L128:
	;
	if v407 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v416 = int32(1)
	goto L131
L130:
	;
	v416 = int32(4)
	goto L131
L131:
	;
	v460 = v412
	v462 = v384 + v416 + v207 + int32(1)
	goto L120
L132:
	;
	v427 = int32(16)
	goto L134
L133:
	;
	v427 = int32(0)
	goto L134
L134:
	;
	if base.Ui32((v424-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v434 = int32(4)
	goto L137
L136:
	;
	v434 = v427
	goto L137
L137:
	;
	if base.Ui32(v434) <= base.Ui32(v207) {
		goto L122
	} else {
		goto L138
	}
L138:
	;
	v460 = v434
	v462 = v384 + v207 + int32(2)
	goto L120
L139:
	;
	goto L122
L140:
	;
	v450 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v448))) = uint8(v450)
	v487 = v448
	goto L119
L141:
	;
	v454 = int32(1)
	goto L143
L142:
	;
	v454 = int32(4)
	goto L143
L143:
	;
	v460 = v442
	v462 = v384 + v454 + v207 + int32(1)
	goto L120
L144:
	;
	v468 = v460 - v207
	if base.Ui32(v468) <= base.Ui32(int32(127)) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	if v463 == int32(0) {
		v487 = v466
		goto L119
	} else {
		goto L150
	}
L146:
	;
	v471 = int32(1)
	v474 = v468<<(uint(v471)%32) | v471
	*(*uint8)(unsafe.Add(mBase, uint32(v466))) = uint8(v474)
	if v463 != 0 {
		v481 = v471
		goto L145
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v466))) = v465 << (uint(int32(2)) % 32)
	v481 = int32(4)
	goto L145
L149:
	;
	v487 = v466
	goto L119
L150:
	;
	base.MemoryCopy(m, v466+v481, v462, v463)
	v487 = v466
	goto L119
L151:
	;
	goto L112
}
