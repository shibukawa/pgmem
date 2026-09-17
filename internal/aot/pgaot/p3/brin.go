package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F__brin_parallel_scan_and_build(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 float64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 float64
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 float64
	_ = v67
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v88 float64
	_ = v88
	var v91 float64
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = F_palloc0(m, int32(12))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(-1)
		v20 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v20)
		v22 = F_tuplesort_begin_index_brin(m, l5, v15)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v22
			v25 = F_BuildIndexInfo(m, l4)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
				*(*uint8)(unsafe.Add(mBase, uint32(v25)+121)) = uint8(v27)
				v29 = int32(1)
				v30 = int32(0)
				v37 = F_table_beginscan_parallel(m, l3, l1+int32(96))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l3)+188))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+140))
					v41 = m.T0[v40].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l3, l4, v25, v29, v30, v29, v30, int32(-1), int32(15), l0, v37)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
						if v44 == int32(0) {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v51 = F_brin_form_tuple(m, v47, v48, v43, v12+int32(12))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
								F_tuplesort_putbrintuple(m, v53, v51, v54)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									v57 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
									*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = base.F64_add(v57, float64(1))
									F_pfree(m, v51)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
										F_tuplesort_performsort(m, v64)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											v67 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
											*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = base.F64_add(v41, v67)
											v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
											*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = int32(1)
											if v70 != 0 {
												F_s_lock(m, l1+int32(44), int32(_a_F__brin_parallel_scan_and_build_0), int32(2847), int32(_a_F__brin_parallel_scan_and_build_1))
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return
												} else {
													v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
													*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v80 + int32(1)
													v84 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
													v85 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
													*(*float64)(unsafe.Add(mBase, uint32(l1)+56)) = base.F64_add(v84, v85)
													v88 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = int32(0)
													v91 = *(*float64)(unsafe.Add(mBase, uint32(l1)+64))
													*(*float64)(unsafe.Add(mBase, uint32(l1)+64)) = base.F64_add(v88, v91)
													F_ConditionVariableSignal(m, l1+int32(32))
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return
													} else {
														v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
														F_tuplesort_end(m, v98)
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return
														} else {
															m.G0 = v12 + int32(16)
															return
														}
													}
												}
											} else {
												v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
												*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v80 + int32(1)
												v84 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
												v85 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
												*(*float64)(unsafe.Add(mBase, uint32(l1)+56)) = base.F64_add(v84, v85)
												v88 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = int32(0)
												v91 = *(*float64)(unsafe.Add(mBase, uint32(l1)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l1)+64)) = base.F64_add(v88, v91)
												F_ConditionVariableSignal(m, l1+int32(32))
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return
												} else {
													v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
													F_tuplesort_end(m, v98)
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return
													} else {
														m.G0 = v12 + int32(16)
														return
													}
												}
											}
										}
									}
								}
							}
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							F_tuplesort_performsort(m, v64)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								v67 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
								*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = base.F64_add(v41, v67)
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = int32(1)
								if v70 != 0 {
									F_s_lock(m, l1+int32(44), int32(_a_F__brin_parallel_scan_and_build_0), int32(2847), int32(_a_F__brin_parallel_scan_and_build_1))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return
									} else {
										v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v80 + int32(1)
										v84 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
										v85 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
										*(*float64)(unsafe.Add(mBase, uint32(l1)+56)) = base.F64_add(v84, v85)
										v88 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = int32(0)
										v91 = *(*float64)(unsafe.Add(mBase, uint32(l1)+64))
										*(*float64)(unsafe.Add(mBase, uint32(l1)+64)) = base.F64_add(v88, v91)
										F_ConditionVariableSignal(m, l1+int32(32))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return
										} else {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
											F_tuplesort_end(m, v98)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												m.G0 = v12 + int32(16)
												return
											}
										}
									}
								} else {
									v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v80 + int32(1)
									v84 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
									v85 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
									*(*float64)(unsafe.Add(mBase, uint32(l1)+56)) = base.F64_add(v84, v85)
									v88 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = int32(0)
									v91 = *(*float64)(unsafe.Add(mBase, uint32(l1)+64))
									*(*float64)(unsafe.Add(mBase, uint32(l1)+64)) = base.F64_add(v88, v91)
									F_ConditionVariableSignal(m, l1+int32(32))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
										F_tuplesort_end(m, v98)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											m.G0 = v12 + int32(16)
											return
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
func F_brinGetTupleForHeapBlock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = base.I32_div_u_s(l1, v17)
	v20 = base.I32_div_u_s(v18, int32(1360))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = base.B2i32(base.Ui32(v20) < base.Ui32(v21))
	if v22 == v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v286
L2:
	;
	v25 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l3))) = uint16(v25)
	v286 = v6
	goto L1
L3:
	;
	goto L4
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v20) < base.Ui32(v21) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v31 = v20 + int32(1)
	goto L7
L6:
	;
	v31 = int32(-1)
	goto L7
L7:
	;
	v32 = v17 * v18
	v33 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)) = uint16(v33)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(-1)
	goto L8
L8:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[0]))
	if v50 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v55 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	return int32(0)
L14:
	;
	goto L12
L15:
	;
	F_LockBuffer(m, v88, int32(1))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L13
	} else {
		goto L26
	}
L16:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v85 = F_ReadBuffer(m, v84, v31)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L13
	} else {
		goto L25
	}
L17:
	;
	if v55 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v76 == v31 {
		v88 = v77
		goto L15
	} else {
		goto L22
	}
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[1]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61+(v55^int32(-1))<<(uint(int32(6))%32))+16))
	v76 = v67
	goto L18
L20:
	;
	goto L21
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[2]))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69+v55<<(uint(int32(6))%32)+int32(-64))+16))
	v76 = v75
	goto L18
L22:
	;
	if v77 == int32(0) {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	F_ReleaseBuffer(m, v77)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	goto L16
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v85
	v88 = v85
	goto L15
L26:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v92 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v112 = base.I32_div_u_s(v32, v111)
	v114 = base.I32_rem_u_s(v112, int32(1360))
	v117 = v110 + v114*int32(6)
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+28)))
	if v118 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[3]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96+(v92^int32(-1))<<(uint(int32(2))%32))))
	v110 = v102
	goto L27
L29:
	;
	goto L30
L30:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[4]))
	v110 = v104 + v92<<(uint(int32(13))%32) + int32(-8192)
	goto L27
L31:
	;
	v121 = int32(0)
	F_LockBuffer(m, v92, v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L13
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v126 = v117 + int32(24)
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)))
	if v127 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v286 = v121
	goto L1
L35:
	;
	F_LockBuffer(m, v200, int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L13
	} else {
		goto L79
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L13
	} else {
		goto L75
	}
L37:
	;
	v129 = v15 + int32(8)
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+2)))
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129))))
	v132 = int32(16)
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+2)))
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126))))
	if v130|v131<<(uint(v132)%32) == v135|v136<<(uint(v132)%32) {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	goto L39
L39:
	;
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)) = uint16(v147)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v149
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+2)))
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126))))
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l3))) = uint16(v153)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_LockBuffer(m, v155, int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L13
	} else {
		goto L47
	}
L40:
	;
	if v146 != 0 {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	goto L40
L42:
	;
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+4)))
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+4)))
	if v142 == v143 {
		v146 = int32(1)
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v146 = int32(0)
	goto L41
L45:
	;
	goto L44
L46:
	;
	goto L39
L47:
	;
	v161 = v151 | v152<<(uint(int32(16))%32)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v162 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	F_LockBuffer(m, v195, int32(1))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L13
	} else {
		goto L59
	}
L49:
	;
	v192 = F_ReadBuffer(m, v27, v161)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L13
	} else {
		goto L58
	}
L50:
	;
	if v162 < int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v161 == v183 {
		v195 = v184
		goto L48
	} else {
		goto L55
	}
L52:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[1]))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168+(v162^int32(-1))<<(uint(int32(6))%32))+16))
	v183 = v174
	goto L51
L53:
	;
	goto L54
L54:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[2]))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v176+v162<<(uint(int32(6))%32)+int32(-64))+16))
	v183 = v182
	goto L51
L55:
	;
	if v184 == int32(0) {
		goto L49
	} else {
		goto L56
	}
L56:
	;
	F_ReleaseBuffer(m, v184)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	goto L49
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v192
	v195 = v192
	goto L48
L59:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v200 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218)+16)))
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219+v218)+6)))
	if v221 != int32(_a_F_brinGetTupleForHeapBlock_0) {
		goto L35
	} else {
		goto L64
	}
L61:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[3]))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v204+(v200^int32(-1))<<(uint(int32(2))%32))))
	v218 = v210
	goto L60
L62:
	;
	goto L63
L63:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[4]))
	v218 = v212 + v200<<(uint(int32(13))%32) + int32(-8192)
	goto L60
L64:
	;
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v225) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v233 = int32(base.Ui32(v225+int32(_a_F_brinGetTupleForHeapBlock_1)) >> (uint(int32(2)) % 32))
	goto L67
L66:
	;
	v233 = int32(0)
	goto L67
L67:
	;
	if base.Ui32(v233&int32(_a_F_brinGetTupleForHeapBlock_2)) < base.Ui32(v224) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v237 = int32(0)
	F_LockBuffer(m, v200, v237)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L13
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v218+v224<<(uint(int32(2))%32))+20))
	if v244&int32(_a_F_brinGetTupleForHeapBlock_3) == int32(0) {
		goto L35
	} else {
		goto L72
	}
L71:
	;
	v286 = v237
	goto L1
L72:
	;
	v251 = v218 + v244&int32(_a_F_brinGetTupleForHeapBlock_4)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	if v252 != v32 {
		goto L35
	} else {
		goto L73
	}
L73:
	;
	if l4 == int32(0) {
		v286 = v251
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(base.Ui32(v244) >> (uint(int32(17)) % 32))
	v286 = v251
	goto L1
L75:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L13
	} else {
		goto L76
	}
L76:
	;
	F_errmsg_internal(m, int32(_a_F_brinGetTupleForHeapBlock_5), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_brinGetTupleForHeapBlock_6), int32(259), int32(_a_F_brinGetTupleForHeapBlock_7))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L13
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	goto L8
}
func F_brinSetHeapBlockItemptr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+2)))
	if l0 < int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_brinSetHeapBlockItemptr[0]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v12+(l0^int32(-1))<<(uint(int32(2))%32))))
		v26 = v18
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_brinSetHeapBlockItemptr[1]))
		v26 = v20 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	v27 = base.I32_div_u_s(l2, l1)
	v29 = base.I32_rem_u_s(v27, int32(1360))
	v32 = v26 + v29*int32(6)
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+28)) = uint16(v33)
	if v33 != 0 {
		v36 = v8
	} else {
		v36 = int32(-1)
	}
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+26)) = uint16(v36)
	if v33 != 0 {
		v39 = v7
	} else {
		v39 = int32(-1)
	}
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+24)) = uint16(v39)
	return
}
func F_brin_bloom_summary_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = v7 + int32(16)
		F_initStringInfo(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_appendStringInfoChar(m, v15, int32(123))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+6)))
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = v22
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v21
				F_appendStringInfo(m, v15, int32(_a_F_brin_bloom_summary_out_0), v7)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_appendStringInfoChar(m, v15, int32(125))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
						m.G0 = v7 + int32(32)
						return v31
					}
				}
			}
		}
	}
}
func F_brin_bloom_summary_recv(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13852(m, l0, int32(_a_F_brin_bloom_summary_recv_0), int32(826), int32(_a_F_brin_bloom_summary_recv_1), int32(_a_F_brin_bloom_summary_recv_2), int32(_a_F_brin_bloom_summary_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_brin_bloom_summary_send(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_byteasend(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_brin_desc(m *base.Module, l0 int32, l1 int32) {
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
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	switch int32(base.Ui32(v12)>>(uint(int32(4))%32)) & int32(7) {
	case 0:
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v17
		F_appendStringInfo(m, l0, int32(_a_F_brin_desc_0), v8)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			m.G0 = v8 + int32(96)
			return
		}
	case 1:
		v24 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+8)))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v25
		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v24
		F_appendStringInfo(m, l0, int32(_a_F_brin_desc_1), v8+int32(16))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			m.G0 = v8 + int32(96)
			return
		}
	case 2:
		v33 = *(*int64)(unsafe.Add(mBase, uint32(v11)+4))
		v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
		v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v35
		*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v34
		*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v33
		F_appendStringInfo(m, l0, int32(_a_F_brin_desc_2), v8+int32(32))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			m.G0 = v8 + int32(96)
			return
		}
	case 3:
		v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v44
		F_appendStringInfo(m, l0, int32(_a_F_brin_desc_3), v8+int32(48))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return
		} else {
			m.G0 = v8 + int32(96)
			return
		}
	case 4:
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v51
		F_appendStringInfo(m, l0, int32(_a_F_brin_desc_4), v8-int32(-64))
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return
		} else {
			m.G0 = v8 + int32(96)
			return
		}
	case 5:
		v58 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+8)))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+88)) = v59
		*(*int64)(unsafe.Add(mBase, uint32(v8)+80)) = v58
		F_appendStringInfo(m, l0, int32(_a_F_brin_desc_5), v8+int32(80))
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return
		} else {
			m.G0 = v8 + int32(96)
			return
		}
	default:
		m.G0 = v8 + int32(96)
		return
	}
}
func F_brin_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v318 int32
	_ = v318
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v477 int32
	_ = v477
	var v513 int32
	_ = v513
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v674 int32
	_ = v674
	var v710 int32
	_ = v710
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v753 int32
	_ = v753
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	v5 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+14)) = uint16(v5)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v37 = F_palloc(m, v34<<(uint(int32(2))%32))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v42 = F_palloc0(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v44 = int32(8)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v49 = base.I32_div_s(v45+int32(7), v44)
	v50 = F_palloc(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v55 = F_palloc(m, v52<<(uint(int32(2))%32))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v58 <= int32(0) {
		v336 = v5
		v339 = v44
		v341 = v5
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v350 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L7:
	;
	v62 = l0 + int32(20)
	v71 = v5
	v74 = v5
	v75 = v5
	v83 = v5
	goto L8
L8:
	;
	v94 = l2 + int32(24) + v74*int32(20)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+3)))
	if v95 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v307 = int32(1)
	if v285&v307 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L10:
	;
	v303 = v74 + int32(1)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	if v303 < v305 {
		v71 = v281
		v74 = v303
		v75 = v285
		v83 = v293
		goto L8
	} else {
		goto L52
	}
L11:
	;
	v98 = int32(0)
	v101 = v62 + v74<<(uint(int32(2))%32)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102))))
	if v103 == v98 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+2)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
	if v146 != 0 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v281 = v71
	v285 = int32(1)
	v293 = v83
	goto L10
L15:
	;
	goto L16
L16:
	;
	v113 = v71
	v115 = v98
	goto L17
L17:
	;
	v134 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v113+v42))) = uint8(v134)
	v139 = v113 + v134
	v141 = v115 + v134
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142))))
	if base.Ui32(v141) < base.Ui32(v143) {
		v113 = v139
		v115 = v141
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v281 = v139
	v285 = v134
	v293 = v83
	goto L10
L19:
	;
	goto L18
L20:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	m.T0[v146].(func(*base.Module, int32, int32, int32))(m, l0, v147, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v151 = v145 | v75
	v154 = v62 + v74<<(uint(int32(2))%32)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155))))
	if v156 == int32(0) {
		v281 = v71
		v285 = v151
		v293 = v83
		goto L10
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v166 = v155
	v168 = v71
	v170 = int32(0)
	v180 = v83
	goto L25
L25:
	;
	v190 = v170 << (uint(int32(2)) % 32)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v190+v191)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v166+v190)+8))
	v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+8)))
	if v196 != int32(_a_F_brin_form_tuple_0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v281 = v269
	v285 = v151
	v293 = v261
	goto L10
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37+v168<<(uint(int32(2))%32)))) = v257
	v268 = int32(1)
	v269 = v168 + v268
	v271 = v170 + v268
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v272))))
	if base.Ui32(v271) < base.Ui32(v273) {
		v166 = v272
		v168 = v269
		v170 = v271
		v180 = v261
		goto L25
	} else {
		goto L51
	}
L28:
	;
	v257 = v193
	v261 = v180
	goto L27
L29:
	;
	goto L30
L30:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	v201 = base.B2i32(v199 != int32(1))
	if v199 != int32(1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v208 = base.B2i32(v199 == int32(1))
	if v206&int32(3) != 0 {
		v243 = v205
		v245 = v208
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v205 = v193
	v206 = v199
	goto L31
L33:
	;
	goto L34
L34:
	;
	v202 = F_detoast_external_attr(m, v193)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	v205 = v202
	v206 = v204
	goto L31
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55+v180<<(uint(int32(2))%32)))) = v248
	v257 = v248
	v261 = v180 + int32(1)
	goto L27
L37:
	;
	if v245 == int32(0) {
		v257 = v243
		v261 = v180
		goto L27
	} else {
		goto L50
	}
L38:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	if base.Ui32(v211) < base.Ui32(int32(2044)) {
		v243 = v205
		v245 = v208
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+12)))
	switch v214 - int32(109) {
	case 0, 11:
		goto L40
	default:
		v243 = v205
		v245 = v208
		goto L37
	}
L40:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	v224 = v219 + v220<<(uint(int32(4))%32) + v74*int32(100)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+88))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	if v225 == v226 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+105)))
	v230 = v228
	goto L43
L42:
	;
	v230 = int32(0)
	goto L43
L43:
	;
	v232 = F_toast_compress_datum(m, v205, base.I32_extend8_s(v230))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v232 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v237 = v232
	goto L47
L46:
	;
	v237 = v205
	goto L47
L47:
	;
	if v201|base.B2i32(v232 == int32(0)) != 0 {
		v243 = v237
		v245 = base.B2i32(v199 == int32(1)) | base.B2i32(v232 != int32(0))
		goto L37
	} else {
		goto L48
	}
L48:
	;
	F_pfree(m, v205)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v248 = v232
	goto L36
L50:
	;
	v248 = v243
	goto L36
L51:
	;
	goto L26
L52:
	;
	goto L9
L53:
	;
	v336 = int32(0)
	v339 = v44
	v341 = v293
	goto L6
L54:
	;
	goto L55
L55:
	;
	v318 = base.I32_div_s(v305<<(uint(int32(1))%32)+int32(7), int32(8))
	v336 = v307
	v339 = (v318 + int32(12)) & int32(-8)
	v341 = v293
	goto L6
L56:
	;
	v353 = int32(_a_F_brin_form_tuple_1)
	v354 = *(*int32)(unsafe.Add(mBase, _c_F_brin_form_tuple[0]))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_form_tuple[0])) = v356
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v359 = F_CreateTemplateTupleDesc(m, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	v513 = v350
	goto L58
L58:
	;
	v536 = F_heap_compute_data_size(m, v513, v37, v42)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L73
	}
L59:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	if int32(0) < v362 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v376 = int32(0)
	v377 = int32(1)
	v379 = v362
	goto L63
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_brin_form_tuple[0])) = v354
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v359
	v513 = v359
	goto L58
L63:
	;
	v399 = l0 + int32(20) + v376<<(uint(int32(2))%32)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	v401 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v400))))
	if v401 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L62
L65:
	;
	v407 = v400
	v408 = int32(0)
	v410 = v377
	goto L68
L66:
	;
	v457 = v377
	v459 = v379
	goto L67
L67:
	;
	v477 = v376 + int32(1)
	if v477 < v459 {
		v376 = v477
		v377 = v457
		v379 = v459
		goto L63
	} else {
		goto L72
	}
L68:
	;
	v430 = int32(0)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v407+v408<<(uint(int32(2))%32))+8))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	F_TupleDescInitEntry(m, v359, base.I32_extend16_s(v410), v430, v435, int32(-1), v430)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	v457 = v441
	v459 = v448
	goto L67
L70:
	;
	v440 = int32(1)
	v441 = v410 + v440
	v443 = v408 + v440
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	v445 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v444))))
	if base.Ui32(v443) < base.Ui32(v445) {
		v407 = v444
		v408 = v443
		v410 = v441
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	goto L64
L73:
	;
	v542 = (v536 + v339 + int32(7)) & int32(-8)
	v543 = F_palloc0(m, v542)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v543)+4)) = uint8(v339)
	*(*int32)(unsafe.Add(mBase, uint32(v543))) = l1
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v547 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v550 = int32(_a_F_brin_form_tuple_1)
	v551 = *(*int32)(unsafe.Add(mBase, _c_F_brin_form_tuple[0]))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_form_tuple[0])) = v553
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v556 = F_CreateTemplateTupleDesc(m, v555)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	v710 = v547
	goto L77
L77:
	;
	F_heap_fill_tuple(m, v710, v37, v42, v543+v339, v30+int32(14), v50)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L92
	}
L78:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v558)))
	if int32(0) < v559 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v573 = int32(0)
	v574 = int32(1)
	v576 = v559
	goto L82
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_brin_form_tuple[0])) = v551
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v556
	v710 = v556
	goto L77
L82:
	;
	v596 = l0 + int32(20) + v573<<(uint(int32(2))%32)
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v596)))
	v598 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v597))))
	if v598 != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	goto L81
L84:
	;
	v604 = v597
	v605 = int32(0)
	v607 = v574
	goto L87
L85:
	;
	v654 = v574
	v656 = v576
	goto L86
L86:
	;
	v674 = v573 + int32(1)
	if v674 < v656 {
		v573 = v674
		v574 = v654
		v576 = v656
		goto L82
	} else {
		goto L91
	}
L87:
	;
	v627 = int32(0)
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v604+v605<<(uint(int32(2))%32))+8))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v631)))
	F_TupleDescInitEntry(m, v556, base.I32_extend16_s(v607), v627, v632, int32(-1), v627)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L89
	}
L88:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v644)))
	v654 = v638
	v656 = v645
	goto L86
L89:
	;
	v637 = int32(1)
	v638 = v607 + v637
	v640 = v605 + v637
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v596)))
	v642 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v641))))
	if base.Ui32(v640) < base.Ui32(v642) {
		v604 = v641
		v605 = v640
		v607 = v638
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	goto L83
L92:
	;
	F_pfree(m, v37)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_pfree(m, v42)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_pfree(m, v50)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	if int32(0) < v341 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v753 = int32(0)
	goto L99
L97:
	;
	goto L98
L98:
	;
	v811 = v543 + int32(4)
	if v336 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L99:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v55+v753<<(uint(int32(2))%32))))
	F_pfree(m, v777)
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L1
	} else {
		goto L101
	}
L100:
	;
	goto L98
L101:
	;
	v781 = v753 + int32(1)
	if v781 != v341 {
		v753 = v781
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v962 == int32(1) {
		goto L125
	} else {
		goto L126
	}
L104:
	;
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v811))))
	v816 = v814 | int32(-128)
	*(*uint8)(unsafe.Add(mBase, uint32(v811))) = uint8(v816)
	v818 = int32(0)
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v819)))
	if v820 <= v818 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v828 = v816
	v829 = v811
	v830 = int32(128)
	v832 = v818
	goto L106
L106:
	;
	if v830 != int32(128) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v878 = int32(0)
	if v876 <= v878 {
		goto L103
	} else {
		goto L115
	}
L108:
	;
	v861 = v828
	v862 = v829
	v863 = v830 << (uint(int32(1)) % 32)
	goto L110
L109:
	;
	v855 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v829)+1)) = uint8(v855)
	v858 = int32(1)
	v861 = v855
	v862 = v829 + v858
	v863 = v858
	goto L110
L110:
	;
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v832*int32(20))+27)))
	if v867 == int32(1) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v870 = v861 | v863
	*(*uint8)(unsafe.Add(mBase, uint32(v862))) = uint8(v870)
	v872 = v870
	goto L113
L112:
	;
	v872 = v861
	goto L113
L113:
	;
	v874 = v832 + int32(1)
	v875 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	if v874 < v876 {
		v828 = v872
		v829 = v862
		v830 = v863
		v832 = v874
		goto L106
	} else {
		goto L114
	}
L114:
	;
	goto L107
L115:
	;
	v885 = v872
	v886 = v862
	v887 = v863
	v889 = v878
	goto L116
L116:
	;
	if v887 != int32(128) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L103
L118:
	;
	v918 = v885
	v919 = v886
	v920 = v887 << (uint(int32(1)) % 32)
	goto L120
L119:
	;
	v912 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v886)+1)) = uint8(v912)
	v915 = int32(1)
	v918 = v912
	v919 = v886 + v915
	v920 = v915
	goto L120
L120:
	;
	v924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v889*int32(20))+26)))
	if v924 == int32(1) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v927 = v918 | v920
	*(*uint8)(unsafe.Add(mBase, uint32(v919))) = uint8(v927)
	v929 = v927
	goto L123
L122:
	;
	v929 = v918
	goto L123
L123:
	;
	v931 = v889 + int32(1)
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v932)))
	if v931 < v933 {
		v885 = v929
		v886 = v919
		v887 = v920
		v889 = v931
		goto L116
	} else {
		goto L124
	}
L124:
	;
	goto L117
L125:
	;
	v965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v811))))
	v967 = v965 | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v811))) = uint8(v967)
	goto L127
L126:
	;
	goto L127
L127:
	;
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v969 == int32(1) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v811))))
	v974 = v972 | int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v811))) = uint8(v974)
	goto L130
L129:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v542
	m.G0 = v30 + int32(16)
	return v543
}
func F_brin_getinsertbuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
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
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int64
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v38 = int32(-1)
	goto L3
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v38 = v36
	goto L3
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_brin_getinsertbuffer[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v36 = v27
	goto L4
L6:
	;
	goto L7
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_brin_getinsertbuffer[1]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+l1<<(uint(int32(6))%32)+int32(-64))+16))
	v36 = v35
	goto L4
L8:
	;
	v59 = v48
	goto L16
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	if v40 != int32(-1) {
		v48 = v40
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v44 = F_GetPageWithFreeSpace(m, l0, l2)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	return int32(0)
L14:
	;
	v48 = v44
	goto L8
L15:
	;
	m.G0 = v16 + int32(32)
	return v288
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_brin_getinsertbuffer[2]))
	if v69 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_brin_initialize_empty_new_buffer(m, l0, v113)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L13
	} else {
		goto L102
	}
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L13
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v72 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v72)
	if v59 == int32(-1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L20
L22:
	;
	if base.B2i32(l1 == int32(0))|base.B2i32(base.Ui32(v112) <= base.Ui32(v38)) != 0 {
		goto L39
	} else {
		goto L40
	}
L23:
	;
	v76 = int32(0)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v77 != 0 {
		v83 = v76
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v108 = int32(0)
	if v59 == v38 {
		goto L35
	} else {
		goto L36
	}
L26:
	;
	v85 = F_ReadBuffer(m, l0, int32(-1))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L13
	} else {
		goto L30
	}
L27:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v78 != 0 {
		v83 = v76
		goto L26
	} else {
		goto L28
	}
L28:
	;
	F_LockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v83 = int32(1)
	goto L26
L30:
	;
	if v85 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v106 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v106)
	v112 = v105
	v113 = v85
	v114 = v83
	goto L22
L32:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_brin_getinsertbuffer[0]))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90+(v85^int32(-1))<<(uint(int32(6))%32))+16))
	v105 = v96
	goto L31
L33:
	;
	goto L34
L34:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_brin_getinsertbuffer[1]))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98+v85<<(uint(int32(6))%32)+int32(-64))+16))
	v105 = v104
	goto L31
L35:
	;
	v112 = v38
	v113 = l1
	v114 = v108
	goto L22
L36:
	;
	goto L37
L37:
	;
	v110 = F_ReadBuffer(m, l0, v59)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	v112 = v59
	v113 = v110
	v114 = v108
	goto L22
L39:
	;
	F_LockBuffer(m, v113, int32(2))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L13
	} else {
		goto L59
	}
L40:
	;
	F_LockBuffer(m, l1, int32(2))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L13
	} else {
		goto L41
	}
L41:
	;
	if l1 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+16)))
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134+v133)+6)))
	if v136 == int32(_a_F_brin_getinsertbuffer_0) {
		goto L39
	} else {
		goto L46
	}
L43:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_brin_getinsertbuffer[3]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v125+(l1^int32(-1))<<(uint(int32(2))%32))))
	v133 = v127
	goto L42
L44:
	;
	goto L45
L45:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_brin_getinsertbuffer[4]))
	v133 = v129 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L42
L46:
	;
	F_LockBuffer(m, l1, int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v142 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_brin_initialize_empty_new_buffer(m, l0, v113)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L13
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v114 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	F_UnlockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	F_ReleaseBuffer(m, v113)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L13
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	v152 = int32(0)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v153 != int32(1) {
		v288 = v152
		goto L15
	} else {
		goto L57
	}
L57:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v112, v112+int32(1))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L13
	} else {
		goto L58
	}
L58:
	;
	v160 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v160)
	v288 = v152
	goto L15
L59:
	;
	if v114 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	F_UnlockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L13
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v113 < int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L62
L64:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v187 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_brin_getinsertbuffer[3]))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v172+(v113^int32(-1))<<(uint(int32(2))%32))))
	v186 = v178
	goto L64
L66:
	;
	goto L67
L67:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_brin_getinsertbuffer[4]))
	v186 = v180 + v113<<(uint(int32(13))%32) + int32(-8192)
	goto L64
L68:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v245 != int32(1) {
		goto L90
	} else {
		goto L91
	}
L69:
	;
	v190 = int32(0)
	v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186)+16)))
	v192 = v186 + v191
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v192)+6)))
	if v193 != int32(_a_F_brin_getinsertbuffer_0) {
		v208 = v190
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v212 != 0 {
		goto L80
	} else {
		goto L81
	}
L72:
	;
	if base.Ui32(v208) < base.Ui32(l2) {
		goto L68
	} else {
		goto L79
	}
L73:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+4)))
	if v196&int32(1) != 0 {
		v208 = v190
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v199 = int32(4)
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186)+14)))
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186)+12)))
	v202 = v200 - v201
	if v202 <= v199 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v208 = v205 - int32(4)
	goto L72
L76:
	;
	v205 = v199
	goto L78
L77:
	;
	v205 = v202
	goto L78
L78:
	;
	goto L75
L79:
	;
	goto L71
L80:
	;
	v236 = v212
	goto L82
L81:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v214
	v216 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v216
	v218 = F_smgropen(m, v16, v213)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L13
	} else {
		goto L83
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v236)+16)) = v112
	if base.B2i32(l1 == int32(0))|base.B2i32(base.Ui32(v38) <= base.Ui32(v112)) != 0 {
		v288 = v113
		goto L15
	} else {
		goto L88
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v218
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v218)+72))
	if v222 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v236 = v234
	goto L82
L85:
	;
	v230 = v222
	goto L87
L86:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v218)+76))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v218)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+4)) = v224
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v218)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = v226
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v218)+72))
	v230 = v228
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v218)+72)) = v230 + int32(1)
	goto L84
L88:
	;
	F_LockBuffer(m, l1, int32(2))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L13
	} else {
		goto L89
	}
L89:
	;
	v288 = v113
	goto L15
L90:
	;
	if v112 != v38 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	goto L17
L93:
	;
	F_UnlockReleaseBuffer(m, v113)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L13
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v251 = int32(0)
	if base.B2i32(l1 == v251)|base.B2i32(base.Ui32(v112) < base.Ui32(v38)) == v251 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L95
L97:
	;
	F_LockBuffer(m, l1, int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L13
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v260 = F_RecordAndGetPageWithFreeSpace(m, l0, v112, v208, l2)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L13
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	v59 = v260
	goto L16
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L13
	} else {
		goto L103
	}
L103:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L13
	} else {
		goto L104
	}
L104:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v271 + int32(4)
	F_errmsg(m, int32(_a_F_brin_getinsertbuffer_1), v16+int32(16))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L13
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_brin_getinsertbuffer_2), int32(852), int32(_a_F_brin_getinsertbuffer_3))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L13
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_brin_minmax_multi_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v28 = F_pg_detoast_datum(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L3
	} else {
		goto L59
	}
L2:
	;
	m.G0 = v19 + int32(32)
	return v240
L3:
	;
	return int32(0)
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v33 = F_brin_range_deserialize(m, v32, v28)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if int32(0) < v35 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if v22 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	if v153 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L9:
	;
	v240 = int32(1)
	goto L2
L10:
	;
	goto L11
L11:
	;
	v53 = int32(0)
	goto L12
L12:
	;
	v61 = v33 + int32(36) + v53<<(uint(int32(3))%32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v67 = int32(0)
	goto L14
L13:
	;
	goto L8
L14:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v23+v67<<(uint(int32(2))%32))))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+44))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+4)))
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+6)))
	switch v88 - int32(1) {
	case 0, 1:
		goto L21
	case 2:
		goto L20
	case 3, 4:
		goto L18
	default:
		goto L19
	}
L15:
	;
	v134 = v53 + int32(1)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v134 < v135 {
		v53 = v134
		goto L12
	} else {
		goto L38
	}
L16:
	;
	goto L15
L17:
	;
	v129 = int32(1)
	v131 = v67 + v129
	if v22 != v131 {
		v67 = v131
		goto L14
	} else {
		goto L37
	}
L18:
	;
	v123 = F_minmax_multi_get_strategy_procinfo(m, v24, v87, v86, v88)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L3
	} else {
		goto L34
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L3
	} else {
		goto L31
	}
L20:
	;
	v98 = F_minmax_multi_get_strategy_procinfo(m, v24, v87, v86, int32(5))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L3
	} else {
		goto L25
	}
L21:
	;
	v91 = F_minmax_multi_get_strategy_procinfo(m, v24, v87, v86, v88)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v93 = F_FunctionCall2Coll(m, v91, v21, v62, v85)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	if v93 == int32(0) {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	goto L17
L25:
	;
	v100 = F_FunctionCall2Coll(m, v98, v21, v62, v85)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	if v100 != 0 {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	v103 = F_minmax_multi_get_strategy_procinfo(m, v24, v87, v86, int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v105 = F_FunctionCall2Coll(m, v103, v21, v63, v85)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	if v105 == int32(0) {
		goto L17
	} else {
		goto L30
	}
L30:
	;
	goto L16
L31:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v113
	F_errmsg_internal(m, int32(_a_F_brin_minmax_multi_consistent_0), v19)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_brin_minmax_multi_consistent_1), int32(2650), int32(_a_F_brin_minmax_multi_consistent_2))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v125 = F_FunctionCall2Coll(m, v123, v21, v63, v85)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	if v125 == int32(0) {
		goto L16
	} else {
		goto L36
	}
L36:
	;
	goto L17
L37:
	;
	v240 = v129
	goto L2
L38:
	;
	goto L13
L39:
	;
	v240 = int32(0)
	goto L2
L40:
	;
	goto L41
L41:
	;
	v157 = int32(0)
	if v22 <= v157 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v240 = int32(1)
	goto L2
L43:
	;
	goto L44
L44:
	;
	v164 = v157
	goto L45
L45:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(36)+v179<<(uint(int32(3))%32)+v164<<(uint(int32(2))%32))))
	v190 = int32(0)
	goto L47
L46:
	;
	v240 = int32(0)
	goto L2
L47:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v23+v190<<(uint(int32(2))%32))))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v208&int32(1) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v237 = v164 + int32(1)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	if v237 < v238 {
		v164 = v237
		goto L45
	} else {
		goto L58
	}
L49:
	;
	goto L48
L50:
	;
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207)+6)))
	if base.Ui32(int32(4)) < base.Ui32((v213-int32(1))&int32(_a_F_brin_minmax_multi_consistent_3)) {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v231 = int32(1)
	v233 = v190 + v231
	if v233 != v22 {
		v190 = v233
		goto L47
	} else {
		goto L57
	}
L53:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v207)+44))
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207)+4)))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v207)+8))
	v223 = F_minmax_multi_get_strategy_procinfo(m, v24, v221, v222, v213)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v225 = F_FunctionCall2Coll(m, v223, v21, v186, v220)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	if v225 == int32(0) {
		goto L49
	} else {
		goto L56
	}
L56:
	;
	goto L52
L57:
	;
	v240 = v231
	goto L2
L58:
	;
	goto L46
L59:
	;
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v264
	F_errmsg_internal(m, int32(_a_F_brin_minmax_multi_consistent_0), v19+int32(16))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_brin_minmax_multi_consistent_1), int32(2709), int32(_a_F_brin_minmax_multi_consistent_2))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L3
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_brin_minmax_multi_distance_float8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v8 = int64(9223372036854775807)
	v9 = base.I64_reinterpret_f64(v6) & v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v11)&v8) {
		if base.Ui64(v9) <= base.Ui64(int64(9218868437227405312)) {
			v24 = F_Float8GetDatum(m, math.Float64frombits(uint64(0x7ff0000000000000)))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				return v24
			}
		} else {
			v30 = float64(0)
			v31 = F_Float8GetDatum(m, v30)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				return v31
			}
		}
	} else {
		if base.Ui64(v9) < base.Ui64(int64(9218868437227405313)) {
			v30 = base.F64_sub(v6, v11)
			v31 = F_Float8GetDatum(m, v30)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				return v31
			}
		} else {
			v24 = F_Float8GetDatum(m, math.Float64frombits(uint64(0x7ff0000000000000)))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				return v24
			}
		}
	}
}
func F_brin_minmax_multi_distance_inet(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
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
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v182 float64
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v194 float64
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v211 float64
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v21 = int32(1)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v23&v21 != 0 {
				v26 = v21
			} else {
				v26 = int32(4)
			}
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v26))))
			v29 = int32(1)
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
			if v31&v29 != 0 {
				v34 = v29
			} else {
				v34 = int32(4)
			}
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v34))))
			if v28 == v36 {
				v38 = int32(4)
				v39 = v13 + v38
				v41 = v13 + int32(1)
				if v28 == int32(2) {
					v46 = v38
				} else {
					v46 = int32(16)
				}
				v47 = F_palloc(m, v46)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v49 = int32(4)
					v51 = int32(1)
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
					v55 = v53 & v51
					if v55 != 0 {
						v56 = v51
					} else {
						v56 = v49
					}
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v56))))
					if v58 == int32(2) {
						v61 = v49
					} else {
						v61 = int32(16)
					}
					if v61 != 0 {
						if v55 != 0 {
							v62 = v41
						} else {
							v62 = v39
						}
						base.MemoryCopy(m, v47, v62+int32(2), v61)
					} else {
					}
					v66 = int32(4)
					v67 = v18 + v66
					v68 = int32(1)
					v69 = v18 + v68
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
					if v74&v68 != 0 {
						v77 = v68
					} else {
						v77 = v66
					}
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v77))))
					if v79 == int32(2) {
						v82 = v66
					} else {
						v82 = int32(16)
					}
					v83 = F_palloc(m, v82)
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						v85 = int32(4)
						v87 = int32(1)
						v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						v91 = v89 & v87
						if v91 != 0 {
							v92 = v87
						} else {
							v92 = v85
						}
						v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v92))))
						if v94 == int32(2) {
							v97 = v85
						} else {
							v97 = int32(16)
						}
						if v97 != 0 {
							if v91 != 0 {
								v98 = v69
							} else {
								v98 = v67
							}
							base.MemoryCopy(m, v83, v98+int32(2), v97)
						} else {
						}
						v102 = int32(4)
						v104 = int32(1)
						v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
						v108 = v106 & v104
						if v108 != 0 {
							v109 = v104
						} else {
							v109 = v102
						}
						v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v109))))
						if v111 == int32(2) {
							v114 = v102
						} else {
							v114 = int32(16)
						}
						v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						if v115&int32(1) != 0 {
							v118 = v69
						} else {
							v118 = v67
						}
						v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
						if v108 != 0 {
							v120 = v41
						} else {
							v120 = v39
						}
						v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
						v123 = int32(0)
						for {
							v135 = v123 << (uint(int32(3)) % 32)
							v136 = v121 - v135
							if v136 <= int32(7) {
								v139 = v123 + v47
								v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
								v143 = int32(0)
								if v143 < v136 {
									v146 = v136
								} else {
									v146 = v143
								}
								v149 = v140 & (int32(255) << (uint(int32(8)-v146) % 32))
								*(*uint8)(unsafe.Add(mBase, uint32(v139))) = uint8(v149)
							} else {
							}
							v152 = v119 - v135
							if v152 <= int32(7) {
								v155 = v123 + v83
								v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
								v159 = int32(0)
								if v159 < v152 {
									v162 = v152
								} else {
									v162 = v159
								}
								v165 = v156 & (int32(255) << (uint(int32(8)-v162) % 32))
								*(*uint8)(unsafe.Add(mBase, uint32(v155))) = uint8(v165)
							} else {
							}
							v169 = v123 + int32(1)
							if v169 != v114 {
								v123 = v169
								continue
							} else {
								break
							}
							break
						}
						v173 = v114
						v182 = float64(0)
						for {
							v183 = int32(1)
							v184 = v173 - v183
							v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v184))))
							v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184+v47))))
							v194 = base.F64_mul(base.F64_add(v182, base.F64_sub(base.F64_convert_i32_u(v186), base.F64_convert_i32_u(v189))), float64(0.00390625))
							if v183 < v173 {
								v173 = v184
								v182 = v194
								continue
							} else {
								break
							}
							break
						}
						F_pfree(m, v47)
						mBase = m.M
						v198 = m.ExcPending
						if v198 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v83)
							mBase = m.M
							v200 = m.ExcPending
							if v200 != 0 {
								return int32(0)
							} else {
								v211 = v194
								v212 = F_Float8GetDatum(m, v211)
								mBase = m.M
								v213 = m.ExcPending
								if v213 != 0 {
									return int32(0)
								} else {
									return v212
								}
							}
						}
					}
				}
			} else {
				v211 = float64(1)
				v212 = F_Float8GetDatum(m, v211)
				mBase = m.M
				v213 = m.ExcPending
				if v213 != 0 {
					return int32(0)
				} else {
					return v212
				}
			}
		}
	}
}
func F_brin_minmax_multi_distance_macaddr8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+7)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)))
	v10 = float64(0.00390625)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+6)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+6)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+5)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+4)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+3)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+2)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+1)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v68 = F_Float8GetDatum(m, base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_sub(base.F64_convert_i32_u(v4), base.F64_convert_i32_u(v7)), v10), base.F64_sub(base.F64_convert_i32_u(v12), base.F64_convert_i32_u(v14))), v10), base.F64_sub(base.F64_convert_i32_u(v20), base.F64_convert_i32_u(v22))), v10), base.F64_sub(base.F64_convert_i32_u(v28), base.F64_convert_i32_u(v30))), v10), base.F64_sub(base.F64_convert_i32_u(v36), base.F64_convert_i32_u(v38))), v10), base.F64_sub(base.F64_convert_i32_u(v44), base.F64_convert_i32_u(v46))), v10), base.F64_sub(base.F64_convert_i32_u(v52), base.F64_convert_i32_u(v54))), v10), base.F64_sub(base.F64_convert_i32_u(v60), base.F64_convert_i32_u(v62))), v10))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		return int32(0)
	} else {
		return v68
	}
}
func F_brin_minmax_multi_serialize(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
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
	var v93 int32
	_ = v93
	var v95 float64
	_ = v95
	var v98 int32
	_ = v98
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v270 int32
	_ = v270
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v408 int32
	_ = v408
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v22+v23<<(uint(int32(1))%32) <= v21 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v425 = F_brin_range_serialize(m, l1)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L6
	} else {
		goto L60
	}
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v28 == v22 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v33 = F_minmax_multi_get_strategy_procinfo(m, l0, v30, v31, int32(1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	return
L7:
	;
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	v36 = F_minmax_multi_get_procinfo(m, l0, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_serialize[0]))
	v44 = F_AllocSetContextCreateInternal(m, v39, int32(_a_F_brin_minmax_multi_serialize_0), int32(0), int32(_a_F_brin_minmax_multi_serialize_1), int32(_a_F_brin_minmax_multi_serialize_2))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v46 = int32(_a_F_brin_minmax_multi_serialize_3)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_serialize[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_serialize[0])) = v44
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v53 = F_build_expanded_ranges(m, v33, v50, l1, v19+int32(12))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v56 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v60 = v56 - int32(1)
	v63 = F_palloc0(m, v60<<(uint(int32(4))%32))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	v124 = v55
	v126 = int32(0)
	goto L13
L13:
	;
	v137 = F_reduce_expanded_ranges(m, v53, v56, v126, v21, v33, v124)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L6
	} else {
		goto L23
	}
L14:
	;
	if int32(0) < v60 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v68 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	F_pg_qsort(m, v63, v60, int32(16), int32(20))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L6
	} else {
		goto L22
	}
L18:
	;
	v86 = v53 + v68*int32(12)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v89 = F_FunctionCall2Coll(m, v36, v55, v87, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L20
	}
L19:
	;
	goto L17
L20:
	;
	v93 = v63 + v68<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v68
	v95 = *(*float64)(unsafe.Add(mBase, uint32(v89)))
	*(*float64)(unsafe.Add(mBase, uint32(v93)+8)) = v95
	v98 = v68 + int32(1)
	if v98 != v60 {
		v68 = v98
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v124 = v120
	v126 = v63
	goto L13
L23:
	;
	v139 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v139
	if v137 <= v139 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v392
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_serialize[0])) = v47
	F_MemoryContextDelete(m, v44)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L6
	} else {
		goto L59
	}
L25:
	;
	v386 = int32(0)
	goto L27
L26:
	;
	v145 = l1 + int32(36)
	v147 = v137 - int32(1)
	if v147 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v386
	v392 = v386
	goto L24
L28:
	;
	v270 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v270
	if v147 == v270 {
		goto L45
	} else {
		goto L46
	}
L29:
	;
	v239 = v53 + v225*int32(12)
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+8)))
	if v240 != 0 {
		v254 = v221
		goto L28
	} else {
		goto L43
	}
L30:
	;
	v150 = int32(0)
	v221 = v150
	v225 = v150
	goto L29
L31:
	;
	goto L32
L32:
	;
	v156 = int32(0)
	v159 = v156
	v163 = v156
	v164 = v156
	goto L33
L33:
	;
	v177 = v53 + v163*int32(12)
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+8)))
	if v178 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v137&int32(1) == int32(0) {
		v254 = v212
		goto L28
	} else {
		goto L42
	}
L35:
	;
	v181 = int32(2)
	v183 = v145 + v159<<(uint(v181)%32)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v184
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v186
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v188 + int32(1)
	v194 = v159 + v181
	goto L37
L36:
	;
	v194 = v159
	goto L37
L37:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+20)))
	if v196 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v199 = int32(2)
	v201 = v145 + v194<<(uint(v199)%32)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v201)+4)) = v204
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v206 + int32(1)
	v212 = v194 + v199
	goto L40
L39:
	;
	v212 = v194
	goto L40
L40:
	;
	v214 = int32(2)
	v215 = v163 + v214
	v217 = v164 + v214
	if v217 != v137&int32(2147483646) {
		v159 = v212
		v163 = v215
		v164 = v217
		goto L33
	} else {
		goto L41
	}
L41:
	;
	goto L34
L42:
	;
	v221 = v212
	v225 = v215
	goto L29
L43:
	;
	v241 = int32(2)
	v243 = v145 + v221<<(uint(v241)%32)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = v244
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+4)) = v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v248 + int32(1)
	v254 = v221 + v241
	goto L28
L44:
	;
	v358 = v53 + v343*int32(12)
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+8)))
	if v359 != int32(1) {
		v392 = v344
		goto L24
	} else {
		goto L58
	}
L45:
	;
	v340 = v254
	v343 = int32(0)
	v344 = v270
	goto L44
L46:
	;
	goto L47
L47:
	;
	v280 = int32(0)
	v282 = v254
	v285 = v280
	v286 = v270
	v287 = v280
	goto L48
L48:
	;
	v300 = v53 + v285*int32(12)
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+8)))
	if v301 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v137&int32(1) == int32(0) {
		v392 = v332
		goto L24
	} else {
		goto L57
	}
L50:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	*(*int32)(unsafe.Add(mBase, uint32(v145+v282<<(uint(int32(2))%32)))) = v307
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v310 = int32(1)
	v311 = v309 + v310
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v311
	v315 = v282 + v310
	v316 = v311
	goto L52
L51:
	;
	v315 = v282
	v316 = v286
	goto L52
L52:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+20)))
	if v317 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v300)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v145+v315<<(uint(int32(2))%32)))) = v323
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v326 = int32(1)
	v327 = v325 + v326
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v327
	v331 = v315 + v326
	v332 = v327
	goto L55
L54:
	;
	v331 = v315
	v332 = v316
	goto L55
L55:
	;
	v333 = int32(2)
	v334 = v285 + v333
	v336 = v287 + v333
	if v336 != v137&int32(2147483646) {
		v282 = v331
		v285 = v334
		v286 = v332
		v287 = v336
		goto L48
	} else {
		goto L56
	}
L56:
	;
	goto L49
L57:
	;
	v340 = v331
	v343 = v334
	v344 = v332
	goto L44
L58:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	*(*int32)(unsafe.Add(mBase, uint32(v145+v340<<(uint(int32(2))%32)))) = v365
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v386 = v367 + int32(1)
	goto L27
L59:
	;
	goto L1
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v425
	m.G0 = v19 + int32(16)
	return
}
func F_brin_summarize_range(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v118 float64
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = int64(0)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_brin_summarize_range[0])))
	if v20 == int32(1) {
		v25 = *(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[1]))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+316))
		v28 = base.B2i32(v26 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_brin_summarize_range[0])) = uint8(v28)
		v30 = v28
	} else {
		v30 = v2
	}
	if v30 == int32(0) {
		if base.Ui64(int64(4294967296)) <= base.Ui64(v15) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v187 = m.ExcPending
			if v187 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v190 = m.ExcPending
				if v190 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v11))) = v15
					F_errmsg(m, int32(_a_F_brin_summarize_range_0), v11)
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1403), int32(_a_F_brin_summarize_range_2))
						mBase = m.M
						v199 = m.ExcPending
						if v199 != 0 {
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
			v36 = F_IndexGetRelation(m, v13, int32(1))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				if v36 != 0 {
					v41 = F_table_open(m, v36, int32(4))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v11+int32(76)))) = v48
						v51 = *(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3]))
						*(*int32)(unsafe.Add(mBase, uint32(v11+int32(72)))) = v51
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v41)+48))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+80))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
						*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v55 | int32(2)
						*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v54
						v63 = int32(_a_F_brin_summarize_range_3)
						v65 = *(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[4]))
						v67 = v65 + int32(1)
						*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[4])) = v67
						F_RestrictSearchPath(m)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v76 = v41
							v77 = v67
							v79 = F_index_open(m, v13, int32(4))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
								v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+119)))
								if v82 != int32(105) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v206 = m.ExcPending
										if v206 != 0 {
											return int32(0)
										} else {
											v207 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v207 + int32(4)
											F_errmsg(m, int32(_a_F_brin_summarize_range_4), v11+int32(48))
											mBase = m.M
											v215 = m.ExcPending
											if v215 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1448), int32(_a_F_brin_summarize_range_2))
												mBase = m.M
												v220 = m.ExcPending
												if v220 != 0 {
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
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+84))
									if v85 != int32(3580) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v203 = m.ExcPending
										if v203 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(151027844))
											mBase = m.M
											v206 = m.ExcPending
											if v206 != 0 {
												return int32(0)
											} else {
												v207 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v207 + int32(4)
												F_errmsg(m, int32(_a_F_brin_summarize_range_4), v11+int32(48))
												mBase = m.M
												v215 = m.ExcPending
												if v215 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1448), int32(_a_F_brin_summarize_range_2))
													mBase = m.M
													v220 = m.ExcPending
													if v220 != 0 {
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
										if v76 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v224 = m.ExcPending
											if v224 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(16908420))
												mBase = m.M
												v227 = m.ExcPending
												if v227 != 0 {
													return int32(0)
												} else {
													v228 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v228 + int32(4)
													F_errmsg(m, int32(_a_F_brin_summarize_range_5), v11+int32(16))
													mBase = m.M
													v236 = m.ExcPending
													if v236 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1464), int32(_a_F_brin_summarize_range_2))
														mBase = m.M
														v241 = m.ExcPending
														if v241 != 0 {
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
											v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
											v92 = F_object_ownercheck(m, int32(1259), v13, v91)
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return int32(0)
											} else {
												if v92 == int32(0) {
													v98 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
													F_aclcheck_error(m, int32(2), int32(20), v98+int32(4))
													mBase = m.M
													v102 = m.ExcPending
													if v102 != 0 {
														return int32(0)
													} else {
														v104 = F_IndexGetRelation(m, v13, int32(0))
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return int32(0)
														} else {
															if v104 != v36 {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v224 = m.ExcPending
																if v224 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(16908420))
																	mBase = m.M
																	v227 = m.ExcPending
																	if v227 != 0 {
																		return int32(0)
																	} else {
																		v228 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
																		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v228 + int32(4)
																		F_errmsg(m, int32(_a_F_brin_summarize_range_5), v11+int32(16))
																		mBase = m.M
																		v236 = m.ExcPending
																		if v236 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1464), int32(_a_F_brin_summarize_range_2))
																			mBase = m.M
																			v241 = m.ExcPending
																			if v241 != 0 {
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
																v107 = *(*int32)(unsafe.Add(mBase, uint32(v79)+192))
																v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+18)))
																if v108 == int32(1) {
																	F_brinsummarize(m, v79, v76, base.I32_wrap_i64(v15), int32(1), v11-int32(-64), int32(0))
																	mBase = m.M
																	v117 = m.ExcPending
																	if v117 != 0 {
																		return int32(0)
																	} else {
																		v118 = *(*float64)(unsafe.Add(mBase, uint32(v11)+64))
																		v144 = base.I32_trunc_sat_f64_s(v118)
																		F_AtEOXact_GUC(m, int32(0), v77)
																		mBase = m.M
																		v147 = m.ExcPending
																		if v147 != 0 {
																			return int32(0)
																		} else {
																			v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																			v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																			*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																			*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																			F_relation_close(m, v79, int32(4))
																			mBase = m.M
																			v156 = m.ExcPending
																			if v156 != 0 {
																				return int32(0)
																			} else {
																				F_relation_close(m, v76, int32(4))
																				mBase = m.M
																				v159 = m.ExcPending
																				if v159 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v11 + int32(80)
																					return v144
																				}
																			}
																		}
																	}
																} else {
																	v120 = int32(0)
																	v123 = F_errstart(m, int32(14), v120)
																	mBase = m.M
																	v124 = m.ExcPending
																	if v124 != 0 {
																		return int32(0)
																	} else {
																		if v123 == int32(0) {
																			v144 = v120
																			F_AtEOXact_GUC(m, int32(0), v77)
																			mBase = m.M
																			v147 = m.ExcPending
																			if v147 != 0 {
																				return int32(0)
																			} else {
																				v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																				v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																				*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																				*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																				F_relation_close(m, v79, int32(4))
																				mBase = m.M
																				v156 = m.ExcPending
																				if v156 != 0 {
																					return int32(0)
																				} else {
																					F_relation_close(m, v76, int32(4))
																					mBase = m.M
																					v159 = m.ExcPending
																					if v159 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v11 + int32(80)
																						return v144
																					}
																				}
																			}
																		} else {
																			F_errcode(m, int32(325))
																			mBase = m.M
																			v129 = m.ExcPending
																			if v129 != 0 {
																				return int32(0)
																			} else {
																				v130 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v130 + int32(4)
																				F_errmsg(m, int32(_a_F_brin_summarize_range_6), v11+int32(32))
																				mBase = m.M
																				v138 = m.ExcPending
																				if v138 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1473), int32(_a_F_brin_summarize_range_2))
																					mBase = m.M
																					v143 = m.ExcPending
																					if v143 != 0 {
																						return int32(0)
																					} else {
																						v144 = v120
																						F_AtEOXact_GUC(m, int32(0), v77)
																						mBase = m.M
																						v147 = m.ExcPending
																						if v147 != 0 {
																							return int32(0)
																						} else {
																							v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																							v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																							*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																							*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																							F_relation_close(m, v79, int32(4))
																							mBase = m.M
																							v156 = m.ExcPending
																							if v156 != 0 {
																								return int32(0)
																							} else {
																								F_relation_close(m, v76, int32(4))
																								mBase = m.M
																								v159 = m.ExcPending
																								if v159 != 0 {
																									return int32(0)
																								} else {
																									m.G0 = v11 + int32(80)
																									return v144
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
												} else {
													v104 = F_IndexGetRelation(m, v13, int32(0))
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return int32(0)
													} else {
														if v104 != v36 {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v224 = m.ExcPending
															if v224 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(16908420))
																mBase = m.M
																v227 = m.ExcPending
																if v227 != 0 {
																	return int32(0)
																} else {
																	v228 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v228 + int32(4)
																	F_errmsg(m, int32(_a_F_brin_summarize_range_5), v11+int32(16))
																	mBase = m.M
																	v236 = m.ExcPending
																	if v236 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1464), int32(_a_F_brin_summarize_range_2))
																		mBase = m.M
																		v241 = m.ExcPending
																		if v241 != 0 {
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
															v107 = *(*int32)(unsafe.Add(mBase, uint32(v79)+192))
															v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+18)))
															if v108 == int32(1) {
																F_brinsummarize(m, v79, v76, base.I32_wrap_i64(v15), int32(1), v11-int32(-64), int32(0))
																mBase = m.M
																v117 = m.ExcPending
																if v117 != 0 {
																	return int32(0)
																} else {
																	v118 = *(*float64)(unsafe.Add(mBase, uint32(v11)+64))
																	v144 = base.I32_trunc_sat_f64_s(v118)
																	F_AtEOXact_GUC(m, int32(0), v77)
																	mBase = m.M
																	v147 = m.ExcPending
																	if v147 != 0 {
																		return int32(0)
																	} else {
																		v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																		v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																		*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																		*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																		F_relation_close(m, v79, int32(4))
																		mBase = m.M
																		v156 = m.ExcPending
																		if v156 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v76, int32(4))
																			mBase = m.M
																			v159 = m.ExcPending
																			if v159 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v11 + int32(80)
																				return v144
																			}
																		}
																	}
																}
															} else {
																v120 = int32(0)
																v123 = F_errstart(m, int32(14), v120)
																mBase = m.M
																v124 = m.ExcPending
																if v124 != 0 {
																	return int32(0)
																} else {
																	if v123 == int32(0) {
																		v144 = v120
																		F_AtEOXact_GUC(m, int32(0), v77)
																		mBase = m.M
																		v147 = m.ExcPending
																		if v147 != 0 {
																			return int32(0)
																		} else {
																			v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																			v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																			*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																			*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																			F_relation_close(m, v79, int32(4))
																			mBase = m.M
																			v156 = m.ExcPending
																			if v156 != 0 {
																				return int32(0)
																			} else {
																				F_relation_close(m, v76, int32(4))
																				mBase = m.M
																				v159 = m.ExcPending
																				if v159 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v11 + int32(80)
																					return v144
																				}
																			}
																		}
																	} else {
																		F_errcode(m, int32(325))
																		mBase = m.M
																		v129 = m.ExcPending
																		if v129 != 0 {
																			return int32(0)
																		} else {
																			v130 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
																			*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v130 + int32(4)
																			F_errmsg(m, int32(_a_F_brin_summarize_range_6), v11+int32(32))
																			mBase = m.M
																			v138 = m.ExcPending
																			if v138 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1473), int32(_a_F_brin_summarize_range_2))
																				mBase = m.M
																				v143 = m.ExcPending
																				if v143 != 0 {
																					return int32(0)
																				} else {
																					v144 = v120
																					F_AtEOXact_GUC(m, int32(0), v77)
																					mBase = m.M
																					v147 = m.ExcPending
																					if v147 != 0 {
																						return int32(0)
																					} else {
																						v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																						v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																						*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																						*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																						F_relation_close(m, v79, int32(4))
																						mBase = m.M
																						v156 = m.ExcPending
																						if v156 != 0 {
																							return int32(0)
																						} else {
																							F_relation_close(m, v76, int32(4))
																							mBase = m.M
																							v159 = m.ExcPending
																							if v159 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v11 + int32(80)
																								return v144
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
											}
										}
									}
								}
							}
						}
					}
				} else {
					v71 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v71
					*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = int32(0)
					v76 = v2
					v77 = v71
					v79 = F_index_open(m, v13, int32(4))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
						v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+119)))
						if v82 != int32(105) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v203 = m.ExcPending
							if v203 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(151027844))
								mBase = m.M
								v206 = m.ExcPending
								if v206 != 0 {
									return int32(0)
								} else {
									v207 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v207 + int32(4)
									F_errmsg(m, int32(_a_F_brin_summarize_range_4), v11+int32(48))
									mBase = m.M
									v215 = m.ExcPending
									if v215 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1448), int32(_a_F_brin_summarize_range_2))
										mBase = m.M
										v220 = m.ExcPending
										if v220 != 0 {
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
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+84))
							if v85 != int32(3580) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v203 = m.ExcPending
								if v203 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(151027844))
									mBase = m.M
									v206 = m.ExcPending
									if v206 != 0 {
										return int32(0)
									} else {
										v207 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v207 + int32(4)
										F_errmsg(m, int32(_a_F_brin_summarize_range_4), v11+int32(48))
										mBase = m.M
										v215 = m.ExcPending
										if v215 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1448), int32(_a_F_brin_summarize_range_2))
											mBase = m.M
											v220 = m.ExcPending
											if v220 != 0 {
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
								if v76 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v224 = m.ExcPending
									if v224 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(16908420))
										mBase = m.M
										v227 = m.ExcPending
										if v227 != 0 {
											return int32(0)
										} else {
											v228 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v228 + int32(4)
											F_errmsg(m, int32(_a_F_brin_summarize_range_5), v11+int32(16))
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1464), int32(_a_F_brin_summarize_range_2))
												mBase = m.M
												v241 = m.ExcPending
												if v241 != 0 {
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
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
									v92 = F_object_ownercheck(m, int32(1259), v13, v91)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int32(0)
									} else {
										if v92 == int32(0) {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
											F_aclcheck_error(m, int32(2), int32(20), v98+int32(4))
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return int32(0)
											} else {
												v104 = F_IndexGetRelation(m, v13, int32(0))
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return int32(0)
												} else {
													if v104 != v36 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v224 = m.ExcPending
														if v224 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(16908420))
															mBase = m.M
															v227 = m.ExcPending
															if v227 != 0 {
																return int32(0)
															} else {
																v228 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
																*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v228 + int32(4)
																F_errmsg(m, int32(_a_F_brin_summarize_range_5), v11+int32(16))
																mBase = m.M
																v236 = m.ExcPending
																if v236 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1464), int32(_a_F_brin_summarize_range_2))
																	mBase = m.M
																	v241 = m.ExcPending
																	if v241 != 0 {
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
														v107 = *(*int32)(unsafe.Add(mBase, uint32(v79)+192))
														v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+18)))
														if v108 == int32(1) {
															F_brinsummarize(m, v79, v76, base.I32_wrap_i64(v15), int32(1), v11-int32(-64), int32(0))
															mBase = m.M
															v117 = m.ExcPending
															if v117 != 0 {
																return int32(0)
															} else {
																v118 = *(*float64)(unsafe.Add(mBase, uint32(v11)+64))
																v144 = base.I32_trunc_sat_f64_s(v118)
																F_AtEOXact_GUC(m, int32(0), v77)
																mBase = m.M
																v147 = m.ExcPending
																if v147 != 0 {
																	return int32(0)
																} else {
																	v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																	v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																	*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																	*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																	F_relation_close(m, v79, int32(4))
																	mBase = m.M
																	v156 = m.ExcPending
																	if v156 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v76, int32(4))
																		mBase = m.M
																		v159 = m.ExcPending
																		if v159 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v11 + int32(80)
																			return v144
																		}
																	}
																}
															}
														} else {
															v120 = int32(0)
															v123 = F_errstart(m, int32(14), v120)
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return int32(0)
															} else {
																if v123 == int32(0) {
																	v144 = v120
																	F_AtEOXact_GUC(m, int32(0), v77)
																	mBase = m.M
																	v147 = m.ExcPending
																	if v147 != 0 {
																		return int32(0)
																	} else {
																		v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																		v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																		*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																		*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																		F_relation_close(m, v79, int32(4))
																		mBase = m.M
																		v156 = m.ExcPending
																		if v156 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v76, int32(4))
																			mBase = m.M
																			v159 = m.ExcPending
																			if v159 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v11 + int32(80)
																				return v144
																			}
																		}
																	}
																} else {
																	F_errcode(m, int32(325))
																	mBase = m.M
																	v129 = m.ExcPending
																	if v129 != 0 {
																		return int32(0)
																	} else {
																		v130 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
																		*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v130 + int32(4)
																		F_errmsg(m, int32(_a_F_brin_summarize_range_6), v11+int32(32))
																		mBase = m.M
																		v138 = m.ExcPending
																		if v138 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1473), int32(_a_F_brin_summarize_range_2))
																			mBase = m.M
																			v143 = m.ExcPending
																			if v143 != 0 {
																				return int32(0)
																			} else {
																				v144 = v120
																				F_AtEOXact_GUC(m, int32(0), v77)
																				mBase = m.M
																				v147 = m.ExcPending
																				if v147 != 0 {
																					return int32(0)
																				} else {
																					v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																					v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																					*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																					*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																					F_relation_close(m, v79, int32(4))
																					mBase = m.M
																					v156 = m.ExcPending
																					if v156 != 0 {
																						return int32(0)
																					} else {
																						F_relation_close(m, v76, int32(4))
																						mBase = m.M
																						v159 = m.ExcPending
																						if v159 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v11 + int32(80)
																							return v144
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
										} else {
											v104 = F_IndexGetRelation(m, v13, int32(0))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return int32(0)
											} else {
												if v104 != v36 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v224 = m.ExcPending
													if v224 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(16908420))
														mBase = m.M
														v227 = m.ExcPending
														if v227 != 0 {
															return int32(0)
														} else {
															v228 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v228 + int32(4)
															F_errmsg(m, int32(_a_F_brin_summarize_range_5), v11+int32(16))
															mBase = m.M
															v236 = m.ExcPending
															if v236 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1464), int32(_a_F_brin_summarize_range_2))
																mBase = m.M
																v241 = m.ExcPending
																if v241 != 0 {
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
													v107 = *(*int32)(unsafe.Add(mBase, uint32(v79)+192))
													v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+18)))
													if v108 == int32(1) {
														F_brinsummarize(m, v79, v76, base.I32_wrap_i64(v15), int32(1), v11-int32(-64), int32(0))
														mBase = m.M
														v117 = m.ExcPending
														if v117 != 0 {
															return int32(0)
														} else {
															v118 = *(*float64)(unsafe.Add(mBase, uint32(v11)+64))
															v144 = base.I32_trunc_sat_f64_s(v118)
															F_AtEOXact_GUC(m, int32(0), v77)
															mBase = m.M
															v147 = m.ExcPending
															if v147 != 0 {
																return int32(0)
															} else {
																v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																F_relation_close(m, v79, int32(4))
																mBase = m.M
																v156 = m.ExcPending
																if v156 != 0 {
																	return int32(0)
																} else {
																	F_relation_close(m, v76, int32(4))
																	mBase = m.M
																	v159 = m.ExcPending
																	if v159 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v11 + int32(80)
																		return v144
																	}
																}
															}
														}
													} else {
														v120 = int32(0)
														v123 = F_errstart(m, int32(14), v120)
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return int32(0)
														} else {
															if v123 == int32(0) {
																v144 = v120
																F_AtEOXact_GUC(m, int32(0), v77)
																mBase = m.M
																v147 = m.ExcPending
																if v147 != 0 {
																	return int32(0)
																} else {
																	v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																	v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																	*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																	*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																	F_relation_close(m, v79, int32(4))
																	mBase = m.M
																	v156 = m.ExcPending
																	if v156 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v76, int32(4))
																		mBase = m.M
																		v159 = m.ExcPending
																		if v159 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v11 + int32(80)
																			return v144
																		}
																	}
																}
															} else {
																F_errcode(m, int32(325))
																mBase = m.M
																v129 = m.ExcPending
																if v129 != 0 {
																	return int32(0)
																} else {
																	v130 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v130 + int32(4)
																	F_errmsg(m, int32(_a_F_brin_summarize_range_6), v11+int32(32))
																	mBase = m.M
																	v138 = m.ExcPending
																	if v138 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1473), int32(_a_F_brin_summarize_range_2))
																		mBase = m.M
																		v143 = m.ExcPending
																		if v143 != 0 {
																			return int32(0)
																		} else {
																			v144 = v120
																			F_AtEOXact_GUC(m, int32(0), v77)
																			mBase = m.M
																			v147 = m.ExcPending
																			if v147 != 0 {
																				return int32(0)
																			} else {
																				v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																				v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																				*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																				*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																				F_relation_close(m, v79, int32(4))
																				mBase = m.M
																				v156 = m.ExcPending
																				if v156 != 0 {
																					return int32(0)
																				} else {
																					F_relation_close(m, v76, int32(4))
																					mBase = m.M
																					v159 = m.ExcPending
																					if v159 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v11 + int32(80)
																						return v144
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
		v167 = m.ExcPending
		if v167 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v170 = m.ExcPending
			if v170 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_brin_summarize_range_7), int32(0))
				mBase = m.M
				v174 = m.ExcPending
				if v174 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(_a_F_brin_summarize_range_8), int32(0))
					mBase = m.M
					v178 = m.ExcPending
					if v178 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1398), int32(_a_F_brin_summarize_range_2))
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
		}
	}
}
