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
												F_s_lock(m, l1+int32(44), int32(496462), int32(2847), int32(430509))
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
									F_s_lock(m, l1+int32(44), int32(496462), int32(2847), int32(430509))
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = base.I32_div_u_s(l1, v17)
	v20 = base.I32_div_u_s(v18, int32(1360))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = base.B2i32(base.Ui32(v20) < base.Ui32(v21))
	if v22 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v285
L2:
	;
	v25 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l3))) = uint16(v25)
	v285 = v25
	goto L1
L3:
	;
	goto L4
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v20) < base.Ui32(v21) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v32 = v20 + int32(1)
	goto L7
L6:
	;
	v32 = int32(-1)
	goto L7
L7:
	;
	v33 = v18 * v17
	v34 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)) = uint16(v34)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(-1)
	goto L8
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v51 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v56 == int32(0) {
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
	F_LockBuffer(m, v90, int32(1))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L13
	} else {
		goto L26
	}
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = F_ReadBuffer(m, v86, v32)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L13
	} else {
		goto L25
	}
L17:
	;
	if v56 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v77 == v32 {
		v90 = v78
		goto L15
	} else {
		goto L22
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62+(v56^int32(-1))<<(uint(int32(6))%32))+16))
	v77 = v68
	goto L18
L20:
	;
	goto L21
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70+v56<<(uint(int32(6))%32)+int32(-64))+16))
	v77 = v76
	goto L18
L22:
	;
	if v78 == int32(0) {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	F_ReleaseBuffer(m, v78)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	goto L16
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v87
	v90 = v87
	goto L15
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v95 < int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)))
	if v128 != 0 {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v115 = base.I32_div_u_s(v33, v114)
	v117 = base.I32_rem_u_s(v115, int32(1360))
	v120 = v113 + v117*int32(6)
	v122 = v120 + int32(24)
	if v122 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99+(v95^int32(-1))<<(uint(int32(2))%32))))
	v113 = v105
	goto L28
L30:
	;
	goto L31
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v113 = v107 + v95<<(uint(int32(13))%32) + int32(-8192)
	goto L28
L32:
	;
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)))
	if v123 != 0 {
		goto L27
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v124 = int32(0)
	F_LockBuffer(m, v95, v124)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L13
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	v285 = v124
	goto L1
L37:
	;
	F_LockBuffer(m, v203, int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L13
	} else {
		goto L81
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L13
	} else {
		goto L77
	}
L39:
	;
	v130 = v15 + int32(8)
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+2)))
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130))))
	v133 = int32(16)
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+2)))
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122))))
	if v131|v132<<(uint(v133)%32) == v136|v137<<(uint(v133)%32) {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	goto L41
L41:
	;
	v149 = v120 + int32(28)
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149))))
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)) = uint16(v150)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v152
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+2)))
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122))))
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149))))
	*(*uint16)(unsafe.Add(mBase, uint32(l3))) = uint16(v156)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_LockBuffer(m, v158, int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L13
	} else {
		goto L49
	}
L42:
	;
	if v147 != 0 {
		goto L38
	} else {
		goto L48
	}
L43:
	;
	goto L42
L44:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+4)))
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)))
	if v143 == v144 {
		v147 = int32(1)
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v147 = int32(0)
	goto L43
L47:
	;
	goto L46
L48:
	;
	goto L41
L49:
	;
	v164 = v154 | v155<<(uint(int32(16))%32)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v165 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	F_LockBuffer(m, v198, int32(1))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L13
	} else {
		goto L61
	}
L51:
	;
	v195 = F_ReadBuffer(m, v28, v164)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L13
	} else {
		goto L60
	}
L52:
	;
	if v165 < int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v164 == v186 {
		v198 = v187
		goto L50
	} else {
		goto L57
	}
L54:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v171+(v165^int32(-1))<<(uint(int32(6))%32))+16))
	v186 = v177
	goto L53
L55:
	;
	goto L56
L56:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v179+v165<<(uint(int32(6))%32)+int32(-64))+16))
	v186 = v185
	goto L53
L57:
	;
	if v187 == int32(0) {
		goto L51
	} else {
		goto L58
	}
L58:
	;
	F_ReleaseBuffer(m, v187)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	goto L51
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v195
	v198 = v195
	goto L50
L61:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v203 < int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221)+16)))
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222+v221)+6)))
	if v224 != int32(61587) {
		goto L37
	} else {
		goto L66
	}
L63:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v207+(v203^int32(-1))<<(uint(int32(2))%32))))
	v221 = v213
	goto L62
L64:
	;
	goto L65
L65:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v221 = v215 + v203<<(uint(int32(13))%32) + int32(-8192)
	goto L62
L66:
	;
	v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v228) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v236 = int32(base.Ui32(v228+int32(262120)) >> (uint(int32(2)) % 32))
	goto L69
L68:
	;
	v236 = int32(0)
	goto L69
L69:
	;
	if base.Ui32(v236&int32(65535)) < base.Ui32(v227) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v240 = int32(0)
	F_LockBuffer(m, v203, v240)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L13
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v227<<(uint(int32(2))%32)+v221)+20))
	if v247&int32(98304) == int32(0) {
		goto L37
	} else {
		goto L74
	}
L73:
	;
	v285 = v240
	goto L1
L74:
	;
	v254 = v221 + v247&int32(32767)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	if v255 != v33 {
		goto L37
	} else {
		goto L75
	}
L75:
	;
	if l4 == int32(0) {
		v285 = v254
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(base.Ui32(v247) >> (uint(int32(17)) % 32))
	v285 = v254
	goto L1
L77:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L13
	} else {
		goto L78
	}
L78:
	;
	F_errmsg_internal(m, int32(238464), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L13
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(495901), int32(259), int32(316968))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L13
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
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
		v12 = *(*int32)(unsafe.Add(mBase, _consts[9]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v12+(l0^int32(-1))<<(uint(int32(2))%32))))
		v26 = v18
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[10]))
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
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_initStringInfo(m, v6+int32(16))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			F_appendStringInfoChar(m, v6+int32(16), int32(123))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+6)))
				v23 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v22
				F_appendStringInfo(m, v6+int32(16), int32(57479), v6)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_appendStringInfoChar(m, v6+int32(16), int32(125))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
						m.G0 = v6 + int32(32)
						return v36
					}
				}
			}
		}
	}
}
func F_brin_bloom_summary_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(17894)
			F_errmsg(m, int32(192450), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(497076), int32(826), int32(36155))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
		F_appendStringInfo(m, l0, int32(51941), v8)
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
		F_appendStringInfo(m, l0, int32(47122), v8+int32(16))
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
		F_appendStringInfo(m, l0, int32(47065), v8+int32(32))
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
		F_appendStringInfo(m, l0, int32(47150), v8+int32(48))
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
		F_appendStringInfo(m, l0, int32(47543), v8-int32(-64))
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
		F_appendStringInfo(m, l0, int32(41335), v8+int32(80))
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v316 int32
	_ = v316
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v471 int32
	_ = v471
	var v509 int32
	_ = v509
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v663 int32
	_ = v663
	var v701 int32
	_ = v701
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v738 int32
	_ = v738
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	v5 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(16)
	m.G0 = v29
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+14)) = uint16(v5)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v36 = F_palloc(m, v33<<(uint(int32(2))%32))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v41 = F_palloc0(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v43 = int32(8)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v48 = base.I32_div_s(v44+int32(7), v43)
	v49 = F_palloc(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v54 = F_palloc(m, v51<<(uint(int32(2))%32))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v57 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v347 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L7:
	;
	v332 = v43
	v334 = v5
	v336 = v5
	goto L6
L8:
	;
	goto L9
L9:
	;
	v61 = l0 + int32(20)
	v68 = v5
	v73 = v5
	v78 = v5
	v79 = v5
	goto L10
L10:
	;
	v92 = l2 + int32(24) + v78*int32(20)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+3)))
	if v93 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v304 = int32(1)
	if v282&v304 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L12:
	;
	v300 = v78 + int32(1)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	if v300 < v302 {
		v68 = v277
		v73 = v282
		v78 = v300
		v79 = v288
		goto L10
	} else {
		goto L55
	}
L13:
	;
	v96 = int32(0)
	v99 = v61 + v78<<(uint(int32(2))%32)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100))))
	if v101 == v96 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+2)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
	if v143 != 0 {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v277 = v68
	v282 = int32(1)
	v288 = v79
	goto L12
L17:
	;
	goto L18
L18:
	;
	v109 = v68
	v110 = v96
	goto L19
L19:
	;
	v131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v109+v41))) = uint8(v131)
	v136 = v109 + v131
	v138 = v110 + v131
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139))))
	if base.Ui32(v138) < base.Ui32(v140) {
		v109 = v136
		v110 = v138
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v277 = v136
	v282 = v131
	v288 = v79
	goto L12
L21:
	;
	goto L20
L22:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	m.T0[v143].(func(*base.Module, int32, int32, int32))(m, l0, v144, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v148 = v142 | v73
	v151 = v61 + v78<<(uint(int32(2))%32)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	if v153 == int32(0) {
		v277 = v68
		v282 = v148
		v288 = v79
		goto L12
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	v163 = v68
	v164 = int32(0)
	v167 = v152
	v174 = v79
	goto L27
L27:
	;
	v186 = v164 << (uint(int32(2)) % 32)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v186+v187)))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v186+v167)+8))
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+8)))
	if v192 != int32(65535) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v277 = v267
	v282 = v148
	v288 = v260
	goto L12
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36+v163<<(uint(int32(2))%32)))) = v256
	v266 = int32(1)
	v267 = v163 + v266
	v269 = v164 + v266
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v270))))
	if base.Ui32(v269) < base.Ui32(v271) {
		v163 = v267
		v164 = v269
		v167 = v270
		v174 = v260
		goto L27
	} else {
		goto L54
	}
L30:
	;
	v256 = v189
	v260 = v174
	goto L29
L31:
	;
	goto L32
L32:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v197 = base.B2i32(v195 != int32(1))
	if v195 != int32(1) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v204 = base.B2i32(v195 == int32(1))
	if v202&int32(3) != 0 {
		v239 = v201
		v240 = v204
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v201 = v189
	v202 = v195
	goto L33
L35:
	;
	goto L36
L36:
	;
	v198 = F_detoast_external_attr(m, v189)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	v201 = v198
	v202 = v200
	goto L33
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54+v174<<(uint(int32(2))%32)))) = v245
	v256 = v245
	v260 = v174 + int32(1)
	goto L29
L39:
	;
	if v240 == int32(0) {
		v256 = v239
		v260 = v174
		goto L29
	} else {
		goto L53
	}
L40:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if base.Ui32(v207) < base.Ui32(int32(2044)) {
		v239 = v201
		v240 = v204
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+12)))
	switch v210 - int32(109) {
	case 0, 11:
		goto L42
	default:
		v239 = v201
		v240 = v204
		goto L39
	}
L42:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v221 = v216 + v217<<(uint(int32(4))%32) + v78*int32(100)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+88))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	if v222 == v223 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+105)))
	v226 = v225
	goto L45
L44:
	;
	v226 = int32(0)
	goto L45
L45:
	;
	v228 = F_toast_compress_datum(m, v201, base.I32_extend8_s(v226))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v232 = base.B2i32(v195 == int32(1)) | base.B2i32(v228 != int32(0))
	if v228 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v233 = v228
	goto L49
L48:
	;
	v233 = v201
	goto L49
L49:
	;
	if v195 != int32(1) {
		v239 = v233
		v240 = v232
		goto L39
	} else {
		goto L50
	}
L50:
	;
	if v228 == int32(0) {
		v239 = v233
		v240 = v232
		goto L39
	} else {
		goto L51
	}
L51:
	;
	F_pfree(m, v201)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v245 = v228
	goto L38
L53:
	;
	v245 = v239
	goto L38
L54:
	;
	goto L28
L55:
	;
	goto L11
L56:
	;
	v332 = int32(8)
	v334 = int32(0)
	v336 = v288
	goto L6
L57:
	;
	goto L58
L58:
	;
	v316 = base.I32_div_s(v302<<(uint(int32(1))%32)+int32(7), int32(8))
	v332 = (v316 + int32(12)) & int32(-8)
	v334 = v304
	v336 = v288
	goto L6
L59:
	;
	v350 = int32(4515392)
	v351 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v353
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v356 = F_CreateTemplateTupleDesc(m, v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	v509 = v347
	goto L61
L61:
	;
	v528 = F_heap_compute_data_size(m, v509, v36, v41)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L76
	}
L62:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	if int32(0) < v359 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v371 = int32(1)
	v375 = v359
	v376 = int32(0)
	goto L66
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v351
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v356
	v509 = v356
	goto L61
L66:
	;
	v395 = l0 + int32(20) + v376<<(uint(int32(2))%32)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	v397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v396))))
	if v397 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L65
L68:
	;
	v402 = int32(0)
	v403 = v371
	v404 = v396
	goto L71
L69:
	;
	v449 = v371
	v453 = v375
	goto L70
L70:
	;
	v471 = v376 + int32(1)
	if v471 < v453 {
		v371 = v449
		v375 = v453
		v376 = v471
		goto L66
	} else {
		goto L75
	}
L71:
	;
	v425 = int32(0)
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v404+v402<<(uint(int32(2))%32))+8))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	F_TupleDescInitEntry(m, v356, base.I32_extend16_s(v403), v425, v430, int32(-1), v425)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	v449 = v436
	v453 = v443
	goto L70
L73:
	;
	v435 = int32(1)
	v436 = v403 + v435
	v438 = v402 + v435
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	v440 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v439))))
	if base.Ui32(v438) < base.Ui32(v440) {
		v402 = v438
		v403 = v436
		v404 = v439
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	goto L67
L76:
	;
	v534 = (v528 + v332 + int32(7)) & int32(-8)
	v535 = F_palloc0(m, v534)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v535)+4)) = uint8(v332)
	*(*int32)(unsafe.Add(mBase, uint32(v535))) = l1
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v539 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v542 = int32(4515392)
	v543 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v545
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v548 = F_CreateTemplateTupleDesc(m, v547)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	v701 = v539
	goto L80
L80:
	;
	F_heap_fill_tuple(m, v701, v36, v41, v332+v535, v29+int32(14), v49)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L95
	}
L81:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	if int32(0) < v551 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v563 = int32(1)
	v567 = v551
	v568 = int32(0)
	goto L85
L83:
	;
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v543
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v548
	v701 = v548
	goto L80
L85:
	;
	v587 = l0 + int32(20) + v568<<(uint(int32(2))%32)
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	v589 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v588))))
	if v589 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L84
L87:
	;
	v594 = int32(0)
	v595 = v563
	v596 = v588
	goto L90
L88:
	;
	v641 = v563
	v645 = v567
	goto L89
L89:
	;
	v663 = v568 + int32(1)
	if v663 < v645 {
		v563 = v641
		v567 = v645
		v568 = v663
		goto L85
	} else {
		goto L94
	}
L90:
	;
	v617 = int32(0)
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v596+v594<<(uint(int32(2))%32))+8))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v621)))
	F_TupleDescInitEntry(m, v548, base.I32_extend16_s(v595), v617, v622, int32(-1), v617)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L92
	}
L91:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v634)))
	v641 = v628
	v645 = v635
	goto L89
L92:
	;
	v627 = int32(1)
	v628 = v595 + v627
	v630 = v594 + v627
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	v632 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v631))))
	if base.Ui32(v630) < base.Ui32(v632) {
		v594 = v630
		v595 = v628
		v596 = v631
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	goto L86
L95:
	;
	F_pfree(m, v36)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_pfree(m, v41)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_pfree(m, v49)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	if int32(0) < v336 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v738 = int32(0)
	goto L102
L100:
	;
	goto L101
L101:
	;
	v796 = v535 + int32(4)
	if v334 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v54+v738<<(uint(int32(2))%32))))
	F_pfree(m, v763)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L104
	}
L103:
	;
	goto L101
L104:
	;
	v767 = v738 + int32(1)
	if v767 != v336 {
		v738 = v767
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v948 == int32(1) {
		goto L128
	} else {
		goto L129
	}
L107:
	;
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v796))))
	v801 = v799 | int32(-128)
	*(*uint8)(unsafe.Add(mBase, uint32(v796))) = uint8(v801)
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v803)))
	if v804 <= int32(0) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v815 = int32(128)
	v816 = int32(0)
	v817 = v801
	v819 = v796
	goto L109
L109:
	;
	if v815 != int32(128) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if v862 <= int32(0) {
		goto L106
	} else {
		goto L118
	}
L111:
	;
	v847 = v817
	v848 = v819
	v849 = v815 << (uint(int32(1)) % 32)
	goto L113
L112:
	;
	v841 = int32(0)
	v842 = int32(1)
	v843 = v819 + v842
	*(*uint8)(unsafe.Add(mBase, uint32(v843))) = uint8(v841)
	v847 = v841
	v848 = v843
	v849 = v842
	goto L113
L113:
	;
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(27)+v816*int32(20)))))
	if v853 == int32(1) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v856 = v849 | v847
	*(*uint8)(unsafe.Add(mBase, uint32(v848))) = uint8(v856)
	v858 = v856
	goto L116
L115:
	;
	v858 = v847
	goto L116
L116:
	;
	v860 = v816 + int32(1)
	v861 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v861)))
	if v860 < v862 {
		v815 = v849
		v816 = v860
		v817 = v858
		v819 = v848
		goto L109
	} else {
		goto L117
	}
L117:
	;
	goto L110
L118:
	;
	v873 = v849
	v874 = int32(0)
	v875 = v858
	v877 = v848
	goto L119
L119:
	;
	if v873 != int32(128) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	goto L106
L121:
	;
	v905 = v875
	v906 = v877
	v907 = v873 << (uint(int32(1)) % 32)
	goto L123
L122:
	;
	v899 = int32(0)
	v900 = int32(1)
	v901 = v877 + v900
	*(*uint8)(unsafe.Add(mBase, uint32(v901))) = uint8(v899)
	v905 = v899
	v906 = v901
	v907 = v900
	goto L123
L123:
	;
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(26)+v874*int32(20)))))
	if v911 == int32(1) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v914 = v907 | v905
	*(*uint8)(unsafe.Add(mBase, uint32(v906))) = uint8(v914)
	v916 = v914
	goto L126
L125:
	;
	v916 = v905
	goto L126
L126:
	;
	v918 = v874 + int32(1)
	v919 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v919)))
	if v918 < v920 {
		v873 = v907
		v874 = v918
		v875 = v916
		v877 = v906
		goto L119
	} else {
		goto L127
	}
L127:
	;
	goto L120
L128:
	;
	v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v796))))
	v953 = v951 | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v796))) = uint8(v953)
	goto L130
L129:
	;
	goto L130
L130:
	;
	v955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v955 == int32(1) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v796))))
	v960 = v958 | int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v796))) = uint8(v960)
	goto L133
L132:
	;
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v534
	m.G0 = v29 + int32(16)
	return v535
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
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int64
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
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
	v21 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v36 = v27
	goto L4
L6:
	;
	goto L7
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[7]))
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
	return v283
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v69 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_brin_initialize_empty_new_buffer(m, l0, v113)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L13
	} else {
		goto L104
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
	if l1 == int32(0) {
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
	v90 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90+(v85^int32(-1))<<(uint(int32(6))%32))+16))
	v105 = v96
	goto L31
L33:
	;
	goto L34
L34:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[7]))
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
	v164 = m.ExcPending
	if v164 != 0 {
		goto L13
	} else {
		goto L60
	}
L40:
	;
	if base.Ui32(v112) <= base.Ui32(v38) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	F_LockBuffer(m, l1, int32(2))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L13
	} else {
		goto L42
	}
L42:
	;
	if l1 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+16)))
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133+v132)+6)))
	if v135 == int32(61587) {
		goto L39
	} else {
		goto L47
	}
L44:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v124+(l1^int32(-1))<<(uint(int32(2))%32))))
	v132 = v126
	goto L43
L45:
	;
	goto L46
L46:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v132 = v128 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L43
L47:
	;
	F_LockBuffer(m, l1, int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v141 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	F_brin_initialize_empty_new_buffer(m, l0, v113)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v114 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	F_UnlockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L13
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_ReleaseBuffer(m, v113)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L13
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v151 = int32(0)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v152 != int32(1) {
		v283 = v151
		goto L15
	} else {
		goto L58
	}
L58:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v112, v112+int32(1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	v159 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v159)
	v283 = v151
	goto L15
L60:
	;
	if v114 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	F_UnlockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L13
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v113 < int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L63
L65:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v186 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v171+(v113^int32(-1))<<(uint(int32(2))%32))))
	v185 = v177
	goto L65
L67:
	;
	goto L68
L68:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v185 = v179 + v113<<(uint(int32(13))%32) + int32(-8192)
	goto L65
L69:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v243 != int32(1) {
		goto L92
	} else {
		goto L93
	}
L70:
	;
	v189 = int32(0)
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185)+16)))
	v191 = v185 + v190
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+6)))
	if v192 != int32(61587) {
		v207 = v189
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v211 != 0 {
		goto L81
	} else {
		goto L82
	}
L73:
	;
	if base.Ui32(v207) < base.Ui32(l2) {
		goto L69
	} else {
		goto L80
	}
L74:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+4)))
	if v195&int32(1) != 0 {
		v207 = v189
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v198 = int32(4)
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185)+14)))
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185)+12)))
	v201 = v199 - v200
	if v201 <= v198 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v207 = v204 - int32(4)
	goto L73
L77:
	;
	v204 = v198
	goto L79
L78:
	;
	v204 = v201
	goto L79
L79:
	;
	goto L76
L80:
	;
	goto L72
L81:
	;
	v235 = v211
	goto L83
L82:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v213
	v215 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v215
	v217 = F_smgropen(m, v16, v212)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L13
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235)+16)) = v112
	if l1 == int32(0) {
		v283 = v113
		goto L15
	} else {
		goto L89
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v217
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217)+72))
	if v221 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v235 = v233
	goto L83
L86:
	;
	v229 = v221
	goto L88
L87:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v217)+76))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v217)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v222)+4)) = v223
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v217)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v217)+72))
	v229 = v227
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+72)) = v229 + int32(1)
	goto L85
L89:
	;
	if base.Ui32(v38) <= base.Ui32(v112) {
		v283 = v113
		goto L15
	} else {
		goto L90
	}
L90:
	;
	F_LockBuffer(m, l1, int32(2))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L13
	} else {
		goto L91
	}
L91:
	;
	v283 = v113
	goto L15
L92:
	;
	if v112 != v38 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	goto L17
L95:
	;
	F_UnlockReleaseBuffer(m, v113)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L13
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	if l1 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L97
L99:
	;
	v255 = F_RecordAndGetPageWithFreeSpace(m, l0, v112, v207, l2)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L13
	} else {
		goto L103
	}
L100:
	;
	if base.Ui32(v112) < base.Ui32(v38) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	F_LockBuffer(m, l1, int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L13
	} else {
		goto L102
	}
L102:
	;
	goto L99
L103:
	;
	v59 = v255
	goto L16
L104:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L13
	} else {
		goto L105
	}
L105:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L13
	} else {
		goto L106
	}
L106:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v266 + int32(4)
	F_errmsg(m, int32(692891), v16+int32(16))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L13
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(493729), int32(852), int32(225252))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L13
	} else {
		goto L108
	}
L108:
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
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
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
	v269 = m.ExcPending
	if v269 != 0 {
		goto L3
	} else {
		goto L59
	}
L2:
	;
	m.G0 = v19 + int32(32)
	return v246
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
	if v21 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	if v157 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L9:
	;
	v246 = int32(1)
	goto L2
L10:
	;
	goto L11
L11:
	;
	v55 = int32(0)
	goto L12
L12:
	;
	v61 = v33 + int32(36) + v55<<(uint(int32(3))%32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v66 = int32(0)
	goto L14
L13:
	;
	goto L8
L14:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v22+v66<<(uint(int32(2))%32))))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+44))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+4)))
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+6)))
	switch v88 - int32(1) {
	case 0, 1:
		goto L21
	case 2:
		goto L18
	case 3, 4:
		goto L20
	default:
		goto L19
	}
L15:
	;
	v138 = v55 + int32(1)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v138 < v139 {
		v55 = v138
		goto L12
	} else {
		goto L38
	}
L16:
	;
	goto L15
L17:
	;
	v132 = int32(1)
	v134 = v66 + v132
	if v21 != v134 {
		v66 = v134
		goto L14
	} else {
		goto L37
	}
L18:
	;
	v120 = v87 & int32(65535)
	v122 = F_minmax_multi_get_strategy_procinfo(m, v24, v120, v86, int32(5))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L3
	} else {
		goto L31
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L28
	}
L20:
	;
	v101 = F_minmax_multi_get_strategy_procinfo(m, v24, v87&int32(65535), v86, v88)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L3
	} else {
		goto L25
	}
L21:
	;
	v93 = F_minmax_multi_get_strategy_procinfo(m, v24, v87&int32(65535), v86, v88)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v95 = F_FunctionCall2Coll(m, v93, v23, v62, v85)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	if v95 == int32(0) {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	goto L17
L25:
	;
	v103 = F_FunctionCall2Coll(m, v101, v23, v63, v85)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	if v103 != 0 {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	goto L16
L28:
	;
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v109
	F_errmsg_internal(m, int32(471321), v19)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(497644), int32(2650), int32(92423))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v124 = F_FunctionCall2Coll(m, v122, v23, v62, v85)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	if v124 != 0 {
		goto L16
	} else {
		goto L33
	}
L33:
	;
	v127 = F_minmax_multi_get_strategy_procinfo(m, v24, v120, v86, int32(1))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v129 = F_FunctionCall2Coll(m, v127, v23, v63, v85)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	if v129 != 0 {
		goto L16
	} else {
		goto L36
	}
L36:
	;
	goto L17
L37:
	;
	v246 = v132
	goto L2
L38:
	;
	goto L13
L39:
	;
	v246 = int32(0)
	goto L2
L40:
	;
	goto L41
L41:
	;
	v161 = int32(0)
	if v21 <= v161 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v246 = int32(1)
	goto L2
L43:
	;
	goto L44
L44:
	;
	v172 = v161
	goto L45
L45:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(36)+(v183<<(uint(int32(1))%32)+v172)<<(uint(int32(2))%32))))
	v193 = int32(0)
	goto L47
L46:
	;
	v246 = int32(0)
	goto L2
L47:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v22+v193<<(uint(int32(2))%32))))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	if v212&int32(1) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v243 = v172 + int32(1)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	if v243 < v244 {
		v172 = v243
		goto L45
	} else {
		goto L58
	}
L49:
	;
	goto L48
L50:
	;
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v211)+6)))
	if base.Ui32(int32(5)) <= base.Ui32((v217-int32(1))&int32(65535)) {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v237 = int32(1)
	v239 = v193 + v237
	if v239 != v21 {
		v193 = v239
		goto L47
	} else {
		goto L57
	}
L53:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v211)+44))
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v211)+4)))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	v229 = F_minmax_multi_get_strategy_procinfo(m, v24, v225, v226, v217&int32(65535))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v231 = F_FunctionCall2Coll(m, v229, v23, v190, v224)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	if v231 == int32(0) {
		goto L49
	} else {
		goto L56
	}
L56:
	;
	goto L52
L57:
	;
	v246 = v237
	goto L2
L58:
	;
	goto L46
L59:
	;
	v270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v211)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v270
	F_errmsg_internal(m, int32(471321), v19+int32(16))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(497644), int32(2709), int32(92423))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
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
	var v181 float64
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v193 float64
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v209 float64
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v20 = int32(1)
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v22&v20 != 0 {
				v25 = v20
			} else {
				v25 = int32(4)
			}
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v25))))
			v28 = int32(1)
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			if v30&v28 != 0 {
				v33 = v28
			} else {
				v33 = int32(4)
			}
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v33))))
			if v27 == v35 {
				if v27 == int32(2) {
					v41 = int32(4)
				} else {
					v41 = int32(16)
				}
				v42 = F_palloc(m, v41)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					v44 = int32(1)
					v45 = v12 + v44
					v47 = v12 + int32(4)
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
					v50 = v48 & v44
					if v50 != 0 {
						v51 = v45
					} else {
						v51 = v47
					}
					v54 = int32(4)
					if v50 != 0 {
						v58 = int32(1)
					} else {
						v58 = v54
					}
					v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v58))))
					if v60 == int32(2) {
						v63 = v54
					} else {
						v63 = int32(16)
					}
					if v63 != 0 {
						v64 = F__emscripten_memcpy_bulkmem(m, v42, v51+int32(2), v63)
						mBase = m.M
						v65 = v64
					} else {
						v65 = v42
					}
					v66 = int32(4)
					v68 = int32(1)
					v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
					if v70&v68 != 0 {
						v73 = v68
					} else {
						v73 = v66
					}
					v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v73))))
					if v75 == int32(2) {
						v78 = v66
					} else {
						v78 = int32(16)
					}
					v79 = F_palloc(m, v78)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						v81 = int32(1)
						v82 = v17 + v81
						v84 = v17 + int32(4)
						v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
						v87 = v85 & v81
						if v87 != 0 {
							v88 = v82
						} else {
							v88 = v84
						}
						v91 = int32(4)
						if v87 != 0 {
							v95 = int32(1)
						} else {
							v95 = v91
						}
						v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v95))))
						if v97 == int32(2) {
							v100 = v91
						} else {
							v100 = int32(16)
						}
						if v100 != 0 {
							v101 = F__emscripten_memcpy_bulkmem(m, v79, v88+int32(2), v100)
							mBase = m.M
							v102 = v101
						} else {
							v102 = v79
						}
						v103 = int32(4)
						v105 = int32(1)
						v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
						v109 = v107 & v105
						if v109 != 0 {
							v110 = v105
						} else {
							v110 = v103
						}
						v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v110))))
						if v112 == int32(2) {
							v115 = v103
						} else {
							v115 = int32(16)
						}
						v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
						if v116&int32(1) != 0 {
							v119 = v82
						} else {
							v119 = v84
						}
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
						if v109 != 0 {
							v121 = v45
						} else {
							v121 = v47
						}
						v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+1)))
						v124 = int32(0)
						for {
							v135 = v124 << (uint(int32(3)) % 32)
							v136 = v122 - v135
							if v136 <= int32(7) {
								v139 = v124 + v65
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
							v152 = v120 - v135
							if v152 <= int32(7) {
								v155 = v124 + v102
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
							v169 = v124 + int32(1)
							if v169 != v115 {
								v124 = v169
								continue
							} else {
								break
							}
							break
						}
						v173 = v115
						v181 = float64(0)
						for {
							v182 = int32(1)
							v183 = v173 - v182
							v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v183))))
							v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+v65))))
							v193 = base.F64_mul(base.F64_add(v181, base.F64_sub(base.F64_convert_i32_u(v185), base.F64_convert_i32_u(v188))), float64(0.00390625))
							if v182 < v173 {
								v173 = v183
								v181 = v193
								continue
							} else {
								break
							}
							break
						}
						F_pfree(m, v65)
						mBase = m.M
						v197 = m.ExcPending
						if v197 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v102)
							mBase = m.M
							v199 = m.ExcPending
							if v199 != 0 {
								return int32(0)
							} else {
								v209 = v193
								v210 = F_Float8GetDatum(m, v209)
								mBase = m.M
								v211 = m.ExcPending
								if v211 != 0 {
									return int32(0)
								} else {
									return v210
								}
							}
						}
					}
				}
			} else {
				v209 = float64(1)
				v210 = F_Float8GetDatum(m, v209)
				mBase = m.M
				v211 = m.ExcPending
				if v211 != 0 {
					return int32(0)
				} else {
					return v210
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
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v331 int32
	_ = v331
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
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
	v348 = F_brin_range_serialize(m, l1)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L6
	} else {
		goto L51
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
	v39 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v44 = F_AllocSetContextCreateInternal(m, v39, int32(60951), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v46 = int32(4515392)
	v47 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v44
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
	if v56 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v137 = F_reduce_expanded_ranges(m, v53, v56, v127, v21, v33, v125)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L6
	} else {
		goto L24
	}
L12:
	;
	v125 = v55
	v127 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v60 = v56 - int32(1)
	v63 = F_palloc0(m, v60<<(uint(int32(4))%32))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	if int32(0) < v60 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v68 = int32(0)
	goto L19
L17:
	;
	goto L18
L18:
	;
	F_pg_qsort(m, v63, v60, int32(16), int32(20))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L6
	} else {
		goto L23
	}
L19:
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
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	v93 = v63 + v68<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v68
	v95 = *(*float64)(unsafe.Add(mBase, uint32(v89)))
	*(*float64)(unsafe.Add(mBase, uint32(v93)+8)) = v95
	v98 = v68 + int32(1)
	if v98 != v60 {
		v68 = v98
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v125 = v120
	v127 = v63
	goto L11
L24:
	;
	v139 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v139
	if v137 <= v139 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v316
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v47
	F_MemoryContextDelete(m, v44)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L6
	} else {
		goto L50
	}
L26:
	;
	v309 = int32(0)
	goto L28
L27:
	;
	v145 = l1 + int32(36)
	v146 = int32(0)
	v148 = v146
	v151 = v146
	goto L29
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v309
	v316 = v309
	goto L25
L29:
	;
	v166 = v53 + v151*int32(12)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+8)))
	if v167 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v188 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v188
	v191 = int32(1)
	if v137 == v191 {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	v170 = int32(2)
	v172 = v145 + v148<<(uint(v170)%32)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = v173
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v172)+4)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v177 + int32(1)
	v183 = v148 + v170
	goto L33
L32:
	;
	v183 = v148
	goto L33
L33:
	;
	v186 = v151 + int32(1)
	if v186 != v137 {
		v148 = v183
		v151 = v186
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	if v137&v191 == int32(0) {
		v316 = v266
		goto L25
	} else {
		goto L48
	}
L36:
	;
	v261 = v183
	v264 = int32(0)
	v266 = v188
	goto L35
L37:
	;
	goto L38
L38:
	;
	v198 = int32(0)
	v200 = v183
	v203 = v198
	v205 = v188
	v207 = v198
	goto L39
L39:
	;
	v218 = v53 + v203*int32(12)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+8)))
	if v219 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v261 = v254
	v264 = v257
	v266 = v255
	goto L35
L41:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	*(*int32)(unsafe.Add(mBase, uint32(v145+v200<<(uint(int32(2))%32)))) = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v228 = int32(1)
	v229 = v227 + v228
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v229
	v233 = v200 + v228
	v234 = v229
	goto L43
L42:
	;
	v233 = v200
	v234 = v205
	goto L43
L43:
	;
	v235 = int32(1)
	v239 = v53 + (v203|v235)*int32(12)
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+8)))
	if v240 == v235 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	*(*int32)(unsafe.Add(mBase, uint32(v145+v233<<(uint(int32(2))%32)))) = v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v249 = int32(1)
	v250 = v248 + v249
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v250
	v254 = v233 + v249
	v255 = v250
	goto L46
L45:
	;
	v254 = v233
	v255 = v234
	goto L46
L46:
	;
	v256 = int32(2)
	v257 = v203 + v256
	v259 = v207 + v256
	if v259 != v137&int32(2147483646) {
		v200 = v254
		v203 = v257
		v205 = v255
		v207 = v259
		goto L39
	} else {
		goto L47
	}
L47:
	;
	goto L40
L48:
	;
	v281 = v53 + v264*int32(12)
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+8)))
	if v282 != int32(1) {
		v316 = v266
		goto L25
	} else {
		goto L49
	}
L49:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	*(*int32)(unsafe.Add(mBase, uint32(v145+v261<<(uint(int32(2))%32)))) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v309 = v290 + int32(1)
	goto L28
L50:
	;
	goto L1
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v348
	m.G0 = v19 + int32(16)
	return
}
func F_brin_summarize_range(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 float64
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = int64(0)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
	if v21 == int32(1) {
		v26 = *(*int32)(unsafe.Add(mBase, _consts[2]))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+316))
		v29 = base.B2i32(v27 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v29)
		v31 = v29
	} else {
		v31 = v2
	}
	if v31 == int32(0) {
		if base.Ui64(int64(4294967296)) <= base.Ui64(v15) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v191 = m.ExcPending
			if v191 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v12))) = v15
					F_errmsg(m, int32(430401), v12)
					mBase = m.M
					v198 = m.ExcPending
					if v198 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(496462), int32(1403), int32(401663))
						mBase = m.M
						v203 = m.ExcPending
						if v203 != 0 {
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
			v37 = F_IndexGetRelation(m, v16, int32(1))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				if v37 != 0 {
					v42 = F_table_open(m, v37, int32(4))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, _consts[3]))
						*(*int32)(unsafe.Add(mBase, uint32(v12+int32(76)))) = v49
						v52 = *(*int32)(unsafe.Add(mBase, _consts[4]))
						*(*int32)(unsafe.Add(mBase, uint32(v12+int32(72)))) = v52
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+48))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+80))
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
						*(*int32)(unsafe.Add(mBase, _consts[4])) = v56 | int32(2)
						*(*int32)(unsafe.Add(mBase, _consts[3])) = v55
						v64 = int32(4513432)
						v66 = *(*int32)(unsafe.Add(mBase, _consts[5]))
						v68 = v66 + int32(1)
						*(*int32)(unsafe.Add(mBase, _consts[5])) = v68
						F_RestrictSearchPath(m)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							v77 = v42
							v78 = v68
							v80 = F_index_open(m, v16, int32(4))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
								v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+119)))
								if v83 != int32(105) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v207 = m.ExcPending
									if v207 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v210 = m.ExcPending
										if v210 != 0 {
											return int32(0)
										} else {
											v211 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v211 + int32(4)
											F_errmsg(m, int32(29085), v12+int32(48))
											mBase = m.M
											v219 = m.ExcPending
											if v219 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(496462), int32(1448), int32(401663))
												mBase = m.M
												v224 = m.ExcPending
												if v224 != 0 {
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
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+84))
									if v86 != int32(3580) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v207 = m.ExcPending
										if v207 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(151027844))
											mBase = m.M
											v210 = m.ExcPending
											if v210 != 0 {
												return int32(0)
											} else {
												v211 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v211 + int32(4)
												F_errmsg(m, int32(29085), v12+int32(48))
												mBase = m.M
												v219 = m.ExcPending
												if v219 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(496462), int32(1448), int32(401663))
													mBase = m.M
													v224 = m.ExcPending
													if v224 != 0 {
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
										if v77 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v228 = m.ExcPending
											if v228 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(16908420))
												mBase = m.M
												v231 = m.ExcPending
												if v231 != 0 {
													return int32(0)
												} else {
													v232 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v232 + int32(4)
													F_errmsg(m, int32(696125), v12+int32(16))
													mBase = m.M
													v240 = m.ExcPending
													if v240 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(496462), int32(1464), int32(401663))
														mBase = m.M
														v245 = m.ExcPending
														if v245 != 0 {
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
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
											v93 = F_object_ownercheck(m, int32(1259), v16, v92)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												if v93 == int32(0) {
													v99 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
													F_aclcheck_error(m, int32(2), int32(20), v99+int32(4))
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														v105 = F_IndexGetRelation(m, v16, int32(0))
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return int32(0)
														} else {
															if v105 != v37 {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v228 = m.ExcPending
																if v228 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(16908420))
																	mBase = m.M
																	v231 = m.ExcPending
																	if v231 != 0 {
																		return int32(0)
																	} else {
																		v232 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v232 + int32(4)
																		F_errmsg(m, int32(696125), v12+int32(16))
																		mBase = m.M
																		v240 = m.ExcPending
																		if v240 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(496462), int32(1464), int32(401663))
																			mBase = m.M
																			v245 = m.ExcPending
																			if v245 != 0 {
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
																v108 = *(*int32)(unsafe.Add(mBase, uint32(v80)+192))
																v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+18)))
																if v109 == int32(1) {
																	F_brinsummarize(m, v80, v77, base.I32_wrap_i64(v15), int32(1), v12-int32(-64), int32(0))
																	mBase = m.M
																	v118 = m.ExcPending
																	if v118 != 0 {
																		return int32(0)
																	} else {
																		F_AtEOXact_GUC(m, int32(0), v78)
																		mBase = m.M
																		v144 = m.ExcPending
																		if v144 != 0 {
																			return int32(0)
																		} else {
																			v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
																			v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
																			*(*int32)(unsafe.Add(mBase, _consts[4])) = v146
																			*(*int32)(unsafe.Add(mBase, _consts[3])) = v145
																			F_relation_close(m, v80, int32(4))
																			mBase = m.M
																			v153 = m.ExcPending
																			if v153 != 0 {
																				return int32(0)
																			} else {
																				F_relation_close(m, v77, int32(4))
																				mBase = m.M
																				v156 = m.ExcPending
																				if v156 != 0 {
																					return int32(0)
																				} else {
																					v157 = *(*float64)(unsafe.Add(mBase, uint32(v12)+64))
																					if base.F64_lt(base.F64_abs(v157), float64(2.147483648e+09)) != 0 {
																						v161 = base.I32_trunc_f64_s(v157)
																						v163 = v161
																					} else {
																						v163 = int32(-2147483648)
																					}
																					m.G0 = v12 + int32(80)
																					return v163
																				}
																			}
																		}
																	}
																} else {
																	v121 = F_errstart(m, int32(14), int32(0))
																	mBase = m.M
																	v122 = m.ExcPending
																	if v122 != 0 {
																		return int32(0)
																	} else {
																		if v121 == int32(0) {
																			F_AtEOXact_GUC(m, int32(0), v78)
																			mBase = m.M
																			v144 = m.ExcPending
																			if v144 != 0 {
																				return int32(0)
																			} else {
																				v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
																				v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
																				*(*int32)(unsafe.Add(mBase, _consts[4])) = v146
																				*(*int32)(unsafe.Add(mBase, _consts[3])) = v145
																				F_relation_close(m, v80, int32(4))
																				mBase = m.M
																				v153 = m.ExcPending
																				if v153 != 0 {
																					return int32(0)
																				} else {
																					F_relation_close(m, v77, int32(4))
																					mBase = m.M
																					v156 = m.ExcPending
																					if v156 != 0 {
																						return int32(0)
																					} else {
																						v157 = *(*float64)(unsafe.Add(mBase, uint32(v12)+64))
																						if base.F64_lt(base.F64_abs(v157), float64(2.147483648e+09)) != 0 {
																							v161 = base.I32_trunc_f64_s(v157)
																							v163 = v161
																						} else {
																							v163 = int32(-2147483648)
																						}
																						m.G0 = v12 + int32(80)
																						return v163
																					}
																				}
																			}
																		} else {
																			F_errcode(m, int32(325))
																			mBase = m.M
																			v127 = m.ExcPending
																			if v127 != 0 {
																				return int32(0)
																			} else {
																				v128 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v128 + int32(4)
																				F_errmsg(m, int32(435866), v12+int32(32))
																				mBase = m.M
																				v136 = m.ExcPending
																				if v136 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(496462), int32(1473), int32(401663))
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						F_AtEOXact_GUC(m, int32(0), v78)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return int32(0)
																						} else {
																							v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
																							v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
																							*(*int32)(unsafe.Add(mBase, _consts[4])) = v146
																							*(*int32)(unsafe.Add(mBase, _consts[3])) = v145
																							F_relation_close(m, v80, int32(4))
																							mBase = m.M
																							v153 = m.ExcPending
																							if v153 != 0 {
																								return int32(0)
																							} else {
																								F_relation_close(m, v77, int32(4))
																								mBase = m.M
																								v156 = m.ExcPending
																								if v156 != 0 {
																									return int32(0)
																								} else {
																									v157 = *(*float64)(unsafe.Add(mBase, uint32(v12)+64))
																									if base.F64_lt(base.F64_abs(v157), float64(2.147483648e+09)) != 0 {
																										v161 = base.I32_trunc_f64_s(v157)
																										v163 = v161
																									} else {
																										v163 = int32(-2147483648)
																									}
																									m.G0 = v12 + int32(80)
																									return v163
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
													v105 = F_IndexGetRelation(m, v16, int32(0))
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return int32(0)
													} else {
														if v105 != v37 {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v228 = m.ExcPending
															if v228 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(16908420))
																mBase = m.M
																v231 = m.ExcPending
																if v231 != 0 {
																	return int32(0)
																} else {
																	v232 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
																	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v232 + int32(4)
																	F_errmsg(m, int32(696125), v12+int32(16))
																	mBase = m.M
																	v240 = m.ExcPending
																	if v240 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(496462), int32(1464), int32(401663))
																		mBase = m.M
																		v245 = m.ExcPending
																		if v245 != 0 {
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
															v108 = *(*int32)(unsafe.Add(mBase, uint32(v80)+192))
															v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+18)))
															if v109 == int32(1) {
																F_brinsummarize(m, v80, v77, base.I32_wrap_i64(v15), int32(1), v12-int32(-64), int32(0))
																mBase = m.M
																v118 = m.ExcPending
																if v118 != 0 {
																	return int32(0)
																} else {
																	F_AtEOXact_GUC(m, int32(0), v78)
																	mBase = m.M
																	v144 = m.ExcPending
																	if v144 != 0 {
																		return int32(0)
																	} else {
																		v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
																		v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
																		*(*int32)(unsafe.Add(mBase, _consts[4])) = v146
																		*(*int32)(unsafe.Add(mBase, _consts[3])) = v145
																		F_relation_close(m, v80, int32(4))
																		mBase = m.M
																		v153 = m.ExcPending
																		if v153 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v77, int32(4))
																			mBase = m.M
																			v156 = m.ExcPending
																			if v156 != 0 {
																				return int32(0)
																			} else {
																				v157 = *(*float64)(unsafe.Add(mBase, uint32(v12)+64))
																				if base.F64_lt(base.F64_abs(v157), float64(2.147483648e+09)) != 0 {
																					v161 = base.I32_trunc_f64_s(v157)
																					v163 = v161
																				} else {
																					v163 = int32(-2147483648)
																				}
																				m.G0 = v12 + int32(80)
																				return v163
																			}
																		}
																	}
																}
															} else {
																v121 = F_errstart(m, int32(14), int32(0))
																mBase = m.M
																v122 = m.ExcPending
																if v122 != 0 {
																	return int32(0)
																} else {
																	if v121 == int32(0) {
																		F_AtEOXact_GUC(m, int32(0), v78)
																		mBase = m.M
																		v144 = m.ExcPending
																		if v144 != 0 {
																			return int32(0)
																		} else {
																			v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
																			v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
																			*(*int32)(unsafe.Add(mBase, _consts[4])) = v146
																			*(*int32)(unsafe.Add(mBase, _consts[3])) = v145
																			F_relation_close(m, v80, int32(4))
																			mBase = m.M
																			v153 = m.ExcPending
																			if v153 != 0 {
																				return int32(0)
																			} else {
																				F_relation_close(m, v77, int32(4))
																				mBase = m.M
																				v156 = m.ExcPending
																				if v156 != 0 {
																					return int32(0)
																				} else {
																					v157 = *(*float64)(unsafe.Add(mBase, uint32(v12)+64))
																					if base.F64_lt(base.F64_abs(v157), float64(2.147483648e+09)) != 0 {
																						v161 = base.I32_trunc_f64_s(v157)
																						v163 = v161
																					} else {
																						v163 = int32(-2147483648)
																					}
																					m.G0 = v12 + int32(80)
																					return v163
																				}
																			}
																		}
																	} else {
																		F_errcode(m, int32(325))
																		mBase = m.M
																		v127 = m.ExcPending
																		if v127 != 0 {
																			return int32(0)
																		} else {
																			v128 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v128 + int32(4)
																			F_errmsg(m, int32(435866), v12+int32(32))
																			mBase = m.M
																			v136 = m.ExcPending
																			if v136 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(496462), int32(1473), int32(401663))
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					F_AtEOXact_GUC(m, int32(0), v78)
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return int32(0)
																					} else {
																						v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
																						v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
																						*(*int32)(unsafe.Add(mBase, _consts[4])) = v146
																						*(*int32)(unsafe.Add(mBase, _consts[3])) = v145
																						F_relation_close(m, v80, int32(4))
																						mBase = m.M
																						v153 = m.ExcPending
																						if v153 != 0 {
																							return int32(0)
																						} else {
																							F_relation_close(m, v77, int32(4))
																							mBase = m.M
																							v156 = m.ExcPending
																							if v156 != 0 {
																								return int32(0)
																							} else {
																								v157 = *(*float64)(unsafe.Add(mBase, uint32(v12)+64))
																								if base.F64_lt(base.F64_abs(v157), float64(2.147483648e+09)) != 0 {
																									v161 = base.I32_trunc_f64_s(v157)
																									v163 = v161
																								} else {
																									v163 = int32(-2147483648)
																								}
																								m.G0 = v12 + int32(80)
																								return v163
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
					v72 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v72
					*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = int32(0)
					v77 = v2
					v78 = v72
					v80 = F_index_open(m, v16, int32(4))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
						v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+119)))
						if v83 != int32(105) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v207 = m.ExcPending
							if v207 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(151027844))
								mBase = m.M
								v210 = m.ExcPending
								if v210 != 0 {
									return int32(0)
								} else {
									v211 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v211 + int32(4)
									F_errmsg(m, int32(29085), v12+int32(48))
									mBase = m.M
									v219 = m.ExcPending
									if v219 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(496462), int32(1448), int32(401663))
										mBase = m.M
										v224 = m.ExcPending
										if v224 != 0 {
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
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+84))
							if v86 != int32(3580) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v207 = m.ExcPending
								if v207 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(151027844))
									mBase = m.M
									v210 = m.ExcPending
									if v210 != 0 {
										return int32(0)
									} else {
										v211 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v211 + int32(4)
										F_errmsg(m, int32(29085), v12+int32(48))
										mBase = m.M
										v219 = m.ExcPending
										if v219 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496462), int32(1448), int32(401663))
											mBase = m.M
											v224 = m.ExcPending
											if v224 != 0 {
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
								if v77 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v228 = m.ExcPending
									if v228 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(16908420))
										mBase = m.M
										v231 = m.ExcPending
										if v231 != 0 {
											return int32(0)
										} else {
											v232 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v232 + int32(4)
											F_errmsg(m, int32(696125), v12+int32(16))
											mBase = m.M
											v240 = m.ExcPending
											if v240 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(496462), int32(1464), int32(401663))
												mBase = m.M
												v245 = m.ExcPending
												if v245 != 0 {
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
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
									v93 = F_object_ownercheck(m, int32(1259), v16, v92)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										if v93 == int32(0) {
											v99 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
											F_aclcheck_error(m, int32(2), int32(20), v99+int32(4))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												v105 = F_IndexGetRelation(m, v16, int32(0))
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													if v105 != v37 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v228 = m.ExcPending
														if v228 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(16908420))
															mBase = m.M
															v231 = m.ExcPending
															if v231 != 0 {
																return int32(0)
															} else {
																v232 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
																*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v232 + int32(4)
																F_errmsg(m, int32(696125), v12+int32(16))
																mBase = m.M
																v240 = m.ExcPending
																if v240 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(496462), int32(1464), int32(401663))
																	mBase = m.M
																	v245 = m.ExcPending
																	if v245 != 0 {
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
														v108 = *(*int32)(unsafe.Add(mBase, uint32(v80)+192))
														v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+18)))
														if v109 == int32(1) {
															F_brinsummarize(m, v80, v77, base.I32_wrap_i64(v15), int32(1), v12-int32(-64), int32(0))
															mBase = m.M
															v118 = m.ExcPending
															if v118 != 0 {
																return int32(0)
															} else {
																F_AtEOXact_GUC(m, int32(0), v78)
																mBase = m.M
																v144 = m.ExcPending
																if v144 != 0 {
																	return int32(0)
																} else {
																	v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
																	v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
																	*(*int32)(unsafe.Add(mBase, _consts[4])) = v146
																	*(*int32)(unsafe.Add(mBase, _consts[3])) = v145
																	F_relation_close(m, v80, int32(4))
																	mBase = m.M
																	v153 = m.ExcPending
																	if v153 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v77, int32(4))
																		mBase = m.M
																		v156 = m.ExcPending
																		if v156 != 0 {
																			return int32(0)
																		} else {
																			v157 = *(*float64)(unsafe.Add(mBase, uint32(v12)+64))
																			if base.F64_lt(base.F64_abs(v157), float64(2.147483648e+09)) != 0 {
																				v161 = base.I32_trunc_f64_s(v157)
																				v163 = v161
																			} else {
																				v163 = int32(-2147483648)
																			}
																			m.G0 = v12 + int32(80)
																			return v163
																		}
																	}
																}
															}
														} else {
															v121 = F_errstart(m, int32(14), int32(0))
															mBase = m.M
															v122 = m.ExcPending
															if v122 != 0 {
																return int32(0)
															} else {
																if v121 == int32(0) {
																	F_AtEOXact_GUC(m, int32(0), v78)
																	mBase = m.M
																	v144 = m.ExcPending
																	if v144 != 0 {
																		return int32(0)
																	} else {
																		v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
																		v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
																		*(*int32)(unsafe.Add(mBase, _consts[4])) = v146
																		*(*int32)(unsafe.Add(mBase, _consts[3])) = v145
																		F_relation_close(m, v80, int32(4))
																		mBase = m.M
																		v153 = m.ExcPending
																		if v153 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v77, int32(4))
																			mBase = m.M
																			v156 = m.ExcPending
																			if v156 != 0 {
																				return int32(0)
																			} else {
																				v157 = *(*float64)(unsafe.Add(mBase, uint32(v12)+64))
																				if base.F64_lt(base.F64_abs(v157), float64(2.147483648e+09)) != 0 {
																					v161 = base.I32_trunc_f64_s(v157)
																					v163 = v161
																				} else {
																					v163 = int32(-2147483648)
																				}
																				m.G0 = v12 + int32(80)
																				return v163
																			}
																		}
																	}
																} else {
																	F_errcode(m, int32(325))
																	mBase = m.M
																	v127 = m.ExcPending
																	if v127 != 0 {
																		return int32(0)
																	} else {
																		v128 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v128 + int32(4)
																		F_errmsg(m, int32(435866), v12+int32(32))
																		mBase = m.M
																		v136 = m.ExcPending
																		if v136 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(496462), int32(1473), int32(401663))
																			mBase = m.M
																			v141 = m.ExcPending
																			if v141 != 0 {
																				return int32(0)
																			} else {
																				F_AtEOXact_GUC(m, int32(0), v78)
																				mBase = m.M
																				v144 = m.ExcPending
																				if v144 != 0 {
																					return int32(0)
																				} else {
																					v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
																					v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
																					*(*int32)(unsafe.Add(mBase, _consts[4])) = v146
																					*(*int32)(unsafe.Add(mBase, _consts[3])) = v145
																					F_relation_close(m, v80, int32(4))
																					mBase = m.M
																					v153 = m.ExcPending
																					if v153 != 0 {
																						return int32(0)
																					} else {
																						F_relation_close(m, v77, int32(4))
																						mBase = m.M
																						v156 = m.ExcPending
																						if v156 != 0 {
																							return int32(0)
																						} else {
																							v157 = *(*float64)(unsafe.Add(mBase, uint32(v12)+64))
																							if base.F64_lt(base.F64_abs(v157), float64(2.147483648e+09)) != 0 {
																								v161 = base.I32_trunc_f64_s(v157)
																								v163 = v161
																							} else {
																								v163 = int32(-2147483648)
																							}
																							m.G0 = v12 + int32(80)
																							return v163
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
											v105 = F_IndexGetRelation(m, v16, int32(0))
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												if v105 != v37 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v228 = m.ExcPending
													if v228 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(16908420))
														mBase = m.M
														v231 = m.ExcPending
														if v231 != 0 {
															return int32(0)
														} else {
															v232 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v232 + int32(4)
															F_errmsg(m, int32(696125), v12+int32(16))
															mBase = m.M
															v240 = m.ExcPending
															if v240 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(496462), int32(1464), int32(401663))
																mBase = m.M
																v245 = m.ExcPending
																if v245 != 0 {
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
													v108 = *(*int32)(unsafe.Add(mBase, uint32(v80)+192))
													v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+18)))
													if v109 == int32(1) {
														F_brinsummarize(m, v80, v77, base.I32_wrap_i64(v15), int32(1), v12-int32(-64), int32(0))
														mBase = m.M
														v118 = m.ExcPending
														if v118 != 0 {
															return int32(0)
														} else {
															F_AtEOXact_GUC(m, int32(0), v78)
															mBase = m.M
															v144 = m.ExcPending
															if v144 != 0 {
																return int32(0)
															} else {
																v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
																v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
																*(*int32)(unsafe.Add(mBase, _consts[4])) = v146
																*(*int32)(unsafe.Add(mBase, _consts[3])) = v145
																F_relation_close(m, v80, int32(4))
																mBase = m.M
																v153 = m.ExcPending
																if v153 != 0 {
																	return int32(0)
																} else {
																	F_relation_close(m, v77, int32(4))
																	mBase = m.M
																	v156 = m.ExcPending
																	if v156 != 0 {
																		return int32(0)
																	} else {
																		v157 = *(*float64)(unsafe.Add(mBase, uint32(v12)+64))
																		if base.F64_lt(base.F64_abs(v157), float64(2.147483648e+09)) != 0 {
																			v161 = base.I32_trunc_f64_s(v157)
																			v163 = v161
																		} else {
																			v163 = int32(-2147483648)
																		}
																		m.G0 = v12 + int32(80)
																		return v163
																	}
																}
															}
														}
													} else {
														v121 = F_errstart(m, int32(14), int32(0))
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return int32(0)
														} else {
															if v121 == int32(0) {
																F_AtEOXact_GUC(m, int32(0), v78)
																mBase = m.M
																v144 = m.ExcPending
																if v144 != 0 {
																	return int32(0)
																} else {
																	v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
																	v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
																	*(*int32)(unsafe.Add(mBase, _consts[4])) = v146
																	*(*int32)(unsafe.Add(mBase, _consts[3])) = v145
																	F_relation_close(m, v80, int32(4))
																	mBase = m.M
																	v153 = m.ExcPending
																	if v153 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v77, int32(4))
																		mBase = m.M
																		v156 = m.ExcPending
																		if v156 != 0 {
																			return int32(0)
																		} else {
																			v157 = *(*float64)(unsafe.Add(mBase, uint32(v12)+64))
																			if base.F64_lt(base.F64_abs(v157), float64(2.147483648e+09)) != 0 {
																				v161 = base.I32_trunc_f64_s(v157)
																				v163 = v161
																			} else {
																				v163 = int32(-2147483648)
																			}
																			m.G0 = v12 + int32(80)
																			return v163
																		}
																	}
																}
															} else {
																F_errcode(m, int32(325))
																mBase = m.M
																v127 = m.ExcPending
																if v127 != 0 {
																	return int32(0)
																} else {
																	v128 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
																	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v128 + int32(4)
																	F_errmsg(m, int32(435866), v12+int32(32))
																	mBase = m.M
																	v136 = m.ExcPending
																	if v136 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(496462), int32(1473), int32(401663))
																		mBase = m.M
																		v141 = m.ExcPending
																		if v141 != 0 {
																			return int32(0)
																		} else {
																			F_AtEOXact_GUC(m, int32(0), v78)
																			mBase = m.M
																			v144 = m.ExcPending
																			if v144 != 0 {
																				return int32(0)
																			} else {
																				v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
																				v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
																				*(*int32)(unsafe.Add(mBase, _consts[4])) = v146
																				*(*int32)(unsafe.Add(mBase, _consts[3])) = v145
																				F_relation_close(m, v80, int32(4))
																				mBase = m.M
																				v153 = m.ExcPending
																				if v153 != 0 {
																					return int32(0)
																				} else {
																					F_relation_close(m, v77, int32(4))
																					mBase = m.M
																					v156 = m.ExcPending
																					if v156 != 0 {
																						return int32(0)
																					} else {
																						v157 = *(*float64)(unsafe.Add(mBase, uint32(v12)+64))
																						if base.F64_lt(base.F64_abs(v157), float64(2.147483648e+09)) != 0 {
																							v161 = base.I32_trunc_f64_s(v157)
																							v163 = v161
																						} else {
																							v163 = int32(-2147483648)
																						}
																						m.G0 = v12 + int32(80)
																						return v163
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
		v171 = m.ExcPending
		if v171 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v174 = m.ExcPending
			if v174 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(128046), int32(0))
				mBase = m.M
				v178 = m.ExcPending
				if v178 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(572935), int32(0))
					mBase = m.M
					v182 = m.ExcPending
					if v182 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(496462), int32(1398), int32(401663))
						mBase = m.M
						v187 = m.ExcPending
						if v187 != 0 {
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
