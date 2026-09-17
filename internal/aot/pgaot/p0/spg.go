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
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v3 = int32(0)
	v8 = l1 + int32(8)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = v8 + int32(base.Ui32(v9)>>(uint(int32(16))%32))
	v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
	if v13 < v3 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L12
	} else {
		goto L25
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L12
	} else {
		goto L22
	}
L3:
	;
	return v93
L4:
	;
	v19 = int32(base.Ui32(v9)>>(uint(int32(3))%32)) & int32(_a_F_spgExtractNodeLabels_0)
	if v19 == int32(0) {
		v93 = v3
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v43 = F_palloc(m, int32(base.Ui32(v9)>>(uint(int32(1))%32))&int32(_a_F_spgExtractNodeLabels_1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v24 = v12
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
	if v35 != v19 {
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
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v47&int32(_a_F_spgExtractNodeLabels_2) == int32(0) {
		v93 = v43
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v58 = v8 + int32(base.Ui32(v47)>>(uint(int32(16))%32))
	v59 = int32(0)
	goto L15
L15:
	;
	v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(v58)+6)))
	if v62 < int32(0) {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v93 = v43
	goto L3
L17:
	;
	v66 = v58 + int32(8)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
	if v70 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v74 = v73
	goto L20
L19:
	;
	v74 = v66
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+v59<<(uint(int32(2))%32)))) = v74
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+6)))
	v77 = int32(_a_F_spgExtractNodeLabels_0)
	v81 = v59 + int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v81) < base.Ui32(int32(base.Ui32(v82)>>(uint(int32(3))%32))&v77) {
		v58 = v58 + v76&v77
		v59 = v81
		goto L15
	} else {
		goto L21
	}
L21:
	;
	goto L16
L22:
	;
	F_errmsg_internal(m, int32(_a_F_spgExtractNodeLabels_3), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_spgExtractNodeLabels_4), int32(1173), int32(_a_F_spgExtractNodeLabels_5))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	F_errmsg_internal(m, int32(_a_F_spgExtractNodeLabels_3), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_spgExtractNodeLabels_4), int32(1184), int32(_a_F_spgExtractNodeLabels_5))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
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
func F_spg_box_quad_choose(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v20 float64
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+12)))
	if v12 == int32(0) {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
		v20 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
		v21 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		if base.F64_gt(v20, v21) != 0 {
			v23 = int32(8)
		} else {
			v23 = int32(0)
		}
		v26 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		v27 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
		if base.F64_gt(v26, v27) != 0 {
			v29 = v23 | int32(4)
		} else {
			v29 = v23
		}
		v32 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
		v33 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
		if base.F64_gt(v32, v33) != 0 {
			v35 = v29 | int32(2)
		} else {
			v35 = v29
		}
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = base.F64_gt(v15, v16) | v35
	} else {
	}
	return int32(0)
}
func F_spg_box_quad_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 float64
	_ = v71
	var v74 float64
	_ = v74
	var v77 float64
	_ = v77
	var v80 float64
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 float64
	_ = v131
	var v134 float64
	_ = v134
	var v137 float64
	_ = v137
	var v140 float64
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
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
	var v190 float64
	_ = v190
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v23 = F_palloc(m, v20<<(uint(int32(3))%32))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
		v30 = F_palloc(m, v27<<(uint(int32(3))%32))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			v35 = F_palloc(m, v32<<(uint(int32(3))%32))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v40 = F_palloc(m, v37<<(uint(int32(3))%32))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					if int32(0) < v42 {
						v46 = int32(0)
						for {
							v64 = v46 << (uint(int32(3)) % 32)
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v66+v46<<(uint(int32(2))%32))))
							v71 = *(*float64)(unsafe.Add(mBase, uint32(v70)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v23+v64))) = v71
							v74 = *(*float64)(unsafe.Add(mBase, uint32(v70)))
							*(*float64)(unsafe.Add(mBase, uint32(v64+v30))) = v74
							v77 = *(*float64)(unsafe.Add(mBase, uint32(v70)+24))
							*(*float64)(unsafe.Add(mBase, uint32(v64+v35))) = v77
							v80 = *(*float64)(unsafe.Add(mBase, uint32(v70)+8))
							*(*float64)(unsafe.Add(mBase, uint32(v64+v40))) = v80
							v83 = v46 + int32(1)
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
							if v83 < v84 {
								v46 = v83
								continue
							} else {
								break
							}
							break
						}
						v89 = v84
					} else {
						v89 = v42
					}
					F_pg_qsort(m, v23, v89, int32(8), int32(1301))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
						F_pg_qsort(m, v30, v107, int32(8), int32(1301))
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return int32(0)
						} else {
							v112 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
							F_pg_qsort(m, v35, v112, int32(8), int32(1301))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
								F_pg_qsort(m, v40, v117, int32(8), int32(1301))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									v122 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
									v124 = F_palloc(m, int32(32))
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return int32(0)
									} else {
										v126 = int32(2)
										v127 = base.I32_div_s(v122, v126)
										v129 = v127 << (uint(int32(3)) % 32)
										v131 = *(*float64)(unsafe.Add(mBase, uint32(v23+v129)))
										*(*float64)(unsafe.Add(mBase, uint32(v124)+16)) = v131
										v134 = *(*float64)(unsafe.Add(mBase, uint32(v129+v30)))
										*(*float64)(unsafe.Add(mBase, uint32(v124))) = v134
										v137 = *(*float64)(unsafe.Add(mBase, uint32(v129+v35)))
										*(*float64)(unsafe.Add(mBase, uint32(v124)+24)) = v137
										v140 = *(*float64)(unsafe.Add(mBase, uint32(v129+v40)))
										*(*float64)(unsafe.Add(mBase, uint32(v124)+8)) = v140
										*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = int64(16)
										*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v124
										v145 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v145)
										v147 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
										v150 = F_palloc(m, v147<<(uint(v126)%32))
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v150
											v153 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
											v156 = F_palloc(m, v153<<(uint(int32(2))%32))
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v156
												v159 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
												if int32(0) < v159 {
													v168 = int32(0)
													for {
														v181 = v168 << (uint(int32(2)) % 32)
														v182 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
														v184 = *(*int32)(unsafe.Add(mBase, uint32(v181+v182)))
														v185 = *(*float64)(unsafe.Add(mBase, uint32(v184)+8))
														v186 = *(*float64)(unsafe.Add(mBase, uint32(v184)+24))
														v187 = *(*float64)(unsafe.Add(mBase, uint32(v184)+16))
														v188 = *(*float64)(unsafe.Add(mBase, uint32(v184)))
														v189 = *(*float64)(unsafe.Add(mBase, uint32(v124)+8))
														v190 = *(*float64)(unsafe.Add(mBase, uint32(v124)+24))
														v191 = *(*float64)(unsafe.Add(mBase, uint32(v124)+16))
														v192 = *(*float64)(unsafe.Add(mBase, uint32(v124)))
														v193 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
														*(*int32)(unsafe.Add(mBase, uint32(v193+v181))) = v184
														v196 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
														if base.F64_gt(v187, v191) != 0 {
															v202 = int32(8)
														} else {
															v202 = int32(0)
														}
														if base.F64_gt(v188, v192) != 0 {
															v206 = v202 | int32(4)
														} else {
															v206 = v202
														}
														if base.F64_gt(v186, v190) != 0 {
															v210 = v206 | int32(2)
														} else {
															v210 = v206
														}
														*(*int32)(unsafe.Add(mBase, uint32(v196+v181))) = base.F64_gt(v185, v189) | v210
														v214 = v168 + int32(1)
														v215 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
														if v214 < v215 {
															v168 = v214
															continue
														} else {
															break
														}
														break
													}
												} else {
												}
												return int32(0)
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
func F_spg_quad_inner_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v45 int64
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 float64
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v183 int64
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 float64
	_ = v187
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int64
	_ = v310
	var v312 int64
	_ = v312
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v318 float64
	_ = v318
	var v320 float64
	_ = v320
	var v322 float64
	_ = v322
	var v324 float64
	_ = v324
	var v326 int64
	_ = v326
	var v328 int64
	_ = v328
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	var v334 float64
	_ = v334
	var v336 float64
	_ = v336
	var v338 float64
	_ = v338
	var v340 float64
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v21 <= v2 {
		v52 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+33)))
	if v53 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v27 = F_palloc(m, v24<<(uint(int32(2))%32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v27
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v35 = F_palloc(m, v32<<(uint(int32(2))%32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v35
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v38 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v41 = int64(-4503599627370496)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v41
	v45 = int64(9218868437227405312)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v45
	*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = v45
	v52 = v13 + int32(-32)
	goto L1
L7:
	;
	goto L8
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v52 = v51
	goto L1
L9:
	;
	m.G0 = v15 - int32(-64)
	return int32(0)
L10:
	;
	v253 = F_palloc(m, int32(16))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L3
	} else {
		goto L57
	}
L11:
	;
	v122 = v2
	v124 = int32(30)
	goto L26
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if int32(0) < v56 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v60
	v64 = F_palloc(m, v60<<(uint(int32(2))%32))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L16
	}
L15:
	;
	v247 = int32(30)
	goto L10
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v64
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	if v67 <= int32(0) {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v75 = v2
	goto L18
L18:
	;
	v83 = v75 << (uint(int32(2)) % 32)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v83+v84))) = v75
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if int32(0) < v87 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L9
L20:
	;
	v90 = int32(_a_F_spg_quad_inner_consistent_0)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0])) = v93
	v95 = F_box_copy(m, v52)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L3
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v113 = v75 + int32(1)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	if v113 < v114 {
		v75 = v113
		goto L18
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0])) = v91
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v99+v83))) = v95
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v105 = F_spg_key_orderbys_distances(m, v95, int32(0), v103, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v107+v83))) = v105
	goto L22
L25:
	;
	goto L19
L26:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v132 = v129 + v122*int32(48)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+44))
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+6)))
	switch v134 - int32(1) {
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
	v247 = v231
	goto L10
L28:
	;
	v237 = v122 + int32(1)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v237 < v238 {
		v122 = v237
		v124 = v231
		goto L26
	} else {
		goto L56
	}
L29:
	;
	v229 = v228 & v124
	if v229 != 0 {
		v231 = v229
		goto L28
	} else {
		goto L55
	}
L30:
	;
	v221 = F_getQuadrant(m, v19, v133)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L3
	} else {
		goto L54
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L3
	} else {
		goto L51
	}
L32:
	;
	v167 = F_DirectFunctionCall2Coll(m, int32(254), int32(0), v133, v19)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L3
	} else {
		goto L45
	}
L33:
	;
	v160 = F_DirectFunctionCall2Coll(m, int32(252), int32(0), v19, v133)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L3
	} else {
		goto L43
	}
L34:
	;
	v153 = F_DirectFunctionCall2Coll(m, int32(248), int32(0), v19, v133)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L3
	} else {
		goto L41
	}
L35:
	;
	v146 = F_DirectFunctionCall2Coll(m, int32(253), int32(0), v19, v133)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L3
	} else {
		goto L39
	}
L36:
	;
	v139 = F_DirectFunctionCall2Coll(m, int32(250), int32(0), v19, v133)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	if v139 == int32(0) {
		v231 = v124
		goto L28
	} else {
		goto L38
	}
L38:
	;
	v228 = int32(24)
	goto L29
L39:
	;
	if v146 == int32(0) {
		v231 = v124
		goto L28
	} else {
		goto L40
	}
L40:
	;
	v228 = int32(6)
	goto L29
L41:
	;
	if v153 == int32(0) {
		v231 = v124
		goto L28
	} else {
		goto L42
	}
L42:
	;
	v228 = int32(12)
	goto L29
L43:
	;
	if v160 == int32(0) {
		v231 = v124
		goto L28
	} else {
		goto L44
	}
L44:
	;
	v228 = int32(18)
	goto L29
L45:
	;
	if v167 != 0 {
		v231 = v124
		goto L28
	} else {
		goto L46
	}
L46:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v133)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v169
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v133)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v171
	v174 = v13 + int32(-48)
	v175 = F_getQuadrant(m, v19, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	v177 = *(*float64)(unsafe.Add(mBase, uint32(v133)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v177
	v179 = F_getQuadrant(m, v19, v174)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v133)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v181
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v133)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v183
	v185 = F_getQuadrant(m, v19, v174)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v133)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v15)+16)) = v187
	v189 = int32(1)
	v198 = F_getQuadrant(m, v19, v174)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v228 = v189<<(uint(v179)%32) | v189<<(uint(v175)%32) | v189<<(uint(v185)%32) | v189<<(uint(v198)%32)
	goto L29
L51:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v206+v122*int32(48))+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v210
	F_errmsg_internal(m, int32(_a_F_spg_quad_inner_consistent_1), v15)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_spg_quad_inner_consistent_2), int32(363), int32(_a_F_spg_quad_inner_consistent_3))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
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
	v228 = int32(1) << (uint(v221) % 32)
	goto L29
L55:
	;
	v247 = int32(0)
	goto L10
L56:
	;
	goto L27
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v253
	v256 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = v256
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v258)+4)) = v256
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v261)+8)) = v256
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+12)) = v256
	v268 = F_palloc(m, int32(16))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v270 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v268
	v274 = v52 + int32(16)
	v278 = v270
	v282 = int32(1)
	goto L59
L59:
	;
	if int32(base.Ui32(v247)>>(uint(v282)%32))&int32(1) != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L9
L61:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v297 = v282 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v292+v278<<(uint(int32(2))%32)))) = v297
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if int32(0) < v299 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v367 = v278
	goto L63
L63:
	;
	v371 = v282 + int32(1)
	if v371 != int32(5) {
		v278 = v367
		v282 = v371
		goto L59
	} else {
		goto L74
	}
L64:
	;
	v302 = int32(_a_F_spg_quad_inner_consistent_0)
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0]))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0])) = v305
	v308 = F_palloc(m, int32(32))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L3
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v365 = v363 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v365
	v367 = v365
	goto L63
L67:
	;
	switch v297 {
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
	*(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0])) = v303
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(v344+v345<<(uint(int32(2))%32)))) = v308
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v353 = F_spg_key_orderbys_distances(m, v308, int32(0), v351, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L3
	} else {
		goto L73
	}
L69:
	;
	v334 = *(*float64)(unsafe.Add(mBase, uint32(v19)))
	*(*float64)(unsafe.Add(mBase, uint32(v308))) = v334
	v336 = *(*float64)(unsafe.Add(mBase, uint32(v52)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v308)+8)) = v336
	v338 = *(*float64)(unsafe.Add(mBase, uint32(v52)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v308)+16)) = v338
	v340 = *(*float64)(unsafe.Add(mBase, uint32(v19)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v308)+24)) = v340
	goto L68
L70:
	;
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v308)+8)) = v326
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	*(*int64)(unsafe.Add(mBase, uint32(v308))) = v328
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v274)))
	*(*int64)(unsafe.Add(mBase, uint32(v308)+16)) = v330
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v274)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v308)+24)) = v332
	goto L68
L71:
	;
	v318 = *(*float64)(unsafe.Add(mBase, uint32(v52)))
	*(*float64)(unsafe.Add(mBase, uint32(v308))) = v318
	v320 = *(*float64)(unsafe.Add(mBase, uint32(v19)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v308)+8)) = v320
	v322 = *(*float64)(unsafe.Add(mBase, uint32(v19)))
	*(*float64)(unsafe.Add(mBase, uint32(v308)+16)) = v322
	v324 = *(*float64)(unsafe.Add(mBase, uint32(v52)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v308)+24)) = v324
	goto L68
L72:
	;
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v52)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v308)+8)) = v310
	v312 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
	*(*int64)(unsafe.Add(mBase, uint32(v308))) = v312
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	*(*int64)(unsafe.Add(mBase, uint32(v308)+16)) = v314
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v308)+24)) = v316
	goto L68
L73:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(v355+v356<<(uint(int32(2))%32)))) = v353
	goto L66
L74:
	;
	goto L60
}
func F_spg_quad_picksplit(m *base.Module, l0 int32) int32 {
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
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
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
	var v53 int32
	_ = v53
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	*(*float64)(unsafe.Add(mBase, uint32(v13))) = base.F64_div(v58, base.F64_convert_i32_s(v53))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = base.F64_div(v57, base.F64_convert_i32_s(v62))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v13
	v69 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v69)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v74 = F_palloc(m, v71<<(uint(int32(2))%32))
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
	v53 = v18
	v57 = v21
	v58 = v17
	goto L3
L5:
	;
	goto L6
L6:
	;
	v22 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
	v27 = int32(0)
	v30 = v22
	v31 = v17
	goto L7
L7:
	;
	v33 = v27 << (uint(int32(2)) % 32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33+v34)))
	v37 = *(*float64)(unsafe.Add(mBase, uint32(v36)))
	v38 = base.F64_add(v37, v31)
	*(*float64)(unsafe.Add(mBase, uint32(v13))) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40+v33)))
	v43 = *(*float64)(unsafe.Add(mBase, uint32(v42)+8))
	v44 = base.F64_add(v43, v30)
	*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = v44
	v47 = v27 + int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v47 < v48 {
		v27 = v47
		v30 = v44
		v31 = v38
		goto L7
	} else {
		goto L9
	}
L8:
	;
	v53 = v48
	v57 = v44
	v58 = v38
	goto L3
L9:
	;
	goto L8
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v74
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v80 = F_palloc(m, v77<<(uint(int32(2))%32))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v80
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if int32(0) < v83 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v91 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	return int32(0)
L15:
	;
	v97 = v91 << (uint(int32(2)) % 32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97+v98)))
	v101 = F_getQuadrant(m, v13, v100)
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
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v103+v97))) = v100
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v108 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v106+v97))) = v101 - v108
	v112 = v91 + v108
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v112 < v113 {
		v91 = v112
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
}
func F_spg_range_quad_inner_consistent(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v53 int32
	_ = v53
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
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
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
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(128)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+33)))
	if v22 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 + int32(128)
	return int32(0)
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v25
	v29 = F_palloc(m, v25<<(uint(int32(2))%32))
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
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+34)))
	if v62 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v29
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
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
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v53+v38<<(uint(int32(2))%32)))) = v38
	v59 = v38 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	if v59 < v60 {
		v38 = v59
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
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	v510 = F_palloc(m, v507<<(uint(int32(2))%32))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L5
	} else {
		goto L145
	}
L12:
	;
	v65 = int32(6)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v66 <= int32(0) {
		v492 = v65
		v503 = v2
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	v153 = F_pg_detoast_datum(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L42
	}
L15:
	;
	v69 = v65
	v77 = v2
	goto L16
L16:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v87 = v84 + v77*int32(48)
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+6)))
	if v88 == int32(16) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v492 = v146
	v503 = v2
	goto L11
L18:
	;
	v149 = v77 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v149 < v150 {
		v69 = v146
		v77 = v149
		goto L16
	} else {
		goto L41
	}
L19:
	;
	v492 = int32(0)
	v503 = v2
	goto L11
L20:
	;
	if v142 != 0 {
		v146 = v142
		goto L18
	} else {
		goto L40
	}
L21:
	;
	v142 = v69 & int32(4)
	goto L20
L22:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+44))
	v92 = F_pg_detoast_datum(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v92+int32(base.Ui32(v94)>>(uint(int32(2))%32))-int32(1)))))
	goto L24
L24:
	;
	if base.Ui32(int32(6)) <= base.Ui32(v88-int32(1)) {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L37
	}
L26:
	;
	v142 = v69 & int32(2)
	goto L20
L27:
	;
	if v100&int32(1) == int32(0) {
		goto L21
	} else {
		goto L36
	}
L28:
	;
	if v100&int32(1) == int32(0) {
		v146 = v69
		goto L18
	} else {
		goto L35
	}
L29:
	;
	if v100&int32(1) == int32(0) {
		goto L21
	} else {
		goto L34
	}
L30:
	;
	switch v88 - int32(7) {
	case 0:
		goto L29
	case 1:
		goto L28
	default:
		goto L25
	case 11:
		goto L27
	}
L31:
	;
	goto L32
L32:
	;
	if v100&int32(1) != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	goto L21
L34:
	;
	v146 = v69
	goto L18
L35:
	;
	goto L26
L36:
	;
	goto L26
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v88
	F_errmsg_internal(m, int32(_a_F_spg_range_quad_inner_consistent_0), v18+int32(16))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_spg_range_quad_inner_consistent_1), int32(401), int32(_a_F_spg_range_quad_inner_consistent_2))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	goto L19
L41:
	;
	goto L17
L42:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v156 = F_range_get_typcache(m, l0, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	F_range_deserialize(m, v156, v153, v18+int32(80), v18+int32(72), v18+int32(71))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v166 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v492 = int32(62)
	v503 = v2
	goto L11
L46:
	;
	goto L47
L47:
	;
	v171 = int32(62)
	v179 = v2
	v182 = v2
	goto L48
L48:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v189 = v186 + v179*int32(48)
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189)+6)))
	if v190 == int32(16) {
		goto L65
	} else {
		goto L66
	}
L49:
	;
	v492 = v480
	v503 = v486
	goto L11
L50:
	;
	v489 = v179 + int32(1)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v489 < v490 {
		v171 = v480
		v179 = v489
		v182 = v486
		goto L48
	} else {
		goto L144
	}
L51:
	;
	v492 = int32(0)
	v503 = v477
	goto L11
L52:
	;
	if v470 != 0 {
		v480 = v470
		v486 = v468
		goto L50
	} else {
		goto L143
	}
L53:
	;
	if v436 != 0 {
		goto L125
	} else {
		goto L126
	}
L54:
	;
	v421 = v412 & int32(56)
	v424 = F_range_cmp_bounds(m, v156, v18+int32(80), v416)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L5
	} else {
		goto L115
	}
L55:
	;
	if v406 == int32(0) {
		v431 = v409
		v432 = v403
		v434 = v405
		v436 = v407
		v437 = v408
		goto L53
	} else {
		goto L114
	}
L56:
	;
	v397 = F_range_cmp_bounds(m, v156, v18+int32(80), v388)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L5
	} else {
		goto L110
	}
L57:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+51)))
	if v381 != 0 {
		v477 = v380
		goto L51
	} else {
		goto L108
	}
L58:
	;
	v374 = v365
	v375 = v372
	v376 = v367
	v377 = int32(0)
	v378 = v369
	v379 = v370
	v380 = v371
	goto L57
L59:
	;
	v365 = v359
	v367 = v361
	v369 = v219
	v370 = v219
	v371 = v363
	v372 = int32(0)
	goto L58
L60:
	;
	v365 = v171
	v367 = v219
	v369 = v219
	v370 = v219
	v371 = v182
	v372 = v18 + int32(52)
	goto L58
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L5
	} else {
		goto L105
	}
L62:
	;
	v305 = v18 + int32(120)
	v307 = v18 + int32(112)
	F_range_deserialize(m, v156, v153, v305, v307, v18+int32(111))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L5
	} else {
		goto L92
	}
L63:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+51)))
	if v294 != int32(1) {
		goto L89
	} else {
		goto L90
	}
L64:
	;
	v292 = int32(0)
	v412 = v171 & int32(30)
	v413 = v292
	v415 = v292
	v416 = v18 + int32(60)
	v417 = v18 + int32(52)
	v418 = v182
	goto L54
L65:
	;
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+66)) = uint8(v193)
	v195 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+64)) = uint16(v195)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v189)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v197
	v199 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+58)) = uint8(v199)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+56)) = uint16(v195)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v189)+44))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+51)) = uint8(v199)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v203
	goto L64
L66:
	;
	goto L67
L67:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v189)+44))
	v208 = F_pg_detoast_datum(m, v207)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	v211 = v18 + int32(60)
	F_range_deserialize(m, v156, v208, v211, v18+int32(52), v18+int32(51))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	v218 = int32(1)
	v219 = int32(0)
	switch v190 - v218 {
	case 0:
		v374 = v171
		v375 = v211
		v376 = v219
		v377 = v218
		v378 = v219
		v379 = v219
		v380 = v182
		goto L57
	case 1:
		goto L60
	case 2:
		goto L74
	case 3:
		goto L73
	case 4:
		goto L72
	case 5:
		goto L71
	case 6:
		goto L70
	case 7:
		goto L63
	default:
		goto L61
	case 17:
		goto L62
	}
L70:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+51)))
	if v279 != 0 {
		v480 = v171
		v486 = v182
		goto L50
	} else {
		goto L88
	}
L71:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+51)))
	if v234 != 0 {
		v359 = v171
		v361 = v219
		v363 = v182
		goto L59
	} else {
		goto L75
	}
L72:
	;
	v374 = v171
	v375 = int32(0)
	v376 = v18 + int32(52)
	v377 = v218
	v378 = v219
	v379 = v219
	v380 = v182
	goto L57
L73:
	;
	v359 = v171
	v361 = v18 + int32(60)
	v363 = v182
	goto L59
L74:
	;
	v365 = v171
	v367 = v219
	v369 = v18 + int32(52)
	v370 = v18 + int32(60)
	v371 = v182
	v372 = int32(0)
	goto L58
L75:
	;
	v235 = int32(0)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	if v237 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v239 = v18 + int32(40)
	v241 = v18 + int32(32)
	F_range_deserialize(m, v156, v237, v239, v241, v18+int32(31))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L5
	} else {
		goto L79
	}
L77:
	;
	v246 = v235
	v247 = v235
	goto L78
L78:
	;
	v252 = F_adjacent_inner_consistent(m, v156, v18+int32(60), v18+int32(72), v247)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L5
	} else {
		goto L80
	}
L79:
	;
	v246 = v239
	v247 = v241
	goto L78
L80:
	;
	v259 = F_adjacent_inner_consistent(m, v156, v18+int32(52), v18+int32(80), v246)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	if int32(0) < v259 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v267 = int32(6)
	goto L84
L83:
	;
	v267 = v259 >> (uint(int32(31)) % 32) & int32(24)
	goto L84
L84:
	;
	if int32(0) < v252 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v275 = int32(18)
	goto L87
L86:
	;
	v275 = v252 >> (uint(int32(31)) % 32) & int32(12)
	goto L87
L87:
	;
	v359 = v171 & (v267 | v275)
	v361 = v219
	v363 = int32(1)
	goto L59
L88:
	;
	goto L64
L89:
	;
	v386 = v171
	v387 = v18 + int32(52)
	v388 = v18 + int32(60)
	v389 = int32(0)
	v390 = v219
	v391 = v219
	v392 = v182
	goto L56
L90:
	;
	goto L91
L91:
	;
	v468 = v182
	v470 = v171 & int32(32)
	goto L52
L92:
	;
	v313 = v18 + int32(100)
	v315 = v18 + int32(92)
	F_range_deserialize(m, v156, v208, v313, v315, v18+int32(91))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L5
	} else {
		goto L93
	}
L93:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+91)))
	if v322 != 0 {
		v341 = int32(5)
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v468 = v182
	v470 = v171 & (int32(1) << (uint(v341) % 32))
	goto L52
L95:
	;
	v323 = F_range_cmp_bounds(m, v156, v313, v305)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	v327 = F_range_cmp_bounds(m, v156, v315, v307)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	if int32(0) <= v327 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v331 = int32(1)
	goto L100
L99:
	;
	v331 = int32(2)
	goto L100
L100:
	;
	if int32(0) <= v323 {
		v341 = v331
		goto L94
	} else {
		goto L101
	}
L101:
	;
	if int32(0) <= v327 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v338 = int32(4)
	goto L104
L103:
	;
	v338 = int32(3)
	goto L104
L104:
	;
	v341 = v338
	goto L94
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v190
	F_errmsg_internal(m, int32(_a_F_spg_range_quad_inner_consistent_0), v18)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_spg_range_quad_inner_consistent_1), int32(651), int32(_a_F_spg_range_quad_inner_consistent_2))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	v383 = v374 & int32(30)
	if v376 == int32(0) {
		v403 = v375
		v405 = v377
		v406 = v378
		v407 = v379
		v408 = v380
		v409 = v383
		goto L55
	} else {
		goto L109
	}
L109:
	;
	v386 = v383
	v387 = v375
	v388 = v376
	v389 = v377
	v390 = v378
	v391 = v379
	v392 = v380
	goto L56
L110:
	;
	if v397 <= int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v401 = v386 & int32(38)
	goto L113
L112:
	;
	v401 = v386
	goto L113
L113:
	;
	v403 = v387
	v405 = v389
	v406 = v390
	v407 = v391
	v408 = v392
	v409 = v401
	goto L55
L114:
	;
	v412 = v409
	v413 = v403
	v415 = v405
	v416 = v406
	v417 = v407
	v418 = v408
	goto L54
L115:
	;
	if v424 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v426 = v412
	goto L118
L117:
	;
	v426 = v421
	goto L118
L118:
	;
	if v415 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v427 = v426
	goto L121
L120:
	;
	v427 = v412
	goto L121
L121:
	;
	if int32(0) < v424 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v430 = v421
	goto L124
L123:
	;
	v430 = v427
	goto L124
L124:
	;
	v431 = v430
	v432 = v413
	v434 = v415
	v436 = v417
	v437 = v418
	goto L53
L125:
	;
	v443 = F_range_cmp_bounds(m, v156, v18+int32(72), v436)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L5
	} else {
		goto L128
	}
L126:
	;
	v448 = v431
	goto L127
L127:
	;
	if v432 == int32(0) {
		v468 = v437
		v470 = v448
		goto L52
	} else {
		goto L132
	}
L128:
	;
	if v443 <= int32(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v447 = v431 & int32(50)
	goto L131
L130:
	;
	v447 = v431
	goto L131
L131:
	;
	v448 = v447
	goto L127
L132:
	;
	v452 = v448 & int32(44)
	v455 = F_range_cmp_bounds(m, v156, v18+int32(72), v432)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	if v455 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v457 = v448
	goto L136
L135:
	;
	v457 = v452
	goto L136
L136:
	;
	if v434 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v458 = v457
	goto L139
L138:
	;
	v458 = v448
	goto L139
L139:
	;
	if int32(0) < v455 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v461 = v452
	goto L142
L141:
	;
	v461 = v458
	goto L142
L142:
	;
	v468 = v437
	v470 = v461
	goto L52
L143:
	;
	v477 = v468
	goto L51
L144:
	;
	goto L49
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v510
	if v503&int32(1) != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	v518 = F_palloc(m, v515<<(uint(int32(2))%32))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L5
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v521 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v521
	v523 = int32(_a_F_spg_range_quad_inner_consistent_3)
	v524 = *(*int32)(unsafe.Add(mBase, _c_F_spg_range_quad_inner_consistent[0]))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_spg_range_quad_inner_consistent[0])) = v526
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	if v521 < v528 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v518
	goto L148
L150:
	;
	v535 = v528
	v540 = int32(1)
	goto L153
L151:
	;
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spg_range_quad_inner_consistent[0])) = v524
	goto L1
L153:
	;
	if int32(base.Ui32(v492)>>(uint(v540)%32))&int32(1) != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	goto L152
L155:
	;
	if v503&int32(1) != 0 {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	v577 = v535
	goto L157
L157:
	;
	v579 = v540 + int32(1)
	if v579 <= v577 {
		v535 = v577
		v540 = v579
		goto L153
	} else {
		goto L162
	}
L158:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	v555 = F_datumCopy(m, v552, int32(0), int32(-1))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L5
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v569 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v564+v565<<(uint(int32(2))%32)))) = v540 - v569
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v572 + v569
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	v577 = v576
	goto L157
L161:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v557+v558<<(uint(int32(2))%32)))) = v555
	goto L160
L162:
	;
	goto L154
}
func F_spg_text_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v390 int32
	_ = v390
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if base.B2i32(v53 < int32(2))|base.B2i32(v52 <= int32(0)) != 0 {
		v204 = v52
		goto L14
	} else {
		goto L15
	}
L2:
	;
	return int32(0)
L3:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v23 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v29 == int32(18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v40 = int32(1)
	if v23&v40 != 0 {
		v52 = int32(base.Ui32(v23)>>(uint(v40)%32)) - v40
		goto L1
	} else {
		goto L13
	}
L7:
	;
	v32 = int32(16)
	goto L9
L8:
	;
	v32 = int32(0)
	goto L9
L9:
	;
	if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v39 = int32(4)
	goto L12
L11:
	;
	v39 = v32
	goto L12
L12:
	;
	v52 = v39
	goto L1
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L14:
	;
	v216 = int32(3964)
	if v216 <= v204 {
		goto L61
	} else {
		goto L62
	}
L15:
	;
	v61 = int32(1)
	v62 = v19 + v61
	v66 = v52
	v72 = v61
	goto L16
L16:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v72<<(uint(int32(2))%32))))
	v83 = F_pg_detoast_datum_packed(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L2
	} else {
		goto L18
	}
L17:
	;
	v204 = v195
	goto L14
L18:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v87 = int32(1)
	v88 = v86 & v87
	v90 = v83 + v87
	v92 = v85 & v87
	v93 = int32(0)
	if v86 == v87 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if v184 < v66 {
		goto L56
	} else {
		goto L57
	}
L20:
	;
	if v85 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L21:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v99 == int32(18) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v110 = int32(1)
	if v88 != 0 {
		v120 = int32(base.Ui32(v86)>>(uint(v110)%32)) - v110
		goto L20
	} else {
		goto L30
	}
L24:
	;
	v102 = int32(16)
	goto L26
L25:
	;
	v102 = int32(0)
	goto L26
L26:
	;
	if base.Ui32((v99-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v109 = int32(4)
	goto L29
L28:
	;
	v109 = v102
	goto L29
L29:
	;
	v120 = v109
	goto L20
L30:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v120 = int32(base.Ui32(v114)>>(uint(int32(2))%32)) - int32(4)
	goto L20
L31:
	;
	if v120 < v147 {
		goto L42
	} else {
		goto L43
	}
L32:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v126 == int32(18) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v137 = int32(1)
	if v92 != 0 {
		v147 = int32(base.Ui32(v85)>>(uint(v137)%32)) - v137
		goto L31
	} else {
		goto L41
	}
L35:
	;
	v129 = int32(16)
	goto L37
L36:
	;
	v129 = int32(0)
	goto L37
L37:
	;
	if base.Ui32((v126-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v136 = int32(4)
	goto L40
L39:
	;
	v136 = v129
	goto L40
L40:
	;
	v147 = v136
	goto L31
L41:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v147 = int32(base.Ui32(v141)>>(uint(int32(2))%32)) - int32(4)
	goto L31
L42:
	;
	v149 = v120
	goto L44
L43:
	;
	v149 = v147
	goto L44
L44:
	;
	if v149 <= int32(0) {
		v184 = v93
		goto L19
	} else {
		goto L45
	}
L45:
	;
	if v92 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v154 = v90
	goto L48
L47:
	;
	v154 = v83 + int32(4)
	goto L48
L48:
	;
	if v88 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v155 = v62
	goto L51
L50:
	;
	v155 = v19 + int32(4)
	goto L51
L51:
	;
	v156 = v154
	v159 = v155
	v160 = v93
	goto L52
L52:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v170 != v171 {
		v184 = v160
		goto L19
	} else {
		goto L54
	}
L53:
	;
	v184 = v149
	goto L19
L54:
	;
	v173 = int32(1)
	v178 = v160 + v173
	if v178 != v149 {
		v156 = v156 + v173
		v159 = v159 + v173
		v160 = v178
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v195 = v184
	goto L58
L57:
	;
	v195 = v66
	goto L58
L58:
	;
	v197 = v72 + int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v198 <= v197 {
		v204 = v195
		goto L14
	} else {
		goto L59
	}
L59:
	;
	if int32(0) < v195 {
		v66 = v195
		v72 = v197
		goto L16
	} else {
		goto L60
	}
L60:
	;
	goto L17
L61:
	;
	v219 = v216
	goto L63
L62:
	;
	v219 = v204
	goto L63
L63:
	;
	if v204 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v262 = F_palloc(m, v259*int32(12))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L2
	} else {
		goto L79
	}
L65:
	;
	v222 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v222)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v224 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v224)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v228 = v219 + int32(4)
	v229 = F_palloc(m, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v232 = v219 + int32(1)
	if base.Ui32(v232) <= base.Ui32(int32(127)) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v219 != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v235 = int32(1)
	v238 = v232<<(uint(v235)%32) | v235
	*(*uint8)(unsafe.Add(mBase, uint32(v229))) = uint8(v238)
	v245 = v235
	goto L69
L71:
	;
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229))) = v228 << (uint(int32(2)) % 32)
	v245 = int32(4)
	goto L69
L73:
	;
	v247 = int32(1)
	if v226&v247 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v229
	goto L64
L76:
	;
	v251 = v247
	goto L78
L77:
	;
	v251 = int32(4)
	goto L78
L78:
	;
	base.MemoryCopy(m, v229+v245, v19+v251, v219)
	goto L75
L79:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if int32(0) < v264 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v268 = int32(0)
	goto L83
L81:
	;
	v346 = v264
	goto L82
L82:
	;
	F_pg_qsort(m, v262, v346, int32(12), int32(259))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L2
	} else {
		goto L104
	}
L83:
	;
	v283 = v268 << (uint(int32(2)) % 32)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v283+v284)))
	v287 = F_pg_detoast_datum_packed(m, v286)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L2
	} else {
		goto L86
	}
L84:
	;
	v346 = v341
	goto L82
L85:
	;
	if base.Ui32(v219) < base.Ui32(v318) {
		goto L97
	} else {
		goto L98
	}
L86:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	if v289 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+1)))
	if v295 == int32(18) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v306 = int32(1)
	if v289&v306 != 0 {
		v318 = int32(base.Ui32(v289)>>(uint(v306)%32)) - v306
		goto L85
	} else {
		goto L96
	}
L90:
	;
	v298 = int32(16)
	goto L92
L91:
	;
	v298 = int32(0)
	goto L92
L92:
	;
	if base.Ui32((v295-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v305 = int32(4)
	goto L95
L94:
	;
	v305 = v298
	goto L95
L95:
	;
	v318 = v305
	goto L85
L96:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	v318 = int32(base.Ui32(v312)>>(uint(int32(2))%32)) - int32(4)
	goto L85
L97:
	;
	v321 = int32(1)
	if v289&v321 != 0 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v329 = int32(_a_F_spg_text_picksplit_0)
	goto L99
L99:
	;
	v332 = v262 + v268*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v332)+4)) = v268
	*(*uint16)(unsafe.Add(mBase, uint32(v332)+8)) = uint16(v329)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v335+v283)))
	*(*int32)(unsafe.Add(mBase, uint32(v332))) = v337
	v340 = v268 + int32(1)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v340 < v341 {
		v268 = v340
		goto L83
	} else {
		goto L103
	}
L100:
	;
	v325 = v321
	goto L102
L101:
	;
	v325 = int32(4)
	goto L102
L102:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287+v325+v219))))
	v329 = v328
	goto L99
L103:
	;
	goto L84
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(0)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v366 = F_palloc(m, v363<<(uint(int32(2))%32))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v366
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v372 = F_palloc(m, v369<<(uint(int32(2))%32))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L2
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v372
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v378 = F_palloc(m, v375<<(uint(int32(2))%32))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v378
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if int32(0) < v381 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v390 = int32(0)
	goto L111
L109:
	;
	goto L110
L110:
	;
	return int32(0)
L111:
	;
	v403 = v262 + v390*int32(12)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	v405 = F_pg_detoast_datum_packed(m, v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L2
	} else {
		goto L113
	}
L112:
	;
	goto L110
L113:
	;
	v407 = int32(*(*int16)(unsafe.Add(mBase, uint32(v403)+8)))
	if v390 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405))))
	if v424 != int32(1) {
		goto L124
	} else {
		goto L125
	}
L115:
	;
	v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v403-int32(4)))))
	if v410 == v407&int32(_a_F_spg_text_picksplit_0) {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v414+v415<<(uint(int32(2))%32)))) = v407
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v420 + int32(1)
	goto L114
L118:
	;
	goto L117
L119:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	v504 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v502+v503<<(uint(v504)%32)))) = v498
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v514 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v508+v509<<(uint(v504)%32)))) = v513 - v514
	v518 = v390 + v514
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v518 < v519 {
		v390 = v518
		goto L111
	} else {
		goto L145
	}
L120:
	;
	v474 = v472 + (v219 ^ int32(-1))
	v476 = v474 + int32(4)
	v477 = F_palloc(m, v476)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L2
	} else {
		goto L138
	}
L121:
	;
	v472 = v459
	v473 = v405 + v219 + int32(2)
	goto L120
L122:
	;
	v464 = F_palloc(m, int32(4))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L2
	} else {
		goto L137
	}
L123:
	;
	v456 = int32(1)
	v459 = int32(base.Ui32(v424)>>(uint(v456)%32)) - v456
	if base.Ui32(v219) < base.Ui32(v459) {
		goto L121
	} else {
		goto L136
	}
L124:
	;
	if v424&int32(1) != 0 {
		goto L123
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+1)))
	if v441 == int32(18) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	v433 = int32(base.Ui32(v429)>>(uint(int32(2))%32)) - int32(4)
	if base.Ui32(v433) <= base.Ui32(v219) {
		goto L122
	} else {
		goto L128
	}
L128:
	;
	v472 = v433
	v473 = v405 + v219 + int32(5)
	goto L120
L129:
	;
	v444 = int32(16)
	goto L131
L130:
	;
	v444 = int32(0)
	goto L131
L131:
	;
	if base.Ui32((v441-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v451 = int32(4)
	goto L134
L133:
	;
	v451 = v444
	goto L134
L134:
	;
	if base.Ui32(v451) <= base.Ui32(v219) {
		goto L122
	} else {
		goto L135
	}
L135:
	;
	v472 = v451
	v473 = v405 + v219 + int32(2)
	goto L120
L136:
	;
	goto L122
L137:
	;
	v466 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v464))) = uint8(v466)
	v498 = v464
	goto L119
L138:
	;
	v479 = v472 - v219
	if base.Ui32(v479) <= base.Ui32(int32(127)) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if v474 == int32(0) {
		v498 = v477
		goto L119
	} else {
		goto L144
	}
L140:
	;
	v482 = int32(1)
	v485 = v479<<(uint(v482)%32) | v482
	*(*uint8)(unsafe.Add(mBase, uint32(v477))) = uint8(v485)
	if v474 != 0 {
		v492 = v482
		goto L139
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v477))) = v476 << (uint(int32(2)) % 32)
	v492 = int32(4)
	goto L139
L143:
	;
	v498 = v477
	goto L119
L144:
	;
	base.MemoryCopy(m, v477+v492, v473, v474)
	v498 = v477
	goto L119
L145:
	;
	goto L112
}
