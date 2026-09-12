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
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v118 float64
	_ = v118
	var v120 float64
	_ = v120
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v170 int64
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int64
	_ = v228
	var v236 int32
	_ = v236
	var v239 int64
	_ = v239
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int64
	_ = v254
	var v262 int32
	_ = v262
	var v265 int64
	_ = v265
	var v273 int32
	_ = v273
	var v276 int64
	_ = v276
	var v284 int32
	_ = v284
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v309 int32
	_ = v309
	v12 = m.G0
	v14 = v12 - int32(128)
	m.G0 = v14
	if l2 == int32(0) {
		m.G0 = v14 + int32(128)
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v18 == int32(0) {
			m.G0 = v14 + int32(128)
			return
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(l2)+40))
			v22 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
			v23 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
			v24 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
			v25 = int32(515729)
			F_ExplainOpenGroup(m, v25, v25, int32(1), l0)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v32 = v21 + (v22 + (v23 + v24))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v33 == int32(0) {
					F_ExplainIndentText(m, l0)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						F_appendStringInfoString(m, v38, int32(733438))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v42 + int32(1)
							v48 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2))))
							F_ExplainPropertyInteger(m, int32(140179), int32(0), v48, l0)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								F_ExplainIndentText(m, l0)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									if l1&int32(16) != 0 {
										v58 = int32(340937)
									} else {
										v58 = int32(357768)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v14)+124)) = v58
									*(*int32)(unsafe.Add(mBase, uint32(v14)+120)) = int32(332138)
									if l1&int32(8) != 0 {
										v66 = int32(340937)
									} else {
										v66 = int32(357768)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v14)+116)) = v66
									*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = int32(145594)
									if l1&int32(2) != 0 {
										v74 = int32(340937)
									} else {
										v74 = int32(357768)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v74
									*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = int32(256440)
									if l1&int32(4) != 0 {
										v82 = int32(340937)
									} else {
										v82 = int32(357768)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v14)+100)) = v82
									*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = int32(331914)
									F_appendStringInfo(m, v53, int32(727127), v14+int32(96))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
										if v91 != int32(1) {
											v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v164 - int32(1)
											F_ExplainCloseGroup(m, int32(515729), int32(1), l0)
											mBase = m.M
											v309 = m.ExcPending
											if v309 != 0 {
												return
											} else {
												m.G0 = v14 + int32(128)
												return
											}
										} else {
											v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
											if v94 != int32(1) {
												v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v164 - int32(1)
												F_ExplainCloseGroup(m, int32(515729), int32(1), l0)
												mBase = m.M
												v309 = m.ExcPending
												if v309 != 0 {
													return
												} else {
													m.G0 = v14 + int32(128)
													return
												}
											} else {
												F_ExplainIndentText(m, l0)
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return
												} else {
													v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v100 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
													v101 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
													v102 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
													v103 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
													v104 = *(*int64)(unsafe.Add(mBase, uint32(l2)+40))
													*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = int32(307176)
													*(*int32)(unsafe.Add(mBase, uint32(v14-int32(-64)))) = int32(265799)
													*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(256440)
													*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(331914)
													*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(285200)
													v118 = float64(1e+09)
													v120 = float64(1000)
													*(*float64)(unsafe.Add(mBase, uint32(v14)+88)) = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v32), v118), v120)
													*(*float64)(unsafe.Add(mBase, uint32(v14)+72)) = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v104), v118), v120)
													*(*float64)(unsafe.Add(mBase, uint32(v14)+56)) = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v103), v118), v120)
													*(*float64)(unsafe.Add(mBase, uint32(v14)+40)) = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v102), v118), v120)
													*(*float64)(unsafe.Add(mBase, uint32(v14)+24)) = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v101), v118), v120)
													*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(258811)
													*(*float64)(unsafe.Add(mBase, uint32(v14)+8)) = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v100), v118), v120)
													F_appendStringInfo(m, v99, int32(726450), v14)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return
													} else {
														v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v164 - int32(1)
														F_ExplainCloseGroup(m, int32(515729), int32(1), l0)
														mBase = m.M
														v309 = m.ExcPending
														if v309 != 0 {
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
					v170 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2))))
					F_ExplainPropertyInteger(m, int32(140179), int32(0), v170, l0)
					mBase = m.M
					v172 = m.ExcPending
					if v172 != 0 {
						return
					} else {
						v173 = int32(137507)
						F_ExplainOpenGroup(m, v173, v173, int32(1), l0)
						mBase = m.M
						v177 = m.ExcPending
						if v177 != 0 {
							return
						} else {
							F_ExplainPropertyBool(m, int32(331914), int32(base.Ui32(l1&int32(4))>>(uint(int32(2))%32)), l0)
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return
							} else {
								F_ExplainPropertyBool(m, int32(256440), int32(base.Ui32(l1&int32(2))>>(uint(int32(1))%32)), l0)
								mBase = m.M
								v191 = m.ExcPending
								if v191 != 0 {
									return
								} else {
									F_ExplainPropertyBool(m, int32(145594), int32(base.Ui32(l1&int32(8))>>(uint(int32(3))%32)), l0)
									mBase = m.M
									v198 = m.ExcPending
									if v198 != 0 {
										return
									} else {
										F_ExplainPropertyBool(m, int32(332138), int32(base.Ui32(l1&int32(16))>>(uint(int32(4))%32)), l0)
										mBase = m.M
										v205 = m.ExcPending
										if v205 != 0 {
											return
										} else {
											F_ExplainCloseGroup(m, int32(137507), int32(1), l0)
											mBase = m.M
											v209 = m.ExcPending
											if v209 != 0 {
												return
											} else {
												v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
												if v210 != int32(1) {
													F_ExplainCloseGroup(m, int32(515729), int32(1), l0)
													mBase = m.M
													v309 = m.ExcPending
													if v309 != 0 {
														return
													} else {
														m.G0 = v14 + int32(128)
														return
													}
												} else {
													v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
													if v213 != int32(1) {
														F_ExplainCloseGroup(m, int32(515729), int32(1), l0)
														mBase = m.M
														v309 = m.ExcPending
														if v309 != 0 {
															return
														} else {
															m.G0 = v14 + int32(128)
															return
														}
													} else {
														v216 = int32(332208)
														F_ExplainOpenGroup(m, v216, v216, int32(1), l0)
														mBase = m.M
														v220 = m.ExcPending
														if v220 != 0 {
															return
														} else {
															v221 = int32(258811)
															F_ExplainOpenGroup(m, v221, v221, int32(1), l0)
															mBase = m.M
															v225 = m.ExcPending
															if v225 != 0 {
																return
															} else {
																v228 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
																F_ExplainPropertyFloat(m, int32(285200), int32(150400), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v228), float64(1e+09)), float64(1000)), int32(3), l0)
																mBase = m.M
																v236 = m.ExcPending
																if v236 != 0 {
																	return
																} else {
																	v239 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
																	F_ExplainPropertyFloat(m, int32(307176), int32(150400), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v239), float64(1e+09)), float64(1000)), int32(3), l0)
																	mBase = m.M
																	v247 = m.ExcPending
																	if v247 != 0 {
																		return
																	} else {
																		F_ExplainCloseGroup(m, int32(258811), int32(1), l0)
																		mBase = m.M
																		v251 = m.ExcPending
																		if v251 != 0 {
																			return
																		} else {
																			v254 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
																			F_ExplainPropertyFloat(m, int32(331914), int32(150400), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v254), float64(1e+09)), float64(1000)), int32(3), l0)
																			mBase = m.M
																			v262 = m.ExcPending
																			if v262 != 0 {
																				return
																			} else {
																				v265 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
																				F_ExplainPropertyFloat(m, int32(256440), int32(150400), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v265), float64(1e+09)), float64(1000)), int32(3), l0)
																				mBase = m.M
																				v273 = m.ExcPending
																				if v273 != 0 {
																					return
																				} else {
																					v276 = *(*int64)(unsafe.Add(mBase, uint32(l2)+40))
																					F_ExplainPropertyFloat(m, int32(265799), int32(150400), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v276), float64(1e+09)), float64(1000)), int32(3), l0)
																					mBase = m.M
																					v284 = m.ExcPending
																					if v284 != 0 {
																						return
																					} else {
																						F_ExplainPropertyFloat(m, int32(307176), int32(150400), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v32), float64(1e+09)), float64(1000)), int32(3), l0)
																						mBase = m.M
																						v294 = m.ExcPending
																						if v294 != 0 {
																							return
																						} else {
																							F_ExplainCloseGroup(m, int32(332208), int32(1), l0)
																							mBase = m.M
																							v298 = m.ExcPending
																							if v298 != 0 {
																								return
																							} else {
																								F_ExplainCloseGroup(m, int32(515729), int32(1), l0)
																								mBase = m.M
																								v309 = m.ExcPending
																								if v309 != 0 {
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
		}
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
	var v11 int32
	_ = v11
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
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
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v5 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v123 = F_CreateTemplateTupleDesc(m, int32(1))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L18
	} else {
		goto L43
	}
L2:
	;
	v119 = int32(25)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v9 = int32(0)
	v10 = int32(25)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v11 <= v9 {
		v119 = v10
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v14 = v9
	v15 = v10
	goto L6
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v14<<(uint(int32(2))%32))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v24 = int32(111255)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _consts[288])))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v28 == int32(0) {
		v47 = v27
		v48 = v28
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v119 = v112
	goto L1
L8:
	;
	v115 = v14 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v115 < v116 {
		v14 = v115
		v15 = v112
		goto L6
	} else {
		goto L42
	}
L9:
	;
	if v48-v47 != 0 {
		v112 = v15
		goto L8
	} else {
		goto L17
	}
L10:
	;
	goto L9
L11:
	;
	if v27 != v28 {
		v47 = v27
		v48 = v28
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v32 = v23
	v33 = v24
	goto L13
L13:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	if v37 == int32(0) {
		v47 = v36
		v48 = v37
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v47 = v36
	v48 = v37
	goto L10
L15:
	;
	v40 = int32(1)
	if v36 == v37 {
		v32 = v32 + v40
		v33 = v33 + v40
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v50 = F_defGetString(m, v22)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	v54 = int32(299074)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, _consts[289])))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v58 == int32(0) {
		v77 = v57
		v78 = v58
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v78-v77 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	goto L20
L22:
	;
	if v57 != v58 {
		v77 = v57
		v78 = v58
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v62 = v50
	v63 = v54
	goto L24
L24:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v67 == int32(0) {
		v77 = v66
		v78 = v67
		goto L21
	} else {
		goto L26
	}
L25:
	;
	v77 = v66
	v78 = v67
	goto L21
L26:
	;
	v70 = int32(1)
	if v66 == v67 {
		v62 = v62 + v70
		v63 = v63 + v70
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v112 = int32(142)
	goto L8
L29:
	;
	goto L30
L30:
	;
	v85 = int32(243182)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, _consts[290])))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v89 == int32(0) {
		v108 = v88
		v109 = v89
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v109-v108 != 0 {
		goto L39
	} else {
		goto L40
	}
L32:
	;
	goto L31
L33:
	;
	if v88 != v89 {
		v108 = v88
		v109 = v89
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v93 = v50
	v94 = v85
	goto L35
L35:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
	if v98 == int32(0) {
		v108 = v97
		v109 = v98
		goto L32
	} else {
		goto L37
	}
L36:
	;
	v108 = v97
	v109 = v98
	goto L32
L37:
	;
	v101 = int32(1)
	if v97 == v98 {
		v93 = v93 + v101
		v94 = v94 + v101
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v111 = int32(25)
	goto L41
L40:
	;
	v111 = int32(114)
	goto L41
L41:
	;
	v112 = v111
	goto L8
L42:
	;
	goto L7
L43:
	;
	F_TupleDescInitEntry(m, v123, int32(1), int32(525632), v119, int32(-1), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L18
	} else {
		goto L44
	}
L44:
	;
	return v123
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
