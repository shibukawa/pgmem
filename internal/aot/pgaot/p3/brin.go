package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 float64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 float64
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 float64
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v91 int32
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
				v38 = F_table_beginscan_parallel(m, l3, l1+int32(96), v30)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l3)+188))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+140))
					v42 = m.T0[v41].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l3, l4, v25, v29, v30, v29, v30, int32(-1), int32(2), l0, v38)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
						if v45 == int32(0) {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v52 = F_brin_form_tuple(m, v48, v49, v44, v12+int32(12))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
								F_tuplesort_putbrintuple(m, v54, v52, v55)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									v58 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
									*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = base.F64_add(v58, float64(1))
									F_pfree(m, v52)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
										F_tuplesort_performsort(m, v65)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											v68 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
											*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = base.F64_add(v42, v68)
											v73 = base.AtomicRmwXchg32(m, l1, int32(44), int32(1))
											if v73 != 0 {
												F_s_lock(m, l1+int32(44), int32(_a_F__brin_parallel_scan_and_build_0))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
													*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v79 + int32(1)
													v83 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
													v84 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
													*(*float64)(unsafe.Add(mBase, uint32(l1)+56)) = base.F64_add(v83, v84)
													v87 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
													v88 = *(*float64)(unsafe.Add(mBase, uint32(l1)+64))
													*(*float64)(unsafe.Add(mBase, uint32(l1)+64)) = base.F64_add(v87, v88)
													v91 = int32(0)
													atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l1)+44)), uint32(v91))
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
												v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
												*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v79 + int32(1)
												v83 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
												v84 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
												*(*float64)(unsafe.Add(mBase, uint32(l1)+56)) = base.F64_add(v83, v84)
												v87 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
												v88 = *(*float64)(unsafe.Add(mBase, uint32(l1)+64))
												*(*float64)(unsafe.Add(mBase, uint32(l1)+64)) = base.F64_add(v87, v88)
												v91 = int32(0)
												atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l1)+44)), uint32(v91))
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
							v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							F_tuplesort_performsort(m, v65)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								v68 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
								*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = base.F64_add(v42, v68)
								v73 = base.AtomicRmwXchg32(m, l1, int32(44), int32(1))
								if v73 != 0 {
									F_s_lock(m, l1+int32(44), int32(_a_F__brin_parallel_scan_and_build_0))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return
									} else {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v79 + int32(1)
										v83 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
										v84 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
										*(*float64)(unsafe.Add(mBase, uint32(l1)+56)) = base.F64_add(v83, v84)
										v87 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
										v88 = *(*float64)(unsafe.Add(mBase, uint32(l1)+64))
										*(*float64)(unsafe.Add(mBase, uint32(l1)+64)) = base.F64_add(v87, v88)
										v91 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l1)+44)), uint32(v91))
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
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v79 + int32(1)
									v83 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
									v84 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
									*(*float64)(unsafe.Add(mBase, uint32(l1)+56)) = base.F64_add(v83, v84)
									v87 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
									v88 = *(*float64)(unsafe.Add(mBase, uint32(l1)+64))
									*(*float64)(unsafe.Add(mBase, uint32(l1)+64)) = base.F64_add(v87, v88)
									v91 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l1)+44)), uint32(v91))
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
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
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
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
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
	return v279
L2:
	;
	v25 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l3))) = uint16(v25)
	v279 = v25
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
	v33 = v17 * v18
	v34 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)) = uint16(v34)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(-1)
	goto L8
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[0]))
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
	F_LockBufferInternal(m, v89, int32(1))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L13
	} else {
		goto L26
	}
L16:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v86 = F_ReadBuffer(m, v85, v32)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
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
		v89 = v78
		goto L15
	} else {
		goto L22
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[1]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62+(v56^int32(-1))*int32(56))+16))
	v77 = v68
	goto L18
L20:
	;
	goto L21
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[2]))
	v71 = int32(56)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70+v56*v71-v71)+16))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v86
	v89 = v86
	goto L15
L26:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v93 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v113 = base.I32_div_u_s(v33, v112)
	v115 = base.I32_rem_u_s(v113, int32(1360))
	v118 = v111 + v115*int32(6)
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118)+28)))
	if v119 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[3]))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v97+(v93^int32(-1))<<(uint(int32(2))%32))))
	v111 = v103
	goto L27
L29:
	;
	goto L30
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[4]))
	v111 = v105 + v93<<(uint(int32(13))%32) + int32(-8192)
	goto L27
L31:
	;
	F_UnlockBuffer(m, v93)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L13
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v126 = v118 + int32(24)
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)))
	if v127 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v279 = int32(0)
	goto L1
L35:
	;
	F_UnlockBuffer(m, v199)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L13
	} else {
		goto L79
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
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
	F_UnlockBuffer(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
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
	v160 = v151 | v152<<(uint(int32(16))%32)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v161 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	F_LockBufferInternal(m, v194, int32(1))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L13
	} else {
		goto L59
	}
L49:
	;
	v191 = F_ReadBuffer(m, v28, v160)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L13
	} else {
		goto L58
	}
L50:
	;
	if v161 < int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v160 == v182 {
		v194 = v183
		goto L48
	} else {
		goto L55
	}
L52:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[1]))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v167+(v161^int32(-1))*int32(56))+16))
	v182 = v173
	goto L51
L53:
	;
	goto L54
L54:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[2]))
	v176 = int32(56)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v175+v161*v176-v176)+16))
	v182 = v181
	goto L51
L55:
	;
	if v183 == int32(0) {
		goto L49
	} else {
		goto L56
	}
L56:
	;
	F_ReleaseBuffer(m, v183)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	goto L49
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v191
	v194 = v191
	goto L48
L59:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v199 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+16)))
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218+v217)+6)))
	if v220 != int32(_a_F_brinGetTupleForHeapBlock_0) {
		goto L35
	} else {
		goto L64
	}
L61:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[3]))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v203+(v199^int32(-1))<<(uint(int32(2))%32))))
	v217 = v209
	goto L60
L62:
	;
	goto L63
L63:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_brinGetTupleForHeapBlock[4]))
	v217 = v211 + v199<<(uint(int32(13))%32) + int32(-8192)
	goto L60
L64:
	;
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v224) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v232 = int32(base.Ui32(v224+int32(_a_F_brinGetTupleForHeapBlock_1)) >> (uint(int32(2)) % 32))
	goto L67
L66:
	;
	v232 = int32(0)
	goto L67
L67:
	;
	if base.Ui32(v232&int32(_a_F_brinGetTupleForHeapBlock_2)) < base.Ui32(v223) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	F_UnlockBuffer(m, v199)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L13
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v217+v223<<(uint(int32(2))%32))+20))
	if v242&int32(_a_F_brinGetTupleForHeapBlock_3) == int32(0) {
		goto L35
	} else {
		goto L72
	}
L71:
	;
	v279 = int32(0)
	goto L1
L72:
	;
	v249 = v217 + v242&int32(_a_F_brinGetTupleForHeapBlock_4)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	if v250 != v33 {
		goto L35
	} else {
		goto L73
	}
L73:
	;
	if l4 == int32(0) {
		v279 = v249
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(base.Ui32(v242) >> (uint(int32(17)) % 32))
	v279 = v249
	goto L1
L75:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L13
	} else {
		goto L76
	}
L76:
	;
	F_errmsg_internal(m, int32(_a_F_brinGetTupleForHeapBlock_5), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_brinGetTupleForHeapBlock_6), int32(259), int32(_a_F_brinGetTupleForHeapBlock_7))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
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
func F_brin_bloom_summary_out(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v16 = v8 + int32(16)
		F_initStringInfo(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			F_appendStringInfoChar(m, v16, int32(123))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+6)))
				v23 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v22
				F_appendStringInfo(m, v16, int32(_a_F_brin_bloom_summary_out_0), v8)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					F_appendStringInfoChar(m, v16, int32(125))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						v32 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+16)))
						m.G0 = v8 + int32(32)
						return v32
					}
				}
			}
		}
	}
}
func F_brin_bloom_summary_recv(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_brin_bloom_summary_recv_0), int32(828), int32(_a_F_brin_bloom_summary_recv_1), int32(_a_F_brin_bloom_summary_recv_2), int32(_a_F_brin_bloom_summary_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_brin_bloom_summary_send(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_byteasend(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v201 int32
	_ = v201
	var v205 int64
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int64
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int64
	_ = v247
	var v248 int32
	_ = v248
	var v250 int64
	_ = v250
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int64
	_ = v264
	var v271 int64
	_ = v271
	var v283 int32
	_ = v283
	var v286 int64
	_ = v286
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
	var v305 int32
	_ = v305
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v346 int32
	_ = v346
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v431 int32
	_ = v431
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	v5 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(16)
	m.G0 = v33
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+14)) = uint16(v5)
	v37 = int32(8)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v40 = F_palloc_mul(m, v37, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v46 = F_palloc0_mul(m, int32(1), v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v53 = base.I32_div_s(v49+int32(7), int32(8))
	v54 = F_palloc_mul(m, int32(1), v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v58 = F_palloc_mul(m, int32(8), v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v61 <= int32(0) {
		v356 = v5
		v362 = v37
		v366 = v5
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v381 = F_brtuple_disk_tupdesc(m, l0)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L54
	}
L7:
	;
	v65 = l0 + int32(20)
	v73 = v5
	v76 = v5
	v82 = v5
	v83 = v5
	goto L8
L8:
	;
	v100 = l2 + int32(24) + v76*int32(24)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+3)))
	if v101 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v335 = int32(1)
	if v314&v335 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L10:
	;
	v331 = v76 + int32(1)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	if v331 < v333 {
		v73 = v305
		v76 = v331
		v82 = v314
		v83 = v315
		goto L8
	} else {
		goto L50
	}
L11:
	;
	v104 = int32(0)
	v107 = v65 + v76<<(uint(int32(2))%32)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108))))
	if v109 == v104 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	if v155 != 0 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v305 = v73
	v314 = int32(1)
	v315 = v83
	goto L10
L15:
	;
	goto L16
L16:
	;
	v118 = v73
	v120 = v104
	goto L17
L17:
	;
	v143 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v118+v46))) = uint8(v143)
	v148 = v118 + v143
	v150 = v120 + v143
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151))))
	if base.Ui32(v150) < base.Ui32(v152) {
		v118 = v148
		v120 = v150
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v305 = v148
	v314 = v143
	v315 = v83
	goto L10
L19:
	;
	goto L18
L20:
	;
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v100)+8))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	m.T0[v155].(func(*base.Module, int32, int64, int32))(m, l0, v156, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v160 = v154 | v82
	v163 = v65 + v76<<(uint(int32(2))%32)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164))))
	if v165 == int32(0) {
		v305 = v73
		v314 = v160
		v315 = v83
		goto L10
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v175 = v164
	v176 = v73
	v178 = int32(0)
	v186 = v83
	goto L25
L25:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v201+v178<<(uint(int32(3))%32))))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v175+v178<<(uint(int32(2))%32))+8))
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209)+8)))
	if v210 != int32(_a_F_brin_form_tuple_0) {
		v283 = v186
		v286 = v205
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v305 = v294
	v314 = v160
	v315 = v283
	goto L10
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40+v176<<(uint(int32(3))%32)))) = v286
	v293 = int32(1)
	v294 = v176 + v293
	v296 = v178 + v293
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297))))
	if base.Ui32(v296) < base.Ui32(v298) {
		v175 = v297
		v176 = v294
		v178 = v296
		v186 = v283
		goto L25
	} else {
		goto L49
	}
L28:
	;
	v213 = base.I32_wrap_i64(v205)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	v216 = base.B2i32(v214 != int32(1))
	if v214 != int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v225 = base.B2i32(v214 == int32(1))
	if v221&int32(3) != 0 {
		v263 = v225
		v264 = v223
		goto L35
	} else {
		goto L36
	}
L30:
	;
	v221 = v214
	v222 = v213
	v223 = v205
	goto L29
L31:
	;
	goto L32
L32:
	;
	v217 = F_detoast_external_attr(m, v213)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	v221 = v219
	v222 = v217
	v223 = base.I64_extend_i32_u(v217)
	goto L29
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v58+v186<<(uint(int32(3))%32)))) = v271
	v283 = v186 + int32(1)
	v286 = v271
	goto L27
L35:
	;
	if v263 == int32(0) {
		v283 = v186
		v286 = v264
		goto L27
	} else {
		goto L48
	}
L36:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	if base.Ui32(v228) < base.Ui32(int32(2044)) {
		v263 = v225
		v264 = v223
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+12)))
	switch v231 - int32(109) {
	case 0, 11:
		goto L38
	default:
		v263 = v225
		v264 = v223
		goto L35
	}
L38:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v239 = v234 + v235<<(uint(int32(3))%32) + v76*int32(100)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+96))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v240 == v241 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+113)))
	v245 = v243
	goto L41
L40:
	;
	v245 = int32(0)
	goto L41
L41:
	;
	v247 = F_toast_compress_datum(m, v223, base.I32_extend8_s(v245))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v250 = v247 & int64(4294967295)
	v252 = base.B2i32(v250 != int64(0))
	if v250 != int64(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v253 = v247
	goto L45
L44:
	;
	v253 = v223
	goto L45
L45:
	;
	if v216|base.B2i32(v250 == int64(0)) != 0 {
		v263 = base.B2i32(v214 == int32(1)) | v252
		v264 = v253
		goto L35
	} else {
		goto L46
	}
L46:
	;
	F_pfree(m, v222)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v271 = v247
	goto L34
L48:
	;
	v271 = v264
	goto L34
L49:
	;
	goto L26
L50:
	;
	goto L9
L51:
	;
	v356 = int32(0)
	v362 = v37
	v366 = v315
	goto L6
L52:
	;
	goto L53
L53:
	;
	v346 = base.I32_div_s(v333<<(uint(int32(1))%32)+int32(7), int32(8))
	v356 = v335
	v362 = (v346 + int32(12)) & int32(-8)
	v366 = v315
	goto L6
L54:
	;
	v383 = F_heap_compute_data_size(m, v381, v40, v46)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v389 = (v383 + v362 + int32(7)) & int32(-8)
	v390 = F_palloc0(m, v389)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v390)+4)) = uint8(v362)
	*(*int32)(unsafe.Add(mBase, uint32(v390))) = l1
	v394 = F_brtuple_disk_tupdesc(m, l0)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_heap_fill_tuple(m, v394, v40, v46, v390+v362, v33+int32(14), v54)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_pfree(m, v40)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_pfree(m, v46)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_pfree(m, v54)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if int32(0) < v366 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v431 = v5
	goto L65
L63:
	;
	goto L64
L64:
	;
	v479 = v390 + int32(4)
	if v356 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v58+v431<<(uint(int32(3))%32))))
	F_pfree(m, v442)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	goto L64
L67:
	;
	v446 = v431 + int32(1)
	if v446 != v366 {
		v431 = v446
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v639 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L70:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	v484 = v482 | int32(-128)
	*(*uint8)(unsafe.Add(mBase, uint32(v479))) = uint8(v484)
	v486 = int32(0)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)))
	if v488 <= v486 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v496 = v479
	v497 = int32(128)
	v498 = v484
	v499 = v486
	goto L72
L72:
	;
	if v497 != int32(128) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v549 = int32(0)
	if v547 <= v549 {
		goto L69
	} else {
		goto L81
	}
L74:
	;
	v532 = v496
	v533 = v498
	v534 = v497 << (uint(int32(1)) % 32)
	goto L76
L75:
	;
	v526 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v496)+1)) = uint8(v526)
	v529 = int32(1)
	v532 = v496 + v529
	v533 = v526
	v534 = v529
	goto L76
L76:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v499*int32(24))+27)))
	if v538 == int32(1) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v541 = v534 | v533
	*(*uint8)(unsafe.Add(mBase, uint32(v532))) = uint8(v541)
	v543 = v541
	goto L79
L78:
	;
	v543 = v533
	goto L79
L79:
	;
	v545 = v499 + int32(1)
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)))
	if v545 < v547 {
		v496 = v532
		v497 = v534
		v498 = v543
		v499 = v545
		goto L72
	} else {
		goto L80
	}
L80:
	;
	goto L73
L81:
	;
	v556 = v532
	v557 = v534
	v558 = v543
	v559 = v549
	goto L82
L82:
	;
	if v557 != int32(128) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	goto L69
L84:
	;
	v592 = v556
	v593 = v558
	v594 = v557 << (uint(int32(1)) % 32)
	goto L86
L85:
	;
	v586 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v556)+1)) = uint8(v586)
	v589 = int32(1)
	v592 = v556 + v589
	v593 = v586
	v594 = v589
	goto L86
L86:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v559*int32(24))+26)))
	if v598 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v601 = v594 | v593
	*(*uint8)(unsafe.Add(mBase, uint32(v592))) = uint8(v601)
	v603 = v601
	goto L89
L88:
	;
	v603 = v593
	goto L89
L89:
	;
	v605 = v559 + int32(1)
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v606)))
	if v605 < v607 {
		v556 = v592
		v557 = v594
		v558 = v603
		v559 = v605
		goto L82
	} else {
		goto L90
	}
L90:
	;
	goto L83
L91:
	;
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	v644 = v642 | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v479))) = uint8(v644)
	goto L93
L92:
	;
	goto L93
L93:
	;
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v646 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	v651 = v649 | int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v479))) = uint8(v651)
	goto L96
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v389
	m.G0 = v33 + int32(16)
	return v390
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
	var v30 int32
	_ = v30
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
	var v99 int32
	_ = v99
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
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
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
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(l1^int32(-1))*int32(56))+16))
	v36 = v27
	goto L4
L6:
	;
	goto L7
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_brin_getinsertbuffer[1]))
	v30 = int32(56)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+l1*v30-v30)+16))
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
	return v286
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
	v261 = m.ExcPending
	if v261 != 0 {
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
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90+(v85^int32(-1))*int32(56))+16))
	v105 = v96
	goto L31
L33:
	;
	goto L34
L34:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_brin_getinsertbuffer[1]))
	v99 = int32(56)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98+v85*v99-v99)+16))
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
	F_LockBufferInternal(m, v113, int32(3))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L13
	} else {
		goto L59
	}
L40:
	;
	F_LockBufferInternal(m, l1, int32(3))
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
	F_UnlockBuffer(m, l1)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v141 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_brin_initialize_empty_new_buffer(m, l0, v113)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
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
	v148 = m.ExcPending
	if v148 != 0 {
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
	v150 = m.ExcPending
	if v150 != 0 {
		goto L13
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	v151 = int32(0)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v152 != int32(1) {
		v286 = v151
		goto L15
	} else {
		goto L57
	}
L57:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v112, v112+int32(1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L13
	} else {
		goto L58
	}
L58:
	;
	v159 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v159)
	v286 = v151
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
	v167 = m.ExcPending
	if v167 != 0 {
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
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v186 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_brin_getinsertbuffer[3]))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v171+(v113^int32(-1))<<(uint(int32(2))%32))))
	v185 = v177
	goto L64
L66:
	;
	goto L67
L67:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_brin_getinsertbuffer[4]))
	v185 = v179 + v113<<(uint(int32(13))%32) + int32(-8192)
	goto L64
L68:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v244 != int32(1) {
		goto L90
	} else {
		goto L91
	}
L69:
	;
	v189 = int32(0)
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185)+16)))
	v191 = v185 + v190
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+6)))
	if v192 != int32(_a_F_brin_getinsertbuffer_0) {
		v207 = v189
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v211 != 0 {
		goto L80
	} else {
		goto L81
	}
L72:
	;
	if base.Ui32(v207) < base.Ui32(l2) {
		goto L68
	} else {
		goto L79
	}
L73:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+4)))
	if v195&int32(1) != 0 {
		v207 = v189
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v198 = int32(4)
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185)+14)))
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185)+12)))
	v201 = v199 - v200
	if v201 <= v198 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v207 = v204 - int32(4)
	goto L72
L76:
	;
	v204 = v198
	goto L78
L77:
	;
	v204 = v201
	goto L78
L78:
	;
	goto L75
L79:
	;
	goto L71
L80:
	;
	v235 = v211
	goto L82
L81:
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
		goto L83
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235)+16)) = v112
	if base.B2i32(l1 == int32(0))|base.B2i32(base.Ui32(v38) <= base.Ui32(v112)) != 0 {
		v286 = v113
		goto L15
	} else {
		goto L88
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v217
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217)+72))
	if v221 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v235 = v233
	goto L82
L85:
	;
	v229 = v221
	goto L87
L86:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v217)+76))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v217)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v222)+4)) = v223
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v217)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v217)+72))
	v229 = v227
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+72)) = v229 + int32(1)
	goto L84
L88:
	;
	F_LockBufferInternal(m, l1, int32(3))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L13
	} else {
		goto L89
	}
L89:
	;
	v286 = v113
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
	v249 = m.ExcPending
	if v249 != 0 {
		goto L13
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v250 = int32(0)
	if base.B2i32(l1 == v250)|base.B2i32(base.Ui32(v112) < base.Ui32(v38)) == v250 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L95
L97:
	;
	F_UnlockBuffer(m, l1)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L13
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v258 = F_RecordAndGetPageWithFreeSpace(m, l0, v112, v207, l2)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L13
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	v59 = v258
	goto L16
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L13
	} else {
		goto L103
	}
L103:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L13
	} else {
		goto L104
	}
L104:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v269 + int32(4)
	F_errmsg(m, int32(_a_F_brin_getinsertbuffer_1), v16+int32(16))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L13
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_brin_getinsertbuffer_2), int32(850), int32(_a_F_brin_getinsertbuffer_3))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
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
func F_brin_minmax_multi_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v158 int64
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v185 int32
	_ = v185
	var v192 int64
	_ = v192
	var v194 int32
	_ = v194
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v227 int64
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int64
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v260 int64
	_ = v260
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v29 = F_pg_detoast_datum(m, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L3
	} else {
		goto L55
	}
L2:
	;
	m.G0 = v20 + int32(32)
	return v260
L3:
	;
	return int64(0)
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v34 = F_brin_range_deserialize(m, v33, v29)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	if int32(0) < v36 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v39 = int64(1)
	if v23 <= int32(0) {
		v260 = v39
		goto L2
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v158 = int64(0)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	if v159 <= int32(0) {
		v260 = v158
		goto L2
	} else {
		goto L37
	}
L9:
	;
	v53 = int32(0)
	goto L10
L10:
	;
	v63 = v34 + int32(40) + v53<<(uint(int32(4))%32)
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v63)+8))
	v74 = int32(0)
	goto L12
L11:
	;
	goto L8
L12:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v24+v74<<(uint(int32(2))%32))))
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v87)+48))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+4)))
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+6)))
	switch v91 - int32(1) {
	case 0, 1:
		goto L19
	case 2:
		goto L18
	case 3, 4:
		goto L16
	default:
		goto L17
	}
L13:
	;
	v138 = v53 + int32(1)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	if v138 < v139 {
		v53 = v138
		goto L10
	} else {
		goto L36
	}
L14:
	;
	goto L13
L15:
	;
	v135 = v74 + int32(1)
	if v23 != v135 {
		v74 = v135
		goto L12
	} else {
		goto L35
	}
L16:
	;
	v128 = F_minmax_multi_get_strategy_procinfo(m, v25, v90, v89, v91)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L32
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L29
	}
L18:
	;
	v101 = F_minmax_multi_get_strategy_procinfo(m, v25, v90, v89, int32(5))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L3
	} else {
		goto L23
	}
L19:
	;
	v94 = F_minmax_multi_get_strategy_procinfo(m, v25, v90, v89, v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v96 = F_FunctionCall2Coll(m, v94, v22, v64, v88)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	if v96 == int64(0) {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	goto L15
L23:
	;
	v103 = F_FunctionCall2Coll(m, v101, v22, v64, v88)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	if v103 != int64(0) {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	v108 = F_minmax_multi_get_strategy_procinfo(m, v25, v90, v89, int32(1))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v110 = F_FunctionCall2Coll(m, v108, v22, v65, v88)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	if v110 == int64(0) {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	goto L14
L29:
	;
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v118
	F_errmsg_internal(m, int32(_a_F_brin_minmax_multi_consistent_0), v20)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_brin_minmax_multi_consistent_1), int32(2648), int32(_a_F_brin_minmax_multi_consistent_2))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	v130 = F_FunctionCall2Coll(m, v128, v22, v65, v88)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	if v130 == int64(0) {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	goto L15
L35:
	;
	v260 = v39
	goto L2
L36:
	;
	goto L11
L37:
	;
	v162 = int32(0)
	if v23 <= v162 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v260 = int64(1)
	goto L2
L39:
	;
	goto L40
L40:
	;
	v170 = v162
	goto L41
L41:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v34+int32(40)+v185<<(uint(int32(4))%32)+v170<<(uint(int32(3))%32))))
	v194 = int32(0)
	goto L44
L42:
	;
	v260 = v158
	goto L2
L43:
	;
	v243 = v170 + int32(1)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	if v243 < v244 {
		v170 = v243
		goto L41
	} else {
		goto L54
	}
L44:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v24+v194<<(uint(int32(2))%32))))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	if v215&int32(1) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v260 = int64(1)
	goto L2
L46:
	;
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214)+6)))
	if base.Ui32(int32(4)) < base.Ui32((v220-int32(1))&int32(_a_F_brin_minmax_multi_consistent_3)) {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v239 = v194 + int32(1)
	if v239 != v23 {
		v194 = v239
		goto L44
	} else {
		goto L53
	}
L49:
	;
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v214)+48))
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214)+4)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v214)+8))
	v230 = F_minmax_multi_get_strategy_procinfo(m, v25, v228, v229, v220)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v232 = F_FunctionCall2Coll(m, v230, v22, v192, v227)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	if v232 == int64(0) {
		goto L43
	} else {
		goto L52
	}
L52:
	;
	goto L48
L53:
	;
	goto L45
L54:
	;
	goto L42
L55:
	;
	v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v271
	F_errmsg_internal(m, int32(_a_F_brin_minmax_multi_consistent_0), v20+int32(16))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L3
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_brin_minmax_multi_consistent_1), int32(2707), int32(_a_F_brin_minmax_multi_consistent_2))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_brin_minmax_multi_distance_float8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v13 int64
	_ = v13
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v26 int64
	_ = v26
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = int64(9223372036854775807)
	v7 = v5 & v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v8&v6) {
		v13 = int64(9218868437227405312)
		if base.Ui64(v7) <= base.Ui64(v13) {
			v17 = v13
		} else {
			v17 = int64(0)
		}
		return v17
	} else {
		v19 = int64(9218868437227405312)
		if base.Ui64(v19) < base.Ui64(v7) {
			v26 = v19
		} else {
			v26 = base.I64_reinterpret_f64(base.F64_sub(base.F64_reinterpret_i64(v5), base.F64_reinterpret_i64(v8)))
		}
		return v26
	}
}
func F_brin_minmax_multi_distance_inet(m *base.Module, l0 int32) int64 {
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
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
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v169 float64
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v181 float64
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v200 int64
	_ = v200
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = int32(1)
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v21&v19 != 0 {
				v24 = v19
			} else {
				v24 = int32(4)
			}
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v24))))
			v27 = int32(1)
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			if v29&v27 != 0 {
				v32 = v27
			} else {
				v32 = int32(4)
			}
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v32))))
			if v26 == v34 {
				if v26 == int32(2) {
					v40 = int32(4)
				} else {
					v40 = int32(16)
				}
				v41 = F_palloc(m, v40)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int64(0)
				} else {
					v43 = int32(4)
					v45 = int32(1)
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
					if v47&v45 != 0 {
						v50 = v45
					} else {
						v50 = v43
					}
					v51 = v12 + v50
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
					if v52 == int32(2) {
						v55 = v43
					} else {
						v55 = int32(16)
					}
					if v55 != 0 {
						base.MemoryCopy(m, v41, v51+int32(2), v55)
					} else {
					}
					v59 = int32(4)
					v61 = int32(1)
					v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
					if v63&v61 != 0 {
						v66 = v61
					} else {
						v66 = v59
					}
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v66))))
					if v68 == int32(2) {
						v71 = v59
					} else {
						v71 = int32(16)
					}
					v72 = F_palloc(m, v71)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int64(0)
					} else {
						v74 = int32(4)
						v76 = int32(1)
						v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
						if v78&v76 != 0 {
							v81 = v76
						} else {
							v81 = v74
						}
						v82 = v17 + v81
						v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
						if v83 == int32(2) {
							v86 = v74
						} else {
							v86 = int32(16)
						}
						if v86 != 0 {
							base.MemoryCopy(m, v72, v82+int32(2), v86)
						} else {
						}
						v90 = int32(4)
						v92 = int32(1)
						v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
						if v94&v92 != 0 {
							v97 = v92
						} else {
							v97 = v90
						}
						v98 = v12 + v97
						v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
						if v99 == int32(2) {
							v102 = v90
						} else {
							v102 = int32(16)
						}
						v103 = int32(1)
						v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
						if v105&v103 != 0 {
							v108 = v103
						} else {
							v108 = int32(4)
						}
						v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v108)+1)))
						v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
						v113 = int32(0)
						for {
							v124 = v113 << (uint(int32(3)) % 32)
							v125 = v111 - v124
							if v125 <= int32(7) {
								v128 = v113 + v41
								v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
								v132 = int32(0)
								if v132 < v125 {
									v135 = v125
								} else {
									v135 = v132
								}
								v138 = v129 & (int32(255) << (uint(int32(8)-v135) % 32))
								*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v138)
							} else {
							}
							v141 = v110 - v124
							if v141 <= int32(7) {
								v144 = v113 + v72
								v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
								v148 = int32(0)
								if v148 < v141 {
									v151 = v141
								} else {
									v151 = v148
								}
								v154 = v145 & (int32(255) << (uint(int32(8)-v151) % 32))
								*(*uint8)(unsafe.Add(mBase, uint32(v144))) = uint8(v154)
							} else {
							}
							v158 = v113 + int32(1)
							if v158 != v102 {
								v113 = v158
								continue
							} else {
								break
							}
							break
						}
						v162 = v102
						v169 = float64(0)
						for {
							v170 = int32(1)
							v171 = v162 - v170
							v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v171))))
							v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171+v41))))
							v181 = base.F64_mul(base.F64_add(v169, base.F64_sub(base.F64_convert_i32_u(v173), base.F64_convert_i32_u(v176))), float64(0.00390625))
							if v170 < v162 {
								v162 = v171
								v169 = v181
								continue
							} else {
								break
							}
							break
						}
						F_pfree(m, v41)
						mBase = m.M
						v185 = m.ExcPending
						if v185 != 0 {
							return int64(0)
						} else {
							F_pfree(m, v72)
							mBase = m.M
							v187 = m.ExcPending
							if v187 != 0 {
								return int64(0)
							} else {
								v200 = base.I64_reinterpret_f64(v181)
								return v200
							}
						}
					}
				}
			} else {
				v200 = int64(4607182418800017408)
				return v200
			}
		}
	}
}
func F_brin_minmax_multi_distance_macaddr8(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+7)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return base.I64_reinterpret_f64(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_sub(base.F64_convert_i32_u(v4), base.F64_convert_i32_u(v7)), v10), base.F64_sub(base.F64_convert_i32_u(v12), base.F64_convert_i32_u(v14))), v10), base.F64_sub(base.F64_convert_i32_u(v20), base.F64_convert_i32_u(v22))), v10), base.F64_sub(base.F64_convert_i32_u(v28), base.F64_convert_i32_u(v30))), v10), base.F64_sub(base.F64_convert_i32_u(v36), base.F64_convert_i32_u(v38))), v10), base.F64_sub(base.F64_convert_i32_u(v44), base.F64_convert_i32_u(v46))), v10), base.F64_sub(base.F64_convert_i32_u(v52), base.F64_convert_i32_u(v54))), v10), base.F64_sub(base.F64_convert_i32_u(v60), base.F64_convert_i32_u(v62))), v10))
}
func F_brin_minmax_multi_serialize(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int64
	_ = v188
	var v190 int64
	_ = v190
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v208 int64
	_ = v208
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int64
	_ = v249
	var v251 int64
	_ = v251
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v276 int32
	_ = v276
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v314 int64
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int64
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int64
	_ = v373
	var v375 int32
	_ = v375
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v418 int32
	_ = v418
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = base.I32_wrap_i64(l1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	if v23<<(uint(int32(1))%32)+v26 <= v28 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v436 = F_brin_range_serialize(m, v22)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L6
	} else {
		goto L60
	}
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	if v30 == v26 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+8)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v35 = F_minmax_multi_get_strategy_procinfo(m, l0, v32, v33, int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
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
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+8)))
	v38 = F_minmax_multi_get_procinfo(m, l0, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_serialize[0]))
	v46 = F_AllocSetContextCreateInternal(m, v41, int32(_a_F_brin_minmax_multi_serialize_0), int32(0), int32(_a_F_brin_minmax_multi_serialize_1), int32(_a_F_brin_minmax_multi_serialize_2))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v48 = int32(_a_F_brin_minmax_multi_serialize_3)
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_serialize[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_serialize[0])) = v46
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v55 = F_build_expanded_ranges(m, v35, v52, v22, v20+int32(12))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v58 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v63 = v58 - int32(1)
	v64 = F_palloc0_mul(m, int32(16), v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	v127 = v57
	v129 = int32(0)
	goto L13
L13:
	;
	v140 = F_reduce_expanded_ranges(m, v55, v58, v129, v28, v35, v127)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L6
	} else {
		goto L23
	}
L14:
	;
	if int32(0) < v63 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v69 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	F_pg_qsort(m, v64, v63, int32(16), int32(21))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L22
	}
L18:
	;
	v88 = v64 + v69<<(uint(int32(4))%32)
	v91 = v55 + v69*int32(24)
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v91)+8))
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v91)+24))
	v94 = F_FunctionCall2Coll(m, v38, v57, v92, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L20
	}
L19:
	;
	goto L17
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v88)+8)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v69
	v99 = v69 + int32(1)
	if v99 != v63 {
		v69 = v99
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v127 = v122
	v129 = v64
	goto L13
L23:
	;
	v142 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v142
	if v140 <= v142 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v402
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_serialize[0])) = v49
	F_MemoryContextDelete(m, v46)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L6
	} else {
		goto L59
	}
L25:
	;
	v395 = int32(0)
	goto L27
L26:
	;
	v148 = v22 + int32(40)
	v150 = v140 - int32(1)
	if v150 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v395
	v402 = v395
	goto L24
L28:
	;
	v276 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v276
	if v150 == v276 {
		goto L45
	} else {
		goto L46
	}
L29:
	;
	v244 = v55 + v230*int32(24)
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+16)))
	if v245 != 0 {
		v259 = v225
		goto L28
	} else {
		goto L43
	}
L30:
	;
	v153 = int32(0)
	v225 = v153
	v230 = v153
	goto L29
L31:
	;
	goto L32
L32:
	;
	v159 = int32(0)
	v162 = v159
	v167 = v159
	v168 = v159
	goto L33
L33:
	;
	v181 = v55 + v167*int32(24)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+16)))
	if v182 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v140&int32(1) == int32(0) {
		v259 = v216
		goto L28
	} else {
		goto L42
	}
L35:
	;
	v187 = v148 + v162<<(uint(int32(3))%32)
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v181)))
	*(*int64)(unsafe.Add(mBase, uint32(v187))) = v188
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v181)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v187)+8)) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v192 + int32(1)
	v198 = v162 + int32(2)
	goto L37
L36:
	;
	v198 = v162
	goto L37
L37:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+40)))
	if v200 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v205 = v148 + v198<<(uint(int32(3))%32)
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v181)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v205))) = v206
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v181)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v205)+8)) = v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v210 + int32(1)
	v216 = v198 + int32(2)
	goto L40
L39:
	;
	v216 = v198
	goto L40
L40:
	;
	v218 = int32(2)
	v219 = v167 + v218
	v221 = v168 + v218
	if v221 != v140&int32(2147483646) {
		v162 = v216
		v167 = v219
		v168 = v221
		goto L33
	} else {
		goto L41
	}
L41:
	;
	goto L34
L42:
	;
	v225 = v216
	v230 = v219
	goto L29
L43:
	;
	v248 = v148 + v225<<(uint(int32(3))%32)
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v244)))
	*(*int64)(unsafe.Add(mBase, uint32(v248))) = v249
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v244)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v248)+8)) = v251
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v253 + int32(1)
	v259 = v225 + int32(2)
	goto L28
L44:
	;
	v366 = v55 + v351*int32(24)
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+16)))
	if v367 != int32(1) {
		v402 = v352
		goto L24
	} else {
		goto L58
	}
L45:
	;
	v347 = v259
	v351 = int32(0)
	v352 = v276
	goto L44
L46:
	;
	goto L47
L47:
	;
	v286 = int32(0)
	v288 = v259
	v292 = v286
	v293 = v276
	v294 = v286
	goto L48
L48:
	;
	v307 = v55 + v292*int32(24)
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+16)))
	if v308 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v140&int32(1) == int32(0) {
		v402 = v339
		goto L24
	} else {
		goto L57
	}
L50:
	;
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v307)))
	*(*int64)(unsafe.Add(mBase, uint32(v148+v288<<(uint(int32(3))%32)))) = v314
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v317 = int32(1)
	v318 = v316 + v317
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v318
	v322 = v288 + v317
	v323 = v318
	goto L52
L51:
	;
	v322 = v288
	v323 = v293
	goto L52
L52:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+40)))
	if v324 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v307)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v148+v322<<(uint(int32(3))%32)))) = v330
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v333 = int32(1)
	v334 = v332 + v333
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v334
	v338 = v322 + v333
	v339 = v334
	goto L55
L54:
	;
	v338 = v322
	v339 = v323
	goto L55
L55:
	;
	v340 = int32(2)
	v341 = v292 + v340
	v343 = v294 + v340
	if v343 != v140&int32(2147483646) {
		v288 = v338
		v292 = v341
		v293 = v339
		v294 = v343
		goto L48
	} else {
		goto L56
	}
L56:
	;
	goto L49
L57:
	;
	v347 = v338
	v351 = v341
	v352 = v339
	goto L44
L58:
	;
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v366)))
	*(*int64)(unsafe.Add(mBase, uint32(v148+v347<<(uint(int32(3))%32)))) = v373
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v395 = v375 + int32(1)
	goto L27
L59:
	;
	goto L1
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = base.I64_extend_i32_u(v436)
	m.G0 = v20 + int32(16)
	return
}
func F_brin_summarize_range(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 float64
	_ = v117
	var v120 int64
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
	var v144 int64
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
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = int64(0)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_brin_summarize_range[0])))
	if v19 == int32(1) {
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[1]))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+308))
		v27 = base.B2i32(v25 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_brin_summarize_range[0])) = uint8(v27)
		v29 = v27
	} else {
		v29 = v2
	}
	if v29 == int32(0) {
		if base.Ui64(int64(4294967296)) <= base.Ui64(v13) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v187 = m.ExcPending
			if v187 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v190 = m.ExcPending
				if v190 != 0 {
					return int64(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v11))) = v13
					F_errmsg(m, int32(_a_F_brin_summarize_range_0), v11)
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1408), int32(_a_F_brin_summarize_range_2))
						mBase = m.M
						v199 = m.ExcPending
						if v199 != 0 {
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
			v35 = F_IndexGetRelation(m, v14, int32(1))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int64(0)
			} else {
				if v35 != 0 {
					v40 = F_table_open(m, v35, int32(4))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int64(0)
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v11+int32(76)))) = v47
						v50 = *(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3]))
						*(*int32)(unsafe.Add(mBase, uint32(v11+int32(72)))) = v50
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v40)+48))
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+80))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
						*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v54 | int32(2)
						*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v53
						v62 = int32(_a_F_brin_summarize_range_3)
						v64 = *(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[4]))
						v66 = v64 + int32(1)
						*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[4])) = v66
						F_RestrictSearchPath(m)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int64(0)
						} else {
							v75 = v40
							v76 = v66
							v78 = F_index_open(m, v14, int32(4))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int64(0)
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+119)))
								if v81 != int32(105) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v206 = m.ExcPending
										if v206 != 0 {
											return int64(0)
										} else {
											v207 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v207 + int32(4)
											F_errmsg(m, int32(_a_F_brin_summarize_range_4), v11+int32(48))
											mBase = m.M
											v215 = m.ExcPending
											if v215 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1453), int32(_a_F_brin_summarize_range_2))
												mBase = m.M
												v220 = m.ExcPending
												if v220 != 0 {
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
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+84))
									if v84 != int32(3580) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v203 = m.ExcPending
										if v203 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(151027844))
											mBase = m.M
											v206 = m.ExcPending
											if v206 != 0 {
												return int64(0)
											} else {
												v207 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v207 + int32(4)
												F_errmsg(m, int32(_a_F_brin_summarize_range_4), v11+int32(48))
												mBase = m.M
												v215 = m.ExcPending
												if v215 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1453), int32(_a_F_brin_summarize_range_2))
													mBase = m.M
													v220 = m.ExcPending
													if v220 != 0 {
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
										if v75 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v224 = m.ExcPending
											if v224 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(16908420))
												mBase = m.M
												v227 = m.ExcPending
												if v227 != 0 {
													return int64(0)
												} else {
													v228 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v228 + int32(4)
													F_errmsg(m, int32(_a_F_brin_summarize_range_5), v11+int32(16))
													mBase = m.M
													v236 = m.ExcPending
													if v236 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1469), int32(_a_F_brin_summarize_range_2))
														mBase = m.M
														v241 = m.ExcPending
														if v241 != 0 {
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
											v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
											v91 = F_object_ownercheck(m, int32(1259), v14, v90)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int64(0)
											} else {
												if v91 == int32(0) {
													v97 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
													F_aclcheck_error(m, int32(2), int32(20), v97+int32(4))
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return int64(0)
													} else {
														v103 = F_IndexGetRelation(m, v14, int32(0))
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return int64(0)
														} else {
															if v103 != v35 {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v224 = m.ExcPending
																if v224 != 0 {
																	return int64(0)
																} else {
																	F_errcode(m, int32(16908420))
																	mBase = m.M
																	v227 = m.ExcPending
																	if v227 != 0 {
																		return int64(0)
																	} else {
																		v228 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
																		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v228 + int32(4)
																		F_errmsg(m, int32(_a_F_brin_summarize_range_5), v11+int32(16))
																		mBase = m.M
																		v236 = m.ExcPending
																		if v236 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1469), int32(_a_F_brin_summarize_range_2))
																			mBase = m.M
																			v241 = m.ExcPending
																			if v241 != 0 {
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
																v106 = *(*int32)(unsafe.Add(mBase, uint32(v78)+192))
																v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+18)))
																if v107 == int32(1) {
																	F_brinsummarize(m, v78, v75, base.I32_wrap_i64(v13), int32(1), v11-int32(-64), int32(0))
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return int64(0)
																	} else {
																		v117 = *(*float64)(unsafe.Add(mBase, uint32(v11)+64))
																		v144 = base.I64_extend_i32_s(base.I32_trunc_sat_f64_s(v117))
																		F_AtEOXact_GUC(m, int32(0), v76)
																		mBase = m.M
																		v147 = m.ExcPending
																		if v147 != 0 {
																			return int64(0)
																		} else {
																			v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																			v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																			*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																			*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																			F_relation_close(m, v78, int32(4))
																			mBase = m.M
																			v156 = m.ExcPending
																			if v156 != 0 {
																				return int64(0)
																			} else {
																				F_relation_close(m, v75, int32(4))
																				mBase = m.M
																				v159 = m.ExcPending
																				if v159 != 0 {
																					return int64(0)
																				} else {
																					m.G0 = v11 + int32(80)
																					return v144
																				}
																			}
																		}
																	}
																} else {
																	v120 = int64(0)
																	v123 = F_errstart(m, int32(14), int32(0))
																	mBase = m.M
																	v124 = m.ExcPending
																	if v124 != 0 {
																		return int64(0)
																	} else {
																		if v123 == int32(0) {
																			v144 = v120
																			F_AtEOXact_GUC(m, int32(0), v76)
																			mBase = m.M
																			v147 = m.ExcPending
																			if v147 != 0 {
																				return int64(0)
																			} else {
																				v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																				v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																				*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																				*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																				F_relation_close(m, v78, int32(4))
																				mBase = m.M
																				v156 = m.ExcPending
																				if v156 != 0 {
																					return int64(0)
																				} else {
																					F_relation_close(m, v75, int32(4))
																					mBase = m.M
																					v159 = m.ExcPending
																					if v159 != 0 {
																						return int64(0)
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
																				return int64(0)
																			} else {
																				v130 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v130 + int32(4)
																				F_errmsg(m, int32(_a_F_brin_summarize_range_6), v11+int32(32))
																				mBase = m.M
																				v138 = m.ExcPending
																				if v138 != 0 {
																					return int64(0)
																				} else {
																					F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1478), int32(_a_F_brin_summarize_range_2))
																					mBase = m.M
																					v143 = m.ExcPending
																					if v143 != 0 {
																						return int64(0)
																					} else {
																						v144 = v120
																						F_AtEOXact_GUC(m, int32(0), v76)
																						mBase = m.M
																						v147 = m.ExcPending
																						if v147 != 0 {
																							return int64(0)
																						} else {
																							v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																							v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																							*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																							*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																							F_relation_close(m, v78, int32(4))
																							mBase = m.M
																							v156 = m.ExcPending
																							if v156 != 0 {
																								return int64(0)
																							} else {
																								F_relation_close(m, v75, int32(4))
																								mBase = m.M
																								v159 = m.ExcPending
																								if v159 != 0 {
																									return int64(0)
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
													v103 = F_IndexGetRelation(m, v14, int32(0))
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return int64(0)
													} else {
														if v103 != v35 {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v224 = m.ExcPending
															if v224 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(16908420))
																mBase = m.M
																v227 = m.ExcPending
																if v227 != 0 {
																	return int64(0)
																} else {
																	v228 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v228 + int32(4)
																	F_errmsg(m, int32(_a_F_brin_summarize_range_5), v11+int32(16))
																	mBase = m.M
																	v236 = m.ExcPending
																	if v236 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1469), int32(_a_F_brin_summarize_range_2))
																		mBase = m.M
																		v241 = m.ExcPending
																		if v241 != 0 {
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
															v106 = *(*int32)(unsafe.Add(mBase, uint32(v78)+192))
															v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+18)))
															if v107 == int32(1) {
																F_brinsummarize(m, v78, v75, base.I32_wrap_i64(v13), int32(1), v11-int32(-64), int32(0))
																mBase = m.M
																v116 = m.ExcPending
																if v116 != 0 {
																	return int64(0)
																} else {
																	v117 = *(*float64)(unsafe.Add(mBase, uint32(v11)+64))
																	v144 = base.I64_extend_i32_s(base.I32_trunc_sat_f64_s(v117))
																	F_AtEOXact_GUC(m, int32(0), v76)
																	mBase = m.M
																	v147 = m.ExcPending
																	if v147 != 0 {
																		return int64(0)
																	} else {
																		v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																		v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																		*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																		*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																		F_relation_close(m, v78, int32(4))
																		mBase = m.M
																		v156 = m.ExcPending
																		if v156 != 0 {
																			return int64(0)
																		} else {
																			F_relation_close(m, v75, int32(4))
																			mBase = m.M
																			v159 = m.ExcPending
																			if v159 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v11 + int32(80)
																				return v144
																			}
																		}
																	}
																}
															} else {
																v120 = int64(0)
																v123 = F_errstart(m, int32(14), int32(0))
																mBase = m.M
																v124 = m.ExcPending
																if v124 != 0 {
																	return int64(0)
																} else {
																	if v123 == int32(0) {
																		v144 = v120
																		F_AtEOXact_GUC(m, int32(0), v76)
																		mBase = m.M
																		v147 = m.ExcPending
																		if v147 != 0 {
																			return int64(0)
																		} else {
																			v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																			v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																			*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																			*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																			F_relation_close(m, v78, int32(4))
																			mBase = m.M
																			v156 = m.ExcPending
																			if v156 != 0 {
																				return int64(0)
																			} else {
																				F_relation_close(m, v75, int32(4))
																				mBase = m.M
																				v159 = m.ExcPending
																				if v159 != 0 {
																					return int64(0)
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
																			return int64(0)
																		} else {
																			v130 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
																			*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v130 + int32(4)
																			F_errmsg(m, int32(_a_F_brin_summarize_range_6), v11+int32(32))
																			mBase = m.M
																			v138 = m.ExcPending
																			if v138 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1478), int32(_a_F_brin_summarize_range_2))
																				mBase = m.M
																				v143 = m.ExcPending
																				if v143 != 0 {
																					return int64(0)
																				} else {
																					v144 = v120
																					F_AtEOXact_GUC(m, int32(0), v76)
																					mBase = m.M
																					v147 = m.ExcPending
																					if v147 != 0 {
																						return int64(0)
																					} else {
																						v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																						v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																						*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																						*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																						F_relation_close(m, v78, int32(4))
																						mBase = m.M
																						v156 = m.ExcPending
																						if v156 != 0 {
																							return int64(0)
																						} else {
																							F_relation_close(m, v75, int32(4))
																							mBase = m.M
																							v159 = m.ExcPending
																							if v159 != 0 {
																								return int64(0)
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
					v70 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v70
					*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = int32(0)
					v75 = v2
					v76 = v70
					v78 = F_index_open(m, v14, int32(4))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int64(0)
					} else {
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
						v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+119)))
						if v81 != int32(105) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v203 = m.ExcPending
							if v203 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(151027844))
								mBase = m.M
								v206 = m.ExcPending
								if v206 != 0 {
									return int64(0)
								} else {
									v207 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v207 + int32(4)
									F_errmsg(m, int32(_a_F_brin_summarize_range_4), v11+int32(48))
									mBase = m.M
									v215 = m.ExcPending
									if v215 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1453), int32(_a_F_brin_summarize_range_2))
										mBase = m.M
										v220 = m.ExcPending
										if v220 != 0 {
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
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+84))
							if v84 != int32(3580) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v203 = m.ExcPending
								if v203 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(151027844))
									mBase = m.M
									v206 = m.ExcPending
									if v206 != 0 {
										return int64(0)
									} else {
										v207 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v207 + int32(4)
										F_errmsg(m, int32(_a_F_brin_summarize_range_4), v11+int32(48))
										mBase = m.M
										v215 = m.ExcPending
										if v215 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1453), int32(_a_F_brin_summarize_range_2))
											mBase = m.M
											v220 = m.ExcPending
											if v220 != 0 {
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
								if v75 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v224 = m.ExcPending
									if v224 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(16908420))
										mBase = m.M
										v227 = m.ExcPending
										if v227 != 0 {
											return int64(0)
										} else {
											v228 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v228 + int32(4)
											F_errmsg(m, int32(_a_F_brin_summarize_range_5), v11+int32(16))
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1469), int32(_a_F_brin_summarize_range_2))
												mBase = m.M
												v241 = m.ExcPending
												if v241 != 0 {
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
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
									v91 = F_object_ownercheck(m, int32(1259), v14, v90)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int64(0)
									} else {
										if v91 == int32(0) {
											v97 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
											F_aclcheck_error(m, int32(2), int32(20), v97+int32(4))
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return int64(0)
											} else {
												v103 = F_IndexGetRelation(m, v14, int32(0))
												mBase = m.M
												v104 = m.ExcPending
												if v104 != 0 {
													return int64(0)
												} else {
													if v103 != v35 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v224 = m.ExcPending
														if v224 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(16908420))
															mBase = m.M
															v227 = m.ExcPending
															if v227 != 0 {
																return int64(0)
															} else {
																v228 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
																*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v228 + int32(4)
																F_errmsg(m, int32(_a_F_brin_summarize_range_5), v11+int32(16))
																mBase = m.M
																v236 = m.ExcPending
																if v236 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1469), int32(_a_F_brin_summarize_range_2))
																	mBase = m.M
																	v241 = m.ExcPending
																	if v241 != 0 {
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
														v106 = *(*int32)(unsafe.Add(mBase, uint32(v78)+192))
														v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+18)))
														if v107 == int32(1) {
															F_brinsummarize(m, v78, v75, base.I32_wrap_i64(v13), int32(1), v11-int32(-64), int32(0))
															mBase = m.M
															v116 = m.ExcPending
															if v116 != 0 {
																return int64(0)
															} else {
																v117 = *(*float64)(unsafe.Add(mBase, uint32(v11)+64))
																v144 = base.I64_extend_i32_s(base.I32_trunc_sat_f64_s(v117))
																F_AtEOXact_GUC(m, int32(0), v76)
																mBase = m.M
																v147 = m.ExcPending
																if v147 != 0 {
																	return int64(0)
																} else {
																	v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																	v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																	*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																	*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																	F_relation_close(m, v78, int32(4))
																	mBase = m.M
																	v156 = m.ExcPending
																	if v156 != 0 {
																		return int64(0)
																	} else {
																		F_relation_close(m, v75, int32(4))
																		mBase = m.M
																		v159 = m.ExcPending
																		if v159 != 0 {
																			return int64(0)
																		} else {
																			m.G0 = v11 + int32(80)
																			return v144
																		}
																	}
																}
															}
														} else {
															v120 = int64(0)
															v123 = F_errstart(m, int32(14), int32(0))
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return int64(0)
															} else {
																if v123 == int32(0) {
																	v144 = v120
																	F_AtEOXact_GUC(m, int32(0), v76)
																	mBase = m.M
																	v147 = m.ExcPending
																	if v147 != 0 {
																		return int64(0)
																	} else {
																		v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																		v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																		*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																		*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																		F_relation_close(m, v78, int32(4))
																		mBase = m.M
																		v156 = m.ExcPending
																		if v156 != 0 {
																			return int64(0)
																		} else {
																			F_relation_close(m, v75, int32(4))
																			mBase = m.M
																			v159 = m.ExcPending
																			if v159 != 0 {
																				return int64(0)
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
																		return int64(0)
																	} else {
																		v130 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
																		*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v130 + int32(4)
																		F_errmsg(m, int32(_a_F_brin_summarize_range_6), v11+int32(32))
																		mBase = m.M
																		v138 = m.ExcPending
																		if v138 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1478), int32(_a_F_brin_summarize_range_2))
																			mBase = m.M
																			v143 = m.ExcPending
																			if v143 != 0 {
																				return int64(0)
																			} else {
																				v144 = v120
																				F_AtEOXact_GUC(m, int32(0), v76)
																				mBase = m.M
																				v147 = m.ExcPending
																				if v147 != 0 {
																					return int64(0)
																				} else {
																					v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																					v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																					*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																					*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																					F_relation_close(m, v78, int32(4))
																					mBase = m.M
																					v156 = m.ExcPending
																					if v156 != 0 {
																						return int64(0)
																					} else {
																						F_relation_close(m, v75, int32(4))
																						mBase = m.M
																						v159 = m.ExcPending
																						if v159 != 0 {
																							return int64(0)
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
											v103 = F_IndexGetRelation(m, v14, int32(0))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int64(0)
											} else {
												if v103 != v35 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v224 = m.ExcPending
													if v224 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(16908420))
														mBase = m.M
														v227 = m.ExcPending
														if v227 != 0 {
															return int64(0)
														} else {
															v228 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v228 + int32(4)
															F_errmsg(m, int32(_a_F_brin_summarize_range_5), v11+int32(16))
															mBase = m.M
															v236 = m.ExcPending
															if v236 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1469), int32(_a_F_brin_summarize_range_2))
																mBase = m.M
																v241 = m.ExcPending
																if v241 != 0 {
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
													v106 = *(*int32)(unsafe.Add(mBase, uint32(v78)+192))
													v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+18)))
													if v107 == int32(1) {
														F_brinsummarize(m, v78, v75, base.I32_wrap_i64(v13), int32(1), v11-int32(-64), int32(0))
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return int64(0)
														} else {
															v117 = *(*float64)(unsafe.Add(mBase, uint32(v11)+64))
															v144 = base.I64_extend_i32_s(base.I32_trunc_sat_f64_s(v117))
															F_AtEOXact_GUC(m, int32(0), v76)
															mBase = m.M
															v147 = m.ExcPending
															if v147 != 0 {
																return int64(0)
															} else {
																v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																F_relation_close(m, v78, int32(4))
																mBase = m.M
																v156 = m.ExcPending
																if v156 != 0 {
																	return int64(0)
																} else {
																	F_relation_close(m, v75, int32(4))
																	mBase = m.M
																	v159 = m.ExcPending
																	if v159 != 0 {
																		return int64(0)
																	} else {
																		m.G0 = v11 + int32(80)
																		return v144
																	}
																}
															}
														}
													} else {
														v120 = int64(0)
														v123 = F_errstart(m, int32(14), int32(0))
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return int64(0)
														} else {
															if v123 == int32(0) {
																v144 = v120
																F_AtEOXact_GUC(m, int32(0), v76)
																mBase = m.M
																v147 = m.ExcPending
																if v147 != 0 {
																	return int64(0)
																} else {
																	v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																	v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																	*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																	*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																	F_relation_close(m, v78, int32(4))
																	mBase = m.M
																	v156 = m.ExcPending
																	if v156 != 0 {
																		return int64(0)
																	} else {
																		F_relation_close(m, v75, int32(4))
																		mBase = m.M
																		v159 = m.ExcPending
																		if v159 != 0 {
																			return int64(0)
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
																	return int64(0)
																} else {
																	v130 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v130 + int32(4)
																	F_errmsg(m, int32(_a_F_brin_summarize_range_6), v11+int32(32))
																	mBase = m.M
																	v138 = m.ExcPending
																	if v138 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1478), int32(_a_F_brin_summarize_range_2))
																		mBase = m.M
																		v143 = m.ExcPending
																		if v143 != 0 {
																			return int64(0)
																		} else {
																			v144 = v120
																			F_AtEOXact_GUC(m, int32(0), v76)
																			mBase = m.M
																			v147 = m.ExcPending
																			if v147 != 0 {
																				return int64(0)
																			} else {
																				v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
																				v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
																				*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[3])) = v149
																				*(*int32)(unsafe.Add(mBase, _c_F_brin_summarize_range[2])) = v148
																				F_relation_close(m, v78, int32(4))
																				mBase = m.M
																				v156 = m.ExcPending
																				if v156 != 0 {
																					return int64(0)
																				} else {
																					F_relation_close(m, v75, int32(4))
																					mBase = m.M
																					v159 = m.ExcPending
																					if v159 != 0 {
																						return int64(0)
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
			return int64(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v170 = m.ExcPending
			if v170 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_brin_summarize_range_7), int32(0))
				mBase = m.M
				v174 = m.ExcPending
				if v174 != 0 {
					return int64(0)
				} else {
					F_errhint(m, int32(_a_F_brin_summarize_range_8), int32(0))
					mBase = m.M
					v178 = m.ExcPending
					if v178 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_brin_summarize_range_1), int32(1403), int32(_a_F_brin_summarize_range_2))
						mBase = m.M
						v183 = m.ExcPending
						if v183 != 0 {
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
	}
}
