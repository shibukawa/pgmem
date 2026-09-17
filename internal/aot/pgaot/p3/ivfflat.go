package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IvfflatAppendPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	v6 = int32(0)
	v12 = F_ReadBufferExtended(m, l0, l4, int32(-1), v6, v6)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		F_LockBuffer(m, v12, int32(2))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			v19 = F_GenericXLogRegisterBuffer(m, v17, v12, int32(1))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				if v12 < int32(0) {
					v24 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatAppendPage[0]))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v24+(v12^int32(-1))<<(uint(int32(6))%32))+16))
					v39 = v30
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatAppendPage[1]))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v32+v12<<(uint(int32(6))%32)+int32(-64))+16))
					v39 = v38
				}
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
				*(*int32)(unsafe.Add(mBase, uint32(v40+v41))) = v39
				v44 = int32(_a_F_IvfflatAppendPage_0)
				v46 = int32(0)
				if v46|(v19&int32(3)|int32(1)) == v46 {
					v62 = v19 + v44
					v64 = v19 + int32(4)
					if base.Ui32(v64) < base.Ui32(v62) {
						v66 = v62
					} else {
						v66 = v64
					}
					v71 = (v19^int32(-1)+v66)&int32(-4) + int32(4)
					if v71 == int32(0) {
					} else {
						base.MemoryFill(m, v19, int32(0), v71)
					}
				} else {
					base.MemoryFill(m, v19, int32(0), v44)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v19)+10)) = int32(_a_F_IvfflatAppendPage_1)
				v85 = int32(_a_F_IvfflatAppendPage_2)
				*(*uint16)(unsafe.Add(mBase, uint32(v19)+18)) = uint16(v85)
				v91 = int32(_a_F_IvfflatAppendPage_3)
				*(*uint16)(unsafe.Add(mBase, uint32(v19)+16)) = uint16(v91)
				*(*uint16)(unsafe.Add(mBase, uint32(v19)+14)) = uint16(v91)
				v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+16)))
				v95 = v19 + v94
				v96 = int32(_a_F_IvfflatAppendPage_4)
				*(*uint16)(unsafe.Add(mBase, uint32(v95)+6)) = uint16(v96)
				*(*int32)(unsafe.Add(mBase, uint32(v95))) = int32(-1)
				v100 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				F_GenericXLogFinish(m, v100)
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return
				} else {
					v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					F_UnlockReleaseBuffer(m, v103)
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return
					} else {
						v106 = F_GenericXLogStart(m, l0)
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v106
							v110 = F_GenericXLogRegisterBuffer(m, v106, v12, int32(1))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v110
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v12
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_IvfflatParallelScanAndSort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 float64
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 float64
	_ = v102
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v108 float64
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	v13 = m.G0
	v15 = v13 - int32(208)
	m.G0 = v15
	v18 = F_palloc0(m, int32(12))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(-1)
		v23 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v23)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v26 = F_BuildIndexInfo(m, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
			*(*uint8)(unsafe.Add(mBase, uint32(v26)+121)) = uint8(v28)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_InitBuildState_2(m, v15+int32(16), v32, v33, v26)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)+84))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
				v39 = v37 * v38
				if v39 != 0 {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
					base.MemoryCopy(m, v40, l3, v39)
				} else {
				}
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v36))) = v42
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v15)+172))
				v45 = int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v15)+206)) = uint16(v45)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+200)) = int32(97)
				v49 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+196)) = v49
				*(*uint8)(unsafe.Add(mBase, uint32(v15)+195)) = uint8(v49)
				v63 = F_tuplesort_begin_heap(m, v44, v45, v15+int32(206), v15+int32(200), v15+int32(196), v15+int32(195), l4, v18, v49)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v63
					*(*int32)(unsafe.Add(mBase, uint32(v15)+168)) = v63
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v70 = F_table_beginscan_parallel(m, v67, l1-int32(-64))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v75 = int32(0)
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v72)+188))
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+140))
						v83 = m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v72, v73, v26, int32(1), v75, l5, v75, int32(-1), int32(_a_F_IvfflatParallelScanAndSort_0), v15+int32(16), v70)
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return
						} else {
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							F_tuplesort_performsort(m, v85)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(1)
								if v88 != 0 {
									F_s_lock(m, l1+int32(28), int32(_a_F_IvfflatParallelScanAndSort_1), int32(692), int32(_a_F_IvfflatParallelScanAndSort_2))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v98 + int32(1)
										v102 = *(*float64)(unsafe.Add(mBase, uint32(l1)+40))
										*(*float64)(unsafe.Add(mBase, uint32(l1)+40)) = base.F64_add(v102, v83)
										v105 = *(*float64)(unsafe.Add(mBase, uint32(v15)+48))
										v106 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v106
										v108 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
										*(*float64)(unsafe.Add(mBase, uint32(l1)+48)) = base.F64_add(v105, v108)
										v113 = F_errstart(m, int32(14), v106)
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return
										} else {
											if v113 != 0 {
												*(*int64)(unsafe.Add(mBase, uint32(v15))) = base.I64_trunc_sat_f64_s(v83)
												if l5 != 0 {
													v119 = int32(_a_F_IvfflatParallelScanAndSort_3)
												} else {
													v119 = int32(_a_F_IvfflatParallelScanAndSort_4)
												}
												F_errmsg(m, v119, v15)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													if l5 != 0 {
														v125 = int32(703)
													} else {
														v125 = int32(705)
													}
													F_errfinish(m, int32(_a_F_IvfflatParallelScanAndSort_1), v125, int32(_a_F_IvfflatParallelScanAndSort_2))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return
													} else {
														F_ConditionVariableSignal(m, l1+int32(16))
														mBase = m.M
														v132 = m.ExcPending
														if v132 != 0 {
															return
														} else {
															v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															F_tuplesort_end(m, v133)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return
															} else {
																v136 = *(*int32)(unsafe.Add(mBase, uint32(v15)+84))
																F_VectorArrayFree(m, v136)
																mBase = m.M
																v138 = m.ExcPending
																if v138 != 0 {
																	return
																} else {
																	v139 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
																	F_pfree(m, v139)
																	mBase = m.M
																	v141 = m.ExcPending
																	if v141 != 0 {
																		return
																	} else {
																		v142 = *(*int32)(unsafe.Add(mBase, uint32(v15)+184))
																		F_MemoryContextDelete(m, v142)
																		mBase = m.M
																		v144 = m.ExcPending
																		if v144 != 0 {
																			return
																		} else {
																			m.G0 = v15 + int32(208)
																			return
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												F_ConditionVariableSignal(m, l1+int32(16))
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return
												} else {
													v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													F_tuplesort_end(m, v133)
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return
													} else {
														v136 = *(*int32)(unsafe.Add(mBase, uint32(v15)+84))
														F_VectorArrayFree(m, v136)
														mBase = m.M
														v138 = m.ExcPending
														if v138 != 0 {
															return
														} else {
															v139 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
															F_pfree(m, v139)
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return
															} else {
																v142 = *(*int32)(unsafe.Add(mBase, uint32(v15)+184))
																F_MemoryContextDelete(m, v142)
																mBase = m.M
																v144 = m.ExcPending
																if v144 != 0 {
																	return
																} else {
																	m.G0 = v15 + int32(208)
																	return
																}
															}
														}
													}
												}
											}
										}
									}
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v98 + int32(1)
									v102 = *(*float64)(unsafe.Add(mBase, uint32(l1)+40))
									*(*float64)(unsafe.Add(mBase, uint32(l1)+40)) = base.F64_add(v102, v83)
									v105 = *(*float64)(unsafe.Add(mBase, uint32(v15)+48))
									v106 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v106
									v108 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
									*(*float64)(unsafe.Add(mBase, uint32(l1)+48)) = base.F64_add(v105, v108)
									v113 = F_errstart(m, int32(14), v106)
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return
									} else {
										if v113 != 0 {
											*(*int64)(unsafe.Add(mBase, uint32(v15))) = base.I64_trunc_sat_f64_s(v83)
											if l5 != 0 {
												v119 = int32(_a_F_IvfflatParallelScanAndSort_3)
											} else {
												v119 = int32(_a_F_IvfflatParallelScanAndSort_4)
											}
											F_errmsg(m, v119, v15)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return
											} else {
												if l5 != 0 {
													v125 = int32(703)
												} else {
													v125 = int32(705)
												}
												F_errfinish(m, int32(_a_F_IvfflatParallelScanAndSort_1), v125, int32(_a_F_IvfflatParallelScanAndSort_2))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return
												} else {
													F_ConditionVariableSignal(m, l1+int32(16))
													mBase = m.M
													v132 = m.ExcPending
													if v132 != 0 {
														return
													} else {
														v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_tuplesort_end(m, v133)
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return
														} else {
															v136 = *(*int32)(unsafe.Add(mBase, uint32(v15)+84))
															F_VectorArrayFree(m, v136)
															mBase = m.M
															v138 = m.ExcPending
															if v138 != 0 {
																return
															} else {
																v139 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
																F_pfree(m, v139)
																mBase = m.M
																v141 = m.ExcPending
																if v141 != 0 {
																	return
																} else {
																	v142 = *(*int32)(unsafe.Add(mBase, uint32(v15)+184))
																	F_MemoryContextDelete(m, v142)
																	mBase = m.M
																	v144 = m.ExcPending
																	if v144 != 0 {
																		return
																	} else {
																		m.G0 = v15 + int32(208)
																		return
																	}
																}
															}
														}
													}
												}
											}
										} else {
											F_ConditionVariableSignal(m, l1+int32(16))
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return
											} else {
												v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_tuplesort_end(m, v133)
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return
												} else {
													v136 = *(*int32)(unsafe.Add(mBase, uint32(v15)+84))
													F_VectorArrayFree(m, v136)
													mBase = m.M
													v138 = m.ExcPending
													if v138 != 0 {
														return
													} else {
														v139 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
														F_pfree(m, v139)
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return
														} else {
															v142 = *(*int32)(unsafe.Add(mBase, uint32(v15)+184))
															F_MemoryContextDelete(m, v142)
															mBase = m.M
															v144 = m.ExcPending
															if v144 != 0 {
																return
															} else {
																m.G0 = v15 + int32(208)
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
