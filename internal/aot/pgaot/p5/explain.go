package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExplainPrintJIT(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v114 float64
	_ = v114
	var v116 float64
	_ = v116
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v165 int64
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int64
	_ = v223
	var v231 int32
	_ = v231
	var v234 int64
	_ = v234
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v249 int64
	_ = v249
	var v257 int32
	_ = v257
	var v260 int64
	_ = v260
	var v268 int32
	_ = v268
	var v271 int64
	_ = v271
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v302 int32
	_ = v302
	v12 = m.G0
	v14 = v12 - int32(128)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v16 != 0 {
		v17 = *(*int64)(unsafe.Add(mBase, uint32(l2)+40))
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
		v20 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
		v21 = int32(_a_F_ExplainPrintJIT_0)
		F_ExplainOpenGroup(m, v21, v21, int32(1), l0)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v28 = v17 + (v18 + (v19 + v20))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v29 == int32(0) {
				F_ExplainIndentText(m, l0)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					F_appendStringInfoString(m, v34, int32(_a_F_ExplainPrintJIT_1))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v38 + int32(1)
						v44 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2))))
						F_ExplainPropertyInteger(m, int32(_a_F_ExplainPrintJIT_2), int32(0), v44, l0)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							F_ExplainIndentText(m, l0)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if l1&int32(16) != 0 {
									v54 = int32(_a_F_ExplainPrintJIT_3)
								} else {
									v54 = int32(_a_F_ExplainPrintJIT_4)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v14)+124)) = v54
								*(*int32)(unsafe.Add(mBase, uint32(v14)+120)) = int32(_a_F_ExplainPrintJIT_5)
								if l1&int32(8) != 0 {
									v62 = int32(_a_F_ExplainPrintJIT_3)
								} else {
									v62 = int32(_a_F_ExplainPrintJIT_4)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v14)+116)) = v62
								*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = int32(_a_F_ExplainPrintJIT_6)
								if l1&int32(2) != 0 {
									v70 = int32(_a_F_ExplainPrintJIT_3)
								} else {
									v70 = int32(_a_F_ExplainPrintJIT_4)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v70
								*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = int32(_a_F_ExplainPrintJIT_7)
								if l1&int32(4) != 0 {
									v78 = int32(_a_F_ExplainPrintJIT_3)
								} else {
									v78 = int32(_a_F_ExplainPrintJIT_4)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v14)+100)) = v78
								*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = int32(_a_F_ExplainPrintJIT_8)
								F_appendStringInfo(m, v49, int32(_a_F_ExplainPrintJIT_9), v14+int32(96))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
									if v87 != int32(1) {
										v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v159 - int32(1)
										F_ExplainCloseGroup(m, int32(_a_F_ExplainPrintJIT_0), int32(1), l0)
										mBase = m.M
										v302 = m.ExcPending
										if v302 != 0 {
											return
										} else {
											m.G0 = v14 + int32(128)
											return
										}
									} else {
										v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
										if v90 != int32(1) {
											v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v159 - int32(1)
											F_ExplainCloseGroup(m, int32(_a_F_ExplainPrintJIT_0), int32(1), l0)
											mBase = m.M
											v302 = m.ExcPending
											if v302 != 0 {
												return
											} else {
												m.G0 = v14 + int32(128)
												return
											}
										} else {
											F_ExplainIndentText(m, l0)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return
											} else {
												v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v96 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
												v97 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
												v98 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
												v99 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
												v100 = *(*int64)(unsafe.Add(mBase, uint32(l2)+40))
												*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = int32(_a_F_ExplainPrintJIT_10)
												*(*int32)(unsafe.Add(mBase, uint32(v14-int32(-64)))) = int32(_a_F_ExplainPrintJIT_11)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(_a_F_ExplainPrintJIT_7)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(_a_F_ExplainPrintJIT_8)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_ExplainPrintJIT_12)
												v114 = float64(1e+09)
												v116 = float64(1000)
												*(*float64)(unsafe.Add(mBase, uint32(v14)+88)) = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v28), v114), v116)
												*(*float64)(unsafe.Add(mBase, uint32(v14)+72)) = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v100), v114), v116)
												*(*float64)(unsafe.Add(mBase, uint32(v14)+56)) = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v99), v114), v116)
												*(*float64)(unsafe.Add(mBase, uint32(v14)+40)) = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v98), v114), v116)
												*(*float64)(unsafe.Add(mBase, uint32(v14)+24)) = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v97), v114), v116)
												*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(_a_F_ExplainPrintJIT_13)
												*(*float64)(unsafe.Add(mBase, uint32(v14)+8)) = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v96), v114), v116)
												F_appendStringInfo(m, v95, int32(_a_F_ExplainPrintJIT_14), v14)
												mBase = m.M
												v153 = m.ExcPending
												if v153 != 0 {
													return
												} else {
													v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v159 - int32(1)
													F_ExplainCloseGroup(m, int32(_a_F_ExplainPrintJIT_0), int32(1), l0)
													mBase = m.M
													v302 = m.ExcPending
													if v302 != 0 {
														return
													} else {
														m.G0 = v14 + int32(128)
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
			} else {
				v165 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2))))
				F_ExplainPropertyInteger(m, int32(_a_F_ExplainPrintJIT_2), int32(0), v165, l0)
				mBase = m.M
				v167 = m.ExcPending
				if v167 != 0 {
					return
				} else {
					v168 = int32(_a_F_ExplainPrintJIT_15)
					F_ExplainOpenGroup(m, v168, v168, int32(1), l0)
					mBase = m.M
					v172 = m.ExcPending
					if v172 != 0 {
						return
					} else {
						F_ExplainPropertyBool(m, int32(_a_F_ExplainPrintJIT_8), int32(base.Ui32(l1&int32(4))>>(uint(int32(2))%32)), l0)
						mBase = m.M
						v179 = m.ExcPending
						if v179 != 0 {
							return
						} else {
							F_ExplainPropertyBool(m, int32(_a_F_ExplainPrintJIT_7), int32(base.Ui32(l1&int32(2))>>(uint(int32(1))%32)), l0)
							mBase = m.M
							v186 = m.ExcPending
							if v186 != 0 {
								return
							} else {
								F_ExplainPropertyBool(m, int32(_a_F_ExplainPrintJIT_6), int32(base.Ui32(l1&int32(8))>>(uint(int32(3))%32)), l0)
								mBase = m.M
								v193 = m.ExcPending
								if v193 != 0 {
									return
								} else {
									F_ExplainPropertyBool(m, int32(_a_F_ExplainPrintJIT_5), int32(base.Ui32(l1&int32(16))>>(uint(int32(4))%32)), l0)
									mBase = m.M
									v200 = m.ExcPending
									if v200 != 0 {
										return
									} else {
										F_ExplainCloseGroup(m, int32(_a_F_ExplainPrintJIT_15), int32(1), l0)
										mBase = m.M
										v204 = m.ExcPending
										if v204 != 0 {
											return
										} else {
											v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
											if v205 != int32(1) {
												F_ExplainCloseGroup(m, int32(_a_F_ExplainPrintJIT_0), int32(1), l0)
												mBase = m.M
												v302 = m.ExcPending
												if v302 != 0 {
													return
												} else {
													m.G0 = v14 + int32(128)
													return
												}
											} else {
												v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
												if v208 != int32(1) {
													F_ExplainCloseGroup(m, int32(_a_F_ExplainPrintJIT_0), int32(1), l0)
													mBase = m.M
													v302 = m.ExcPending
													if v302 != 0 {
														return
													} else {
														m.G0 = v14 + int32(128)
														return
													}
												} else {
													v211 = int32(_a_F_ExplainPrintJIT_16)
													F_ExplainOpenGroup(m, v211, v211, int32(1), l0)
													mBase = m.M
													v215 = m.ExcPending
													if v215 != 0 {
														return
													} else {
														v216 = int32(_a_F_ExplainPrintJIT_13)
														F_ExplainOpenGroup(m, v216, v216, int32(1), l0)
														mBase = m.M
														v220 = m.ExcPending
														if v220 != 0 {
															return
														} else {
															v223 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
															F_ExplainPropertyFloat(m, int32(_a_F_ExplainPrintJIT_12), int32(_a_F_ExplainPrintJIT_17), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v223), float64(1e+09)), float64(1000)), int32(3), l0)
															mBase = m.M
															v231 = m.ExcPending
															if v231 != 0 {
																return
															} else {
																v234 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
																F_ExplainPropertyFloat(m, int32(_a_F_ExplainPrintJIT_10), int32(_a_F_ExplainPrintJIT_17), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v234), float64(1e+09)), float64(1000)), int32(3), l0)
																mBase = m.M
																v242 = m.ExcPending
																if v242 != 0 {
																	return
																} else {
																	F_ExplainCloseGroup(m, int32(_a_F_ExplainPrintJIT_13), int32(1), l0)
																	mBase = m.M
																	v246 = m.ExcPending
																	if v246 != 0 {
																		return
																	} else {
																		v249 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
																		F_ExplainPropertyFloat(m, int32(_a_F_ExplainPrintJIT_8), int32(_a_F_ExplainPrintJIT_17), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v249), float64(1e+09)), float64(1000)), int32(3), l0)
																		mBase = m.M
																		v257 = m.ExcPending
																		if v257 != 0 {
																			return
																		} else {
																			v260 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
																			F_ExplainPropertyFloat(m, int32(_a_F_ExplainPrintJIT_7), int32(_a_F_ExplainPrintJIT_17), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v260), float64(1e+09)), float64(1000)), int32(3), l0)
																			mBase = m.M
																			v268 = m.ExcPending
																			if v268 != 0 {
																				return
																			} else {
																				v271 = *(*int64)(unsafe.Add(mBase, uint32(l2)+40))
																				F_ExplainPropertyFloat(m, int32(_a_F_ExplainPrintJIT_11), int32(_a_F_ExplainPrintJIT_17), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v271), float64(1e+09)), float64(1000)), int32(3), l0)
																				mBase = m.M
																				v279 = m.ExcPending
																				if v279 != 0 {
																					return
																				} else {
																					F_ExplainPropertyFloat(m, int32(_a_F_ExplainPrintJIT_10), int32(_a_F_ExplainPrintJIT_17), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v28), float64(1e+09)), float64(1000)), int32(3), l0)
																					mBase = m.M
																					v289 = m.ExcPending
																					if v289 != 0 {
																						return
																					} else {
																						F_ExplainCloseGroup(m, int32(_a_F_ExplainPrintJIT_16), int32(1), l0)
																						mBase = m.M
																						v293 = m.ExcPending
																						if v293 != 0 {
																							return
																						} else {
																							F_ExplainCloseGroup(m, int32(_a_F_ExplainPrintJIT_0), int32(1), l0)
																							mBase = m.M
																							v302 = m.ExcPending
																							if v302 != 0 {
																								return
																							} else {
																								m.G0 = v14 + int32(128)
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
		m.G0 = v14 + int32(128)
		return
	}
}
func F_ExplainResultDesc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v5 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v125 = F_CreateTemplateTupleDesc(m, int32(1))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L17
	} else {
		goto L40
	}
L2:
	;
	v121 = int32(25)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v9 = int32(25)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v10 <= int32(0) {
		v121 = v9
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v14 = int32(0)
	v15 = v9
	goto L6
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v14<<(uint(int32(2))%32))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v24 = int32(_a_F_ExplainResultDesc_0)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExplainResultDesc[0])))
	if base.B2i32(v27 == int32(0))|base.B2i32(v27 != v30) != 0 {
		v48 = v27
		v49 = v30
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v121 = v115
	goto L1
L8:
	;
	v117 = v14 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v117 < v118 {
		v14 = v117
		v15 = v115
		goto L6
	} else {
		goto L39
	}
L9:
	;
	if v48-v49 != 0 {
		v115 = v15
		goto L8
	} else {
		goto L16
	}
L10:
	;
	goto L9
L11:
	;
	v33 = v23
	v34 = v24
	goto L12
L12:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	if v38 == int32(0) {
		v48 = v38
		v49 = v37
		goto L10
	} else {
		goto L14
	}
L13:
	;
	v48 = v38
	v49 = v37
	goto L10
L14:
	;
	v41 = int32(1)
	if v38 == v37 {
		v33 = v33 + v41
		v34 = v34 + v41
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v51 = F_defGetString(m, v22)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	v55 = int32(_a_F_ExplainResultDesc_1)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExplainResultDesc[1])))
	if base.B2i32(v58 == int32(0))|base.B2i32(v58 != v61) != 0 {
		v79 = v58
		v80 = v61
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v79-v80 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	goto L19
L21:
	;
	v64 = v51
	v65 = v55
	goto L22
L22:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	if v69 == int32(0) {
		v79 = v69
		v80 = v68
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v79 = v69
	v80 = v68
	goto L20
L24:
	;
	v72 = int32(1)
	if v69 == v68 {
		v64 = v64 + v72
		v65 = v65 + v72
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v115 = int32(142)
	goto L8
L27:
	;
	goto L28
L28:
	;
	v87 = int32(_a_F_ExplainResultDesc_2)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExplainResultDesc[2])))
	if base.B2i32(v90 == int32(0))|base.B2i32(v90 != v93) != 0 {
		v111 = v90
		v112 = v93
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v111-v112 != 0 {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	v96 = v51
	v97 = v87
	goto L32
L32:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
	if v101 == int32(0) {
		v111 = v101
		v112 = v100
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v111 = v101
	v112 = v100
	goto L30
L34:
	;
	v104 = int32(1)
	if v101 == v100 {
		v96 = v96 + v104
		v97 = v97 + v104
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v114 = int32(25)
	goto L38
L37:
	;
	v114 = int32(114)
	goto L38
L38:
	;
	v115 = v114
	goto L8
L39:
	;
	goto L7
L40:
	;
	F_TupleDescInitEntry(m, v125, int32(1), int32(_a_F_ExplainResultDesc_3), v121, int32(-1), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L17
	} else {
		goto L41
	}
L41:
	;
	return v125
}
func F_ExplainSaveGroup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	switch v3 - int32(1) {
	case 0:
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v6 - int32(2)
		return
	case 1:
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v10 - int32(2)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = F_list_delete_first(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v19
			return
		}
	case 2:
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v22 - int32(2)
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v28
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v31 = F_list_delete_first(m, v30)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v31
			return
		}
	default:
		return
	}
}
func F_explain_ExecutorStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v41 float64
	_ = v41
	var v43 float64
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[0]))
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[1]))
	if v7 != 0 {
		v51 = v5
	} else {
		if v5 < int32(0) {
			v49 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[2])) = uint8(v49)
			v51 = v5
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[3]))
			if int32(0) <= v11 {
				v49 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[2])) = uint8(v49)
				v51 = v5
			} else {
				v15 = int32(_a_F_explain_ExecutorStart_0)
				v18 = *(*int64)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[4]))
				v19 = *(*int64)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[5]))
				v20 = v18 ^ v19
				*(*int64)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[5])) = base.I64_rotl(v20, int64(37))
				*(*int64)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[4])) = v20<<(uint(int64(16))%64) ^ base.I64_rotl(v18, int64(24)) ^ v20
				v41 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v18*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
				mBase = m.M
				v43 = *(*float64)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[6]))
				v44 = base.F64_lt(v41, v43)
				*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[2])) = uint8(v44)
				v47 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[0]))
				v51 = v47
			}
		}
	}
	if l1&int32(1)|base.B2i32(v51 < int32(0)) != 0 {
	} else {
		v58 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[7])))
		v60 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[1]))
		v61 = int32(0)
		if (v58|base.B2i32(v60 == v61))&int32(1) == v61 {
		} else {
			v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[2])))
			if v69&int32(1) == int32(0) {
			} else {
				v75 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[8])))
				if v75&int32(1) == int32(0) {
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v84 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[9])))
					if v84 != 0 {
						v85 = int32(1)
					} else {
						v85 = int32(4)
					}
					v86 = v80 | v85
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v86
					v89 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[10])))
					if v89 == int32(1) {
						v93 = v86 | int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v93
						v95 = v93
					} else {
						v95 = v86
					}
					v97 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[11])))
					if v97 != int32(1) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v95 | int32(8)
					}
				}
			}
		}
	}
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[12]))
	if v105 != 0 {
		m.T0[v105].(func(*base.Module, int32, int32))(m, l0, l1)
		mBase = m.M
		v107 = m.ExcPending
		if v107 != 0 {
			return
		} else {
			v111 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[0]))
			if v111 < int32(0) {
				return
			} else {
				v115 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[7])))
				v117 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[1]))
				v118 = int32(0)
				if (v115|base.B2i32(v117 == v118))&int32(1) == v118 {
					return
				} else {
					v126 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[2])))
					if v126&int32(1) == int32(0) {
						return
					} else {
						v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						if v131 != 0 {
							return
						} else {
							v132 = int32(_a_F_explain_ExecutorStart_1)
							v133 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[13]))
							v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+100))
							*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[13])) = v136
							v141 = F_InstrAlloc(m, int32(1), int32(2147483647), int32(0))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v141
								*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[13])) = v133
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_standard_ExecutorStart(m, l0, l1)
		mBase = m.M
		v109 = m.ExcPending
		if v109 != 0 {
			return
		} else {
			v111 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[0]))
			if v111 < int32(0) {
				return
			} else {
				v115 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[7])))
				v117 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[1]))
				v118 = int32(0)
				if (v115|base.B2i32(v117 == v118))&int32(1) == v118 {
					return
				} else {
					v126 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[2])))
					if v126&int32(1) == int32(0) {
						return
					} else {
						v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						if v131 != 0 {
							return
						} else {
							v132 = int32(_a_F_explain_ExecutorStart_1)
							v133 = *(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[13]))
							v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+100))
							*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[13])) = v136
							v141 = F_InstrAlloc(m, int32(1), int32(2147483647), int32(0))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v141
								*(*int32)(unsafe.Add(mBase, _c_F_explain_ExecutorStart[13])) = v133
								return
							}
						}
					}
				}
			}
		}
	}
}
