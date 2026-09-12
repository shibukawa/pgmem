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
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
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
					v24 = *(*int32)(unsafe.Add(mBase, _consts[6]))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v24+(v12^int32(-1))<<(uint(int32(6))%32))+16))
					v39 = v30
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, _consts[7]))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v32+v12<<(uint(int32(6))%32)+int32(-64))+16))
					v39 = v38
				}
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
				*(*int32)(unsafe.Add(mBase, uint32(v40+v41))) = v39
				if v19&int32(3) != 0 {
				} else {
				}
				v70 = F___memset(m, v19, int32(0), int32(8192))
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v19)+10)) = int32(1572864)
				v76 = int32(8196)
				*(*uint16)(unsafe.Add(mBase, uint32(v19)+18)) = uint16(v76)
				v82 = int32(8184)
				*(*uint16)(unsafe.Add(mBase, uint32(v19)+16)) = uint16(v82)
				*(*uint16)(unsafe.Add(mBase, uint32(v19)+14)) = uint16(v82)
				v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+16)))
				v86 = v19 + v85
				v87 = int32(65412)
				*(*uint16)(unsafe.Add(mBase, uint32(v86)+6)) = uint16(v87)
				*(*int32)(unsafe.Add(mBase, uint32(v86))) = int32(-1)
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				F_GenericXLogFinish(m, v91)
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return
				} else {
					v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					F_UnlockReleaseBuffer(m, v94)
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return
					} else {
						v97 = F_GenericXLogStart(m, l0)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v97
							v101 = F_GenericXLogRegisterBuffer(m, v97, v12, int32(1))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v101
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
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
	var v118 int64
	_ = v118
	var v120 int64
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	v12 = m.G0
	v14 = v12 - int32(208)
	m.G0 = v14
	v17 = F_palloc0(m, int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = int32(-1)
		v22 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v22)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v25 = F_BuildIndexInfo(m, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
			*(*uint8)(unsafe.Add(mBase, uint32(v25)+121)) = uint8(v27)
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_InitBuildState_2(m, v14+int32(16), v31, v32, v25)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
				v39 = v37 * v38
				if v39 != 0 {
					v40 = F__emscripten_memcpy_bulkmem(m, v36, l3, v39)
					mBase = m.M
				} else {
				}
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v35))) = v42
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v14)+172))
				v45 = int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v14)+206)) = uint16(v45)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+200)) = int32(97)
				v49 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+196)) = v49
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+195)) = uint8(v49)
				v63 = F_tuplesort_begin_heap(m, v44, v45, v14+int32(206), v14+int32(200), v14+int32(196), v14+int32(195), l4, v17, v49)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v63
					*(*int32)(unsafe.Add(mBase, uint32(v14)+168)) = v63
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
						v83 = m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v72, v73, v25, int32(1), v75, l5, v75, int32(-1), int32(7648), v14+int32(16), v70)
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
									F_s_lock(m, l1+int32(28), int32(522701), int32(692), int32(87249))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v98 + int32(1)
										v102 = *(*float64)(unsafe.Add(mBase, uint32(l1)+40))
										*(*float64)(unsafe.Add(mBase, uint32(l1)+40)) = base.F64_add(v102, v83)
										v105 = *(*float64)(unsafe.Add(mBase, uint32(v14)+48))
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
												if base.F64_lt(base.F64_abs(v83), float64(9.223372036854776e+18)) != 0 {
													v118 = base.I64_trunc_f64_s(v83)
													v120 = v118
												} else {
													v120 = int64(-9223372036854775807 - 1)
												}
												*(*int64)(unsafe.Add(mBase, uint32(v14))) = v120
												if l5 != 0 {
													v124 = int32(174324)
												} else {
													v124 = int32(174295)
												}
												F_errmsg(m, v124, v14)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return
												} else {
													if l5 != 0 {
														v130 = int32(703)
													} else {
														v130 = int32(705)
													}
													F_errfinish(m, int32(522701), v130, int32(87249))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return
													} else {
														F_ConditionVariableSignal(m, l1+int32(16))
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
															return
														} else {
															v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															F_tuplesort_end(m, v138)
															mBase = m.M
															v140 = m.ExcPending
															if v140 != 0 {
																return
															} else {
																v141 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
																F_VectorArrayFree(m, v141)
																mBase = m.M
																v143 = m.ExcPending
																if v143 != 0 {
																	return
																} else {
																	v144 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
																	F_pfree(m, v144)
																	mBase = m.M
																	v146 = m.ExcPending
																	if v146 != 0 {
																		return
																	} else {
																		v147 = *(*int32)(unsafe.Add(mBase, uint32(v14)+184))
																		F_MemoryContextDelete(m, v147)
																		mBase = m.M
																		v149 = m.ExcPending
																		if v149 != 0 {
																			return
																		} else {
																			m.G0 = v14 + int32(208)
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
												v137 = m.ExcPending
												if v137 != 0 {
													return
												} else {
													v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													F_tuplesort_end(m, v138)
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return
													} else {
														v141 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
														F_VectorArrayFree(m, v141)
														mBase = m.M
														v143 = m.ExcPending
														if v143 != 0 {
															return
														} else {
															v144 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
															F_pfree(m, v144)
															mBase = m.M
															v146 = m.ExcPending
															if v146 != 0 {
																return
															} else {
																v147 = *(*int32)(unsafe.Add(mBase, uint32(v14)+184))
																F_MemoryContextDelete(m, v147)
																mBase = m.M
																v149 = m.ExcPending
																if v149 != 0 {
																	return
																} else {
																	m.G0 = v14 + int32(208)
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
									v105 = *(*float64)(unsafe.Add(mBase, uint32(v14)+48))
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
											if base.F64_lt(base.F64_abs(v83), float64(9.223372036854776e+18)) != 0 {
												v118 = base.I64_trunc_f64_s(v83)
												v120 = v118
											} else {
												v120 = int64(-9223372036854775807 - 1)
											}
											*(*int64)(unsafe.Add(mBase, uint32(v14))) = v120
											if l5 != 0 {
												v124 = int32(174324)
											} else {
												v124 = int32(174295)
											}
											F_errmsg(m, v124, v14)
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
												return
											} else {
												if l5 != 0 {
													v130 = int32(703)
												} else {
													v130 = int32(705)
												}
												F_errfinish(m, int32(522701), v130, int32(87249))
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return
												} else {
													F_ConditionVariableSignal(m, l1+int32(16))
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return
													} else {
														v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_tuplesort_end(m, v138)
														mBase = m.M
														v140 = m.ExcPending
														if v140 != 0 {
															return
														} else {
															v141 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
															F_VectorArrayFree(m, v141)
															mBase = m.M
															v143 = m.ExcPending
															if v143 != 0 {
																return
															} else {
																v144 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
																F_pfree(m, v144)
																mBase = m.M
																v146 = m.ExcPending
																if v146 != 0 {
																	return
																} else {
																	v147 = *(*int32)(unsafe.Add(mBase, uint32(v14)+184))
																	F_MemoryContextDelete(m, v147)
																	mBase = m.M
																	v149 = m.ExcPending
																	if v149 != 0 {
																		return
																	} else {
																		m.G0 = v14 + int32(208)
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
											v137 = m.ExcPending
											if v137 != 0 {
												return
											} else {
												v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_tuplesort_end(m, v138)
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return
												} else {
													v141 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
													F_VectorArrayFree(m, v141)
													mBase = m.M
													v143 = m.ExcPending
													if v143 != 0 {
														return
													} else {
														v144 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
														F_pfree(m, v144)
														mBase = m.M
														v146 = m.ExcPending
														if v146 != 0 {
															return
														} else {
															v147 = *(*int32)(unsafe.Add(mBase, uint32(v14)+184))
															F_MemoryContextDelete(m, v147)
															mBase = m.M
															v149 = m.ExcPending
															if v149 != 0 {
																return
															} else {
																m.G0 = v14 + int32(208)
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
