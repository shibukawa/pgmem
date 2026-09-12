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
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v22 int32
	_ = v22
	var v25 float64
	_ = v25
	var v26 float64
	_ = v26
	var v28 int32
	_ = v28
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v34 int32
	_ = v34
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
		v18 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
		v19 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		v22 = base.F64_gt(v18, v19) << (uint(int32(3)) % 32)
		v25 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		v26 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
		if base.F64_gt(v25, v26) != 0 {
			v28 = v22 | int32(4)
		} else {
			v28 = v22
		}
		v31 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
		v32 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
		if base.F64_gt(v31, v32) != 0 {
			v34 = v28 | int32(2)
		} else {
			v34 = v28
		}
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = (base.F64_gt(v15, v16) | v34) & int32(255)
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
	var v88 int32
	_ = v88
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
	var v167 int32
	_ = v167
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
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
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
						v88 = v84
					} else {
						v88 = v42
					}
					F_pg_qsort(m, v23, v88, int32(8), int32(1317))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
						F_pg_qsort(m, v30, v107, int32(8), int32(1317))
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return int32(0)
						} else {
							v112 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
							F_pg_qsort(m, v35, v112, int32(8), int32(1317))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
								F_pg_qsort(m, v40, v117, int32(8), int32(1317))
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
													v167 = int32(0)
													for {
														v181 = v167 << (uint(int32(2)) % 32)
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
														v201 = base.F64_gt(v187, v191) << (uint(int32(3)) % 32)
														if base.F64_gt(v188, v192) != 0 {
															v205 = v201 | int32(4)
														} else {
															v205 = v201
														}
														if base.F64_gt(v186, v190) != 0 {
															v209 = v205 | int32(2)
														} else {
															v209 = v205
														}
														*(*int32)(unsafe.Add(mBase, uint32(v196+v181))) = (base.F64_gt(v185, v189) | v209) & int32(255)
														v215 = v167 + int32(1)
														v216 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
														if v215 < v216 {
															v167 = v215
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v44 int64
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 float64
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 float64
	_ = v193
	var v195 int32
	_ = v195
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int64
	_ = v315
	var v317 int64
	_ = v317
	var v319 int64
	_ = v319
	var v321 int64
	_ = v321
	var v323 float64
	_ = v323
	var v325 float64
	_ = v325
	var v327 float64
	_ = v327
	var v329 float64
	_ = v329
	var v331 int64
	_ = v331
	var v333 int64
	_ = v333
	var v335 int64
	_ = v335
	var v337 int64
	_ = v337
	var v339 float64
	_ = v339
	var v341 float64
	_ = v341
	var v343 float64
	_ = v343
	var v345 float64
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v20 <= v2 {
		v51 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+33)))
	if v52 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v26 = F_palloc(m, v23<<(uint(int32(2))%32))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v26
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v34 = F_palloc(m, v31<<(uint(int32(2))%32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v34
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	if v37 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v40 = int64(-4503599627370496)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v14)+48)) = v40
	v44 = int64(9218868437227405312)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v44
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v44
	v51 = v12 + int32(-32)
	goto L1
L7:
	;
	goto L8
L8:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v51 = v50
	goto L1
L9:
	;
	m.G0 = v14 - int32(-64)
	return int32(0)
L10:
	;
	v259 = F_palloc(m, int32(16))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L3
	} else {
		goto L57
	}
L11:
	;
	v120 = v2
	v122 = int32(30)
	goto L26
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if int32(0) < v55 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v59
	v63 = F_palloc(m, v59<<(uint(int32(2))%32))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L3
	} else {
		goto L16
	}
L15:
	;
	v253 = int32(30)
	goto L10
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v63
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	if v66 <= int32(0) {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v74 = int32(0)
	goto L18
L18:
	;
	v82 = v74 << (uint(int32(2)) % 32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v82+v83))) = v74
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if int32(0) < v86 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L9
L20:
	;
	v89 = int32(_a_F_spg_quad_inner_consistent_0)
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0])) = v92
	v94 = F_box_copy(m, v51)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L3
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v112 = v74 + int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	if v112 < v113 {
		v74 = v112
		goto L18
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0])) = v90
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v98+v82))) = v94
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v104 = F_spg_key_orderbys_distances(m, v94, int32(0), v102, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v106+v82))) = v104
	goto L22
L25:
	;
	goto L19
L26:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v130 = v127 + v120*int32(48)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+44))
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+6)))
	switch v132 - int32(1) {
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
	v253 = v239
	goto L10
L28:
	;
	v244 = v120 + int32(1)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v244 < v245 {
		v120 = v244
		v122 = v239
		goto L26
	} else {
		goto L56
	}
L29:
	;
	v237 = v122 & v236
	if v237 != 0 {
		v239 = v237
		goto L28
	} else {
		goto L55
	}
L30:
	;
	v229 = F_getQuadrant(m, v18, v131)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L3
	} else {
		goto L54
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L3
	} else {
		goto L51
	}
L32:
	;
	v165 = F_DirectFunctionCall2Coll(m, int32(254), int32(0), v131, v18)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L3
	} else {
		goto L45
	}
L33:
	;
	v158 = F_DirectFunctionCall2Coll(m, int32(252), int32(0), v18, v131)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L43
	}
L34:
	;
	v151 = F_DirectFunctionCall2Coll(m, int32(248), int32(0), v18, v131)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L3
	} else {
		goto L41
	}
L35:
	;
	v144 = F_DirectFunctionCall2Coll(m, int32(253), int32(0), v18, v131)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L3
	} else {
		goto L39
	}
L36:
	;
	v137 = F_DirectFunctionCall2Coll(m, int32(250), int32(0), v18, v131)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	if v137 == int32(0) {
		v239 = v122
		goto L28
	} else {
		goto L38
	}
L38:
	;
	v236 = int32(24)
	goto L29
L39:
	;
	if v144 == int32(0) {
		v239 = v122
		goto L28
	} else {
		goto L40
	}
L40:
	;
	v236 = int32(6)
	goto L29
L41:
	;
	if v151 == int32(0) {
		v239 = v122
		goto L28
	} else {
		goto L42
	}
L42:
	;
	v236 = int32(12)
	goto L29
L43:
	;
	if v158 == int32(0) {
		v239 = v122
		goto L28
	} else {
		goto L44
	}
L44:
	;
	v236 = int32(18)
	goto L29
L45:
	;
	if v165 != 0 {
		v239 = v122
		goto L28
	} else {
		goto L46
	}
L46:
	;
	v168 = v12 + int32(-40)
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v131)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v168))) = v169
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v131)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v171
	v175 = F_getQuadrant(m, v18, v12+int32(-48))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	v178 = v131 + int32(8)
	v179 = *(*float64)(unsafe.Add(mBase, uint32(v178)))
	*(*float64)(unsafe.Add(mBase, uint32(v168))) = v179
	v183 = F_getQuadrant(m, v18, v12+int32(-48))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v178)))
	*(*int64)(unsafe.Add(mBase, uint32(v168))) = v185
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v131)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v187
	v191 = F_getQuadrant(m, v18, v12+int32(-48))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v193 = *(*float64)(unsafe.Add(mBase, uint32(v131)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v14)+16)) = v193
	v195 = int32(1)
	v206 = F_getQuadrant(m, v18, v12+int32(-48))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v236 = v195<<(uint(v183)%32) | v195<<(uint(v175)%32) | v195<<(uint(v191)%32) | v195<<(uint(v206)%32)
	goto L29
L51:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214+v120*int32(48))+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v218
	F_errmsg_internal(m, int32(_a_F_spg_quad_inner_consistent_1), v14)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_spg_quad_inner_consistent_2), int32(363), int32(_a_F_spg_quad_inner_consistent_3))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
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
	v236 = int32(1) << (uint(v229) % 32)
	goto L29
L55:
	;
	v253 = int32(0)
	goto L10
L56:
	;
	goto L27
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v259
	v262 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = v262
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = v262
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v267)+8)) = v262
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v262
	v274 = F_palloc(m, int32(16))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v276 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v274
	v280 = v51 + int32(16)
	v284 = v276
	v287 = int32(1)
	goto L59
L59:
	;
	if int32(base.Ui32(v253)>>(uint(v287)%32))&int32(1) != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L9
L61:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v302 = v287 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v297+v284<<(uint(int32(2))%32)))) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if int32(0) < v304 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v372 = v284
	goto L63
L63:
	;
	v376 = v287 + int32(1)
	if v376 != int32(5) {
		v284 = v372
		v287 = v376
		goto L59
	} else {
		goto L74
	}
L64:
	;
	v307 = int32(_a_F_spg_quad_inner_consistent_0)
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0]))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0])) = v310
	v313 = F_palloc(m, int32(32))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L3
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v370 = v368 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v370
	v372 = v370
	goto L63
L67:
	;
	switch v302 {
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
	*(*int32)(unsafe.Add(mBase, _c_F_spg_quad_inner_consistent[0])) = v308
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v349+v350<<(uint(int32(2))%32)))) = v313
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v358 = F_spg_key_orderbys_distances(m, v313, int32(0), v356, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L3
	} else {
		goto L73
	}
L69:
	;
	v339 = *(*float64)(unsafe.Add(mBase, uint32(v18)))
	*(*float64)(unsafe.Add(mBase, uint32(v313))) = v339
	v341 = *(*float64)(unsafe.Add(mBase, uint32(v51)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v313)+8)) = v341
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v51)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v313)+16)) = v343
	v345 = *(*float64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v313)+24)) = v345
	goto L68
L70:
	;
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v313))) = v331
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v313)+8)) = v333
	v335 = *(*int64)(unsafe.Add(mBase, uint32(v280)))
	*(*int64)(unsafe.Add(mBase, uint32(v313)+16)) = v335
	v337 = *(*int64)(unsafe.Add(mBase, uint32(v280)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v313)+24)) = v337
	goto L68
L71:
	;
	v323 = *(*float64)(unsafe.Add(mBase, uint32(v51)))
	*(*float64)(unsafe.Add(mBase, uint32(v313))) = v323
	v325 = *(*float64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v313)+8)) = v325
	v327 = *(*float64)(unsafe.Add(mBase, uint32(v18)))
	*(*float64)(unsafe.Add(mBase, uint32(v313)+16)) = v327
	v329 = *(*float64)(unsafe.Add(mBase, uint32(v51)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v313)+24)) = v329
	goto L68
L72:
	;
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
	*(*int64)(unsafe.Add(mBase, uint32(v313))) = v315
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v313)+8)) = v317
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v313)+16)) = v319
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v313)+24)) = v321
	goto L68
L73:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v360+v361<<(uint(int32(2))%32)))) = v358
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
	var v26 int32
	_ = v26
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
	var v54 int32
	_ = v54
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
	var v90 int32
	_ = v90
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
	*(*float64)(unsafe.Add(mBase, uint32(v13))) = base.F64_div(v58, base.F64_convert_i32_s(v54))
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
	v54 = v18
	v57 = v21
	v58 = v17
	goto L3
L5:
	;
	goto L6
L6:
	;
	v22 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
	v26 = int32(0)
	v30 = v22
	v31 = v17
	goto L7
L7:
	;
	v33 = v26 << (uint(int32(2)) % 32)
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
	v47 = v26 + int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v47 < v48 {
		v26 = v47
		v30 = v44
		v31 = v38
		goto L7
	} else {
		goto L9
	}
L8:
	;
	v54 = v48
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
	v90 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	return int32(0)
L15:
	;
	v97 = v90 << (uint(int32(2)) % 32)
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
	v112 = v90 + v108
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v112 < v113 {
		v90 = v112
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
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
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
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
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
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
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
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
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
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
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
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	v539 = F_palloc(m, v536<<(uint(int32(2))%32))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L5
	} else {
		goto L146
	}
L12:
	;
	v65 = int32(6)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v66 <= int32(0) {
		v521 = v65
		v528 = v2
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
	v521 = v146
	v528 = v2
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
	v521 = int32(0)
	v528 = v2
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
	v521 = int32(62)
	v528 = v2
	goto L11
L46:
	;
	goto L47
L47:
	;
	v171 = int32(62)
	v178 = v2
	v179 = v2
	goto L48
L48:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v189 = v186 + v179*int32(48)
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189)+6)))
	if v190 == int32(16) {
		goto L66
	} else {
		goto L67
	}
L49:
	;
	v521 = v508
	v528 = v511
	goto L11
L50:
	;
	v518 = v179 + int32(1)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v518 < v519 {
		v171 = v508
		v178 = v511
		v179 = v518
		goto L48
	} else {
		goto L145
	}
L51:
	;
	v521 = int32(0)
	v528 = v501
	goto L11
L52:
	;
	if v497 != 0 {
		v508 = v497
		v511 = v491
		goto L50
	} else {
		goto L144
	}
L53:
	;
	if v458 != 0 {
		goto L126
	} else {
		goto L127
	}
L54:
	;
	v446 = v437 & int32(56)
	v449 = F_range_cmp_bounds(m, v156, v18+int32(80), v441)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L5
	} else {
		goto L116
	}
L55:
	;
	if v431 == int32(0) {
		v456 = v434
		v457 = v428
		v458 = v429
		v459 = v430
		v462 = v433
		goto L53
	} else {
		goto L115
	}
L56:
	;
	v422 = F_range_cmp_bounds(m, v156, v18+int32(80), v416)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L5
	} else {
		goto L111
	}
L57:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+51)))
	if v406 != 0 {
		v501 = v402
		goto L51
	} else {
		goto L109
	}
L58:
	;
	v399 = v391
	v400 = v397
	v401 = v393
	v402 = v394
	v403 = v395
	v404 = v396
	v405 = int32(0)
	goto L57
L59:
	;
	v389 = int32(0)
	v391 = v383
	v393 = v389
	v394 = v386
	v395 = v388
	v396 = v387
	v397 = v389
	goto L58
L60:
	;
	v383 = v378
	v386 = v381
	v387 = v223
	v388 = int32(0)
	goto L59
L61:
	;
	v391 = v171
	v393 = v223
	v394 = v178
	v395 = v223
	v396 = v223
	v397 = v18 + int32(52)
	goto L58
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L5
	} else {
		goto L106
	}
L63:
	;
	F_range_deserialize(m, v156, v153, v18+int32(120), v18+int32(112), v18+int32(111))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L5
	} else {
		goto L93
	}
L64:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+51)))
	if v305 != int32(1) {
		goto L90
	} else {
		goto L91
	}
L65:
	;
	v303 = int32(0)
	v437 = v171 & int32(30)
	v438 = v303
	v439 = v18 + int32(52)
	v440 = v178
	v441 = v18 + int32(60)
	v443 = v303
	goto L54
L66:
	;
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+66)) = uint8(v193)
	v195 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+64)) = uint16(v195)
	v198 = v189 + int32(44)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v199
	v201 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+58)) = uint8(v201)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+56)) = uint16(v195)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+51)) = uint8(v201)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v205
	goto L65
L67:
	;
	goto L68
L68:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v189)+44))
	v210 = F_pg_detoast_datum(m, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	F_range_deserialize(m, v156, v210, v18+int32(60), v18+int32(52), v18+int32(51))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	v220 = int32(1)
	v223 = int32(0)
	switch v190 - v220 {
	case 0:
		v399 = v171
		v400 = v18 + int32(60)
		v401 = v223
		v402 = v178
		v403 = v223
		v404 = v223
		v405 = v220
		goto L57
	case 1:
		goto L61
	case 2:
		goto L75
	case 3:
		goto L74
	case 4:
		goto L73
	case 5:
		goto L72
	case 6:
		goto L71
	case 7:
		goto L64
	default:
		goto L62
	case 17:
		goto L63
	}
L71:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+51)))
	if v288&int32(1) != 0 {
		v508 = v171
		v511 = v178
		goto L50
	} else {
		goto L89
	}
L72:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+51)))
	if v239 != 0 {
		v378 = v171
		v381 = v178
		goto L60
	} else {
		goto L76
	}
L73:
	;
	v399 = v171
	v400 = int32(0)
	v401 = v223
	v402 = v178
	v403 = v223
	v404 = v18 + int32(52)
	v405 = v220
	goto L57
L74:
	;
	v383 = v171
	v386 = v178
	v387 = v18 + int32(60)
	v388 = int32(0)
	goto L59
L75:
	;
	v391 = v171
	v393 = v18 + int32(60)
	v394 = v178
	v395 = v18 + int32(52)
	v396 = v223
	v397 = int32(0)
	goto L58
L76:
	;
	v240 = int32(0)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	if v242 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_range_deserialize(m, v156, v242, v18+int32(40), v18+int32(32), v18+int32(31))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L5
	} else {
		goto L80
	}
L78:
	;
	v255 = v240
	v256 = v240
	goto L79
L79:
	;
	v261 = F_adjacent_inner_consistent(m, v156, v18+int32(60), v18+int32(72), v256)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L5
	} else {
		goto L81
	}
L80:
	;
	v255 = v18 + int32(40)
	v256 = v18 + int32(32)
	goto L79
L81:
	;
	v268 = F_adjacent_inner_consistent(m, v156, v18+int32(52), v18+int32(80), v255)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	if int32(0) < v268 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v276 = int32(6)
	goto L85
L84:
	;
	v276 = v268 >> (uint(int32(31)) % 32) & int32(24)
	goto L85
L85:
	;
	if int32(0) < v261 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v284 = int32(18)
	goto L88
L87:
	;
	v284 = v261 >> (uint(int32(31)) % 32) & int32(12)
	goto L88
L88:
	;
	v378 = (v276 | v284) & v171
	v381 = int32(1)
	goto L60
L89:
	;
	goto L65
L90:
	;
	v411 = v171
	v412 = v18 + int32(52)
	v413 = v223
	v414 = v178
	v415 = v223
	v416 = v18 + int32(60)
	v417 = int32(0)
	goto L56
L91:
	;
	goto L92
L92:
	;
	v491 = v178
	v497 = v171 & int32(32)
	goto L52
L93:
	;
	F_range_deserialize(m, v156, v210, v18+int32(100), v18+int32(92), v18+int32(91))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+91)))
	if v333 != 0 {
		v360 = int32(5)
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v491 = v178
	v497 = int32(1) << (uint(v360) % 32) & v171
	goto L52
L96:
	;
	v338 = F_range_cmp_bounds(m, v156, v18+int32(100), v18+int32(120))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	v344 = F_range_cmp_bounds(m, v156, v18+int32(92), v18+int32(112))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	if int32(0) <= v344 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v350 = int32(1)
	goto L101
L100:
	;
	v350 = int32(2)
	goto L101
L101:
	;
	if int32(0) <= v338 {
		v360 = v350
		goto L95
	} else {
		goto L102
	}
L102:
	;
	if int32(0) <= v344 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v357 = int32(4)
	goto L105
L104:
	;
	v357 = int32(3)
	goto L105
L105:
	;
	v360 = v357
	goto L95
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v190
	F_errmsg_internal(m, int32(_a_F_spg_range_quad_inner_consistent_0), v18)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_spg_range_quad_inner_consistent_1), int32(651), int32(_a_F_spg_range_quad_inner_consistent_2))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	v408 = v399 & int32(30)
	if v404 == int32(0) {
		v428 = v400
		v429 = v401
		v430 = v402
		v431 = v403
		v433 = v405
		v434 = v408
		goto L55
	} else {
		goto L110
	}
L110:
	;
	v411 = v408
	v412 = v400
	v413 = v401
	v414 = v402
	v415 = v403
	v416 = v404
	v417 = v405
	goto L56
L111:
	;
	if v422 <= int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v426 = v411 & int32(38)
	goto L114
L113:
	;
	v426 = v411
	goto L114
L114:
	;
	v428 = v412
	v429 = v413
	v430 = v414
	v431 = v415
	v433 = v417
	v434 = v426
	goto L55
L115:
	;
	v437 = v434
	v438 = v428
	v439 = v429
	v440 = v430
	v441 = v431
	v443 = v433
	goto L54
L116:
	;
	if v449 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v451 = v437
	goto L119
L118:
	;
	v451 = v446
	goto L119
L119:
	;
	if v443 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v452 = v451
	goto L122
L121:
	;
	v452 = v437
	goto L122
L122:
	;
	if int32(0) < v449 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v455 = v446
	goto L125
L124:
	;
	v455 = v452
	goto L125
L125:
	;
	v456 = v455
	v457 = v438
	v458 = v439
	v459 = v440
	v462 = v443
	goto L53
L126:
	;
	v469 = F_range_cmp_bounds(m, v156, v18+int32(72), v458)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L5
	} else {
		goto L129
	}
L127:
	;
	v474 = v456
	goto L128
L128:
	;
	if v457 == int32(0) {
		v491 = v459
		v497 = v474
		goto L52
	} else {
		goto L133
	}
L129:
	;
	if v469 <= int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v473 = v456 & int32(50)
	goto L132
L131:
	;
	v473 = v456
	goto L132
L132:
	;
	v474 = v473
	goto L128
L133:
	;
	v478 = v474 & int32(44)
	v481 = F_range_cmp_bounds(m, v156, v18+int32(72), v457)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L5
	} else {
		goto L134
	}
L134:
	;
	if v481 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v483 = v474
	goto L137
L136:
	;
	v483 = v478
	goto L137
L137:
	;
	if v462 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v484 = v483
	goto L140
L139:
	;
	v484 = v474
	goto L140
L140:
	;
	if int32(0) < v481 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v487 = v478
	goto L143
L142:
	;
	v487 = v484
	goto L143
L143:
	;
	v491 = v459
	v497 = v487
	goto L52
L144:
	;
	v501 = v491
	goto L51
L145:
	;
	goto L49
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v539
	if v528&int32(1) != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	v547 = F_palloc(m, v544<<(uint(int32(2))%32))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L5
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v550 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v550
	v552 = int32(_a_F_spg_range_quad_inner_consistent_3)
	v553 = *(*int32)(unsafe.Add(mBase, _c_F_spg_range_quad_inner_consistent[0]))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_spg_range_quad_inner_consistent[0])) = v555
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	if v550 < v557 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v547
	goto L149
L151:
	;
	v564 = v557
	v569 = int32(1)
	goto L154
L152:
	;
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spg_range_quad_inner_consistent[0])) = v553
	goto L1
L154:
	;
	if int32(base.Ui32(v521)>>(uint(v569)%32))&int32(1) != 0 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	goto L153
L156:
	;
	if v528&int32(1) != 0 {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	v606 = v564
	goto L158
L158:
	;
	v608 = v569 + int32(1)
	if v608 <= v606 {
		v564 = v606
		v569 = v608
		goto L154
	} else {
		goto L163
	}
L159:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	v584 = F_datumCopy(m, v581, int32(0), int32(-1))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L5
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v598 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v593+v594<<(uint(int32(2))%32)))) = v569 - v598
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v601 + v598
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	v606 = v605
	goto L158
L162:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v586+v587<<(uint(int32(2))%32)))) = v584
	goto L161
L163:
	;
	goto L155
}
func F_spg_text_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v20 = F_pg_detoast_datum_packed(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v55 < int32(2) {
		v210 = v54
		goto L14
	} else {
		goto L15
	}
L2:
	;
	return int32(0)
L3:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v24 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = int32(4)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v29&int32(254) == int32(2) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v42 = int32(1)
	if v24&v42 != 0 {
		v54 = int32(base.Ui32(v24)>>(uint(v42)%32)) - v42
		goto L1
	} else {
		goto L13
	}
L7:
	;
	v38 = v27
	goto L9
L8:
	;
	v38 = base.B2i32(v29 == int32(18)) << (uint(v27) % 32)
	goto L9
L9:
	;
	if v29 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v41 = v27
	goto L12
L11:
	;
	v41 = v38
	goto L12
L12:
	;
	v54 = v41
	goto L1
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L14:
	;
	v223 = int32(3964)
	if v223 <= v210 {
		goto L62
	} else {
		goto L63
	}
L15:
	;
	if v54 <= int32(0) {
		v210 = v54
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v62 = int32(1)
	v63 = v20 + v62
	v67 = v54
	v73 = v62
	goto L17
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v73<<(uint(int32(2))%32))))
	v85 = F_pg_detoast_datum_packed(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L19
	}
L18:
	;
	v210 = v201
	goto L14
L19:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v89 = int32(1)
	v90 = v88 & v89
	v92 = v85 + v89
	v94 = v87 & v89
	v95 = int32(0)
	if v88 == v89 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v189 < v67 {
		goto L57
	} else {
		goto L58
	}
L21:
	;
	if v87 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L22:
	;
	v98 = int32(4)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v100&int32(254) == int32(2) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v113 = int32(1)
	if v90 != 0 {
		v123 = int32(base.Ui32(v88)>>(uint(v113)%32)) - v113
		goto L21
	} else {
		goto L31
	}
L25:
	;
	v109 = v98
	goto L27
L26:
	;
	v109 = base.B2i32(v100 == int32(18)) << (uint(v98) % 32)
	goto L27
L27:
	;
	if v100 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v112 = v98
	goto L30
L29:
	;
	v112 = v109
	goto L30
L30:
	;
	v123 = v112
	goto L21
L31:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v123 = int32(base.Ui32(v117)>>(uint(int32(2))%32)) - int32(4)
	goto L21
L32:
	;
	if v123 < v151 {
		goto L43
	} else {
		goto L44
	}
L33:
	;
	v126 = int32(4)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v128&int32(254) == int32(2) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v141 = int32(1)
	if v94 != 0 {
		v151 = int32(base.Ui32(v87)>>(uint(v141)%32)) - v141
		goto L32
	} else {
		goto L42
	}
L36:
	;
	v137 = v126
	goto L38
L37:
	;
	v137 = base.B2i32(v128 == int32(18)) << (uint(v126) % 32)
	goto L38
L38:
	;
	if v128 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v140 = v126
	goto L41
L40:
	;
	v140 = v137
	goto L41
L41:
	;
	v151 = v140
	goto L32
L42:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v151 = int32(base.Ui32(v145)>>(uint(int32(2))%32)) - int32(4)
	goto L32
L43:
	;
	v153 = v123
	goto L45
L44:
	;
	v153 = v151
	goto L45
L45:
	;
	if v153 <= int32(0) {
		v189 = v95
		goto L20
	} else {
		goto L46
	}
L46:
	;
	if v94 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v158 = v92
	goto L49
L48:
	;
	v158 = v85 + int32(4)
	goto L49
L49:
	;
	if v90 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v159 = v63
	goto L52
L51:
	;
	v159 = v20 + int32(4)
	goto L52
L52:
	;
	v160 = v158
	v161 = v159
	v164 = v95
	goto L53
L53:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if v175 != v176 {
		v189 = v164
		goto L20
	} else {
		goto L55
	}
L54:
	;
	v189 = v153
	goto L20
L55:
	;
	v178 = int32(1)
	v183 = v164 + v178
	if v183 != v153 {
		v160 = v160 + v178
		v161 = v161 + v178
		v164 = v183
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v201 = v189
	goto L59
L58:
	;
	v201 = v67
	goto L59
L59:
	;
	v203 = v73 + int32(1)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v204 <= v203 {
		v210 = v201
		goto L14
	} else {
		goto L60
	}
L60:
	;
	if int32(0) < v201 {
		v67 = v201
		v73 = v203
		goto L17
	} else {
		goto L61
	}
L61:
	;
	goto L18
L62:
	;
	v226 = v223
	goto L64
L63:
	;
	v226 = v210
	goto L64
L64:
	;
	if v210 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v271 = F_palloc(m, v268*int32(12))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L2
	} else {
		goto L81
	}
L66:
	;
	v229 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v229)
	goto L65
L67:
	;
	goto L68
L68:
	;
	v231 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v231)
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v236&v231 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v239 = v231
	goto L71
L70:
	;
	v239 = int32(4)
	goto L71
L71:
	;
	v242 = v226 + int32(4)
	v243 = F_palloc(m, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v246 = v226 + int32(1)
	if base.Ui32(v246) <= base.Ui32(int32(127)) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v226 != 0 {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v249 = int32(1)
	v252 = v246<<(uint(v249)%32) | v249
	*(*uint8)(unsafe.Add(mBase, uint32(v243))) = uint8(v252)
	v258 = v231
	goto L73
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = v242 << (uint(int32(2)) % 32)
	v258 = int32(4)
	goto L73
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v243
	goto L65
L78:
	;
	v260 = F__emscripten_memcpy_bulkmem(m, v243+v258, v20+v239, v226)
	mBase = m.M
	goto L80
L79:
	;
	goto L80
L80:
	;
	goto L77
L81:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if int32(0) < v273 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v277 = int32(0)
	goto L85
L83:
	;
	v355 = v273
	goto L84
L84:
	;
	F_pg_qsort(m, v271, v355, int32(12), int32(259))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L2
	} else {
		goto L106
	}
L85:
	;
	v293 = v277 << (uint(int32(2)) % 32)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v293+v294)))
	v297 = F_pg_detoast_datum_packed(m, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L2
	} else {
		goto L88
	}
L86:
	;
	v355 = v352
	goto L84
L87:
	;
	if base.Ui32(v226) < base.Ui32(v329) {
		goto L99
	} else {
		goto L100
	}
L88:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297))))
	if v299 == int32(1) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v302 = int32(4)
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+1)))
	if v304&int32(254) == int32(2) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v317 = int32(1)
	if v299&v317 != 0 {
		v329 = int32(base.Ui32(v299)>>(uint(v317)%32)) - v317
		goto L87
	} else {
		goto L98
	}
L92:
	;
	v313 = v302
	goto L94
L93:
	;
	v313 = base.B2i32(v304 == int32(18)) << (uint(v302) % 32)
	goto L94
L94:
	;
	if v304 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v316 = v302
	goto L97
L96:
	;
	v316 = v313
	goto L97
L97:
	;
	v329 = v316
	goto L87
L98:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v329 = int32(base.Ui32(v323)>>(uint(int32(2))%32)) - int32(4)
	goto L87
L99:
	;
	v332 = int32(1)
	if v299&v332 != 0 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v340 = int32(_a_F_spg_text_picksplit_0)
	goto L101
L101:
	;
	v343 = v271 + v277*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v343)+4)) = v277
	*(*uint16)(unsafe.Add(mBase, uint32(v343)+8)) = uint16(v340)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v346+v293)))
	*(*int32)(unsafe.Add(mBase, uint32(v343))) = v348
	v351 = v277 + int32(1)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v351 < v352 {
		v277 = v351
		goto L85
	} else {
		goto L105
	}
L102:
	;
	v336 = v332
	goto L104
L103:
	;
	v336 = int32(4)
	goto L104
L104:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297+v336+v226))))
	v340 = v339
	goto L101
L105:
	;
	goto L86
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(0)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v378
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v384 = F_palloc(m, v381<<(uint(int32(2))%32))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L2
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v384
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v390 = F_palloc(m, v387<<(uint(int32(2))%32))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L2
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v390
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if int32(0) < v393 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v400 = int32(0)
	goto L113
L111:
	;
	goto L112
L112:
	;
	return int32(0)
L113:
	;
	v416 = v271 + v400*int32(12)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	v418 = F_pg_detoast_datum_packed(m, v417)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L2
	} else {
		goto L115
	}
L114:
	;
	goto L112
L115:
	;
	v420 = int32(*(*int16)(unsafe.Add(mBase, uint32(v416)+8)))
	if v400 != 0 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418))))
	if v437 != int32(1) {
		goto L126
	} else {
		goto L127
	}
L117:
	;
	v423 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v416-int32(4)))))
	if v423 == v420&int32(_a_F_spg_text_picksplit_0) {
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v427+v428<<(uint(int32(2))%32)))) = v420
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v433 + int32(1)
	goto L116
L120:
	;
	goto L119
L121:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	v521 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v519+v520<<(uint(v521)%32)))) = v514
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v531 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v525+v526<<(uint(v521)%32)))) = v530 - v531
	v535 = v400 + v531
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v535 < v536 {
		v400 = v535
		goto L113
	} else {
		goto L150
	}
L122:
	;
	v491 = v489 + (v226 ^ int32(-1))
	v493 = v491 + int32(4)
	v494 = F_palloc(m, v493)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L2
	} else {
		goto L140
	}
L123:
	;
	v489 = v478
	v490 = v418 + v226 + int32(2)
	goto L122
L124:
	;
	v482 = F_palloc(m, int32(4))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L2
	} else {
		goto L139
	}
L125:
	;
	v475 = int32(1)
	v478 = int32(base.Ui32(v437)>>(uint(v475)%32)) - v475
	if base.Ui32(v226) < base.Ui32(v478) {
		goto L123
	} else {
		goto L138
	}
L126:
	;
	if v437&int32(1) != 0 {
		goto L125
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v456 = int32(4)
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418)+1)))
	if v458&int32(254) == int32(2) {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	if base.Ui32(int32(base.Ui32(v442)>>(uint(int32(2))%32))-int32(4)) <= base.Ui32(v226) {
		goto L124
	} else {
		goto L130
	}
L130:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	v489 = int32(base.Ui32(v448)>>(uint(int32(2))%32)) - int32(4)
	v490 = v418 + v226 + int32(5)
	goto L122
L131:
	;
	v467 = v456
	goto L133
L132:
	;
	v467 = base.B2i32(v458 == int32(18)) << (uint(v456) % 32)
	goto L133
L133:
	;
	if v458 == int32(1) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v470 = v456
	goto L136
L135:
	;
	v470 = v467
	goto L136
L136:
	;
	if base.Ui32(v470) <= base.Ui32(v226) {
		goto L124
	} else {
		goto L137
	}
L137:
	;
	v489 = v470
	v490 = v418 + v226 + int32(2)
	goto L122
L138:
	;
	goto L124
L139:
	;
	v484 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v482))) = uint8(v484)
	v514 = v482
	goto L121
L140:
	;
	v496 = v489 - v226
	if base.Ui32(v496) <= base.Ui32(int32(127)) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	if v491 != 0 {
		goto L147
	} else {
		goto L148
	}
L142:
	;
	v499 = int32(1)
	v502 = v496<<(uint(v499)%32) | v499
	*(*uint8)(unsafe.Add(mBase, uint32(v494))) = uint8(v502)
	if v491 != 0 {
		v509 = v499
		goto L141
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v494))) = v493 << (uint(int32(2)) % 32)
	v509 = int32(4)
	goto L141
L145:
	;
	v514 = v494
	goto L121
L146:
	;
	v514 = v494
	goto L121
L147:
	;
	v511 = F__emscripten_memcpy_bulkmem(m, v509+v494, v490, v491)
	mBase = m.M
	goto L149
L148:
	;
	goto L149
L149:
	;
	goto L146
L150:
	;
	goto L114
}
