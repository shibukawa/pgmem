package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_multirange_adjacent_multirange(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v21 = F_pg_detoast_datum(m, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
			if v23 == int32(0) {
				v105 = v2
				m.G0 = v13 + int32(80)
				return v105
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
				if v26 == int32(0) {
					v105 = v2
					m.G0 = v13 + int32(80)
					return v105
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
					if v31 != 0 {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
						if v32 == v29 {
							v44 = v31
							v45 = v23
							v46 = v26
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
							v51 = v13 + int32(72)
							v53 = v13 - int32(-64)
							F_multirange_get_bounds(m, v47, v16, v45-int32(1), v51, v53)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
								F_multirange_get_bounds(m, v56, v21, int32(0), v13+int32(56), v13+int32(48))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
									v65 = *(*int64)(unsafe.Add(mBase, uint32(v13)+64))
									*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v65
									v67 = *(*int64)(unsafe.Add(mBase, uint32(v13)+56))
									*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v67
									v74 = F_bounds_adjacent(m, v64, v13+int32(40), v13+int32(32))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										if v74 != 0 {
											v105 = int32(1)
											m.G0 = v13 + int32(80)
											return v105
										} else {
											if int32(2) <= v45 {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
												F_multirange_get_bounds(m, v78, v16, int32(0), v51, v53)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													if int32(2) <= v46 {
														v84 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
														F_multirange_get_bounds(m, v84, v21, v46-int32(1), v13+int32(56), v13+int32(48))
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return int32(0)
														} else {
															v93 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
															v94 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
															*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v94
															v96 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
															*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v96
															v102 = F_bounds_adjacent(m, v93, v13+int32(24), v13+int32(16))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return int32(0)
															} else {
																v105 = v102
																m.G0 = v13 + int32(80)
																return v105
															}
														}
													} else {
														v93 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
														v94 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
														*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v94
														v96 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
														*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v96
														v102 = F_bounds_adjacent(m, v93, v13+int32(24), v13+int32(16))
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return int32(0)
														} else {
															v105 = v102
															m.G0 = v13 + int32(80)
															return v105
														}
													}
												}
											} else {
												if int32(2) <= v46 {
													v84 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
													F_multirange_get_bounds(m, v84, v21, v46-int32(1), v13+int32(56), v13+int32(48))
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return int32(0)
													} else {
														v93 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
														v94 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
														*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v94
														v96 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
														*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v96
														v102 = F_bounds_adjacent(m, v93, v13+int32(24), v13+int32(16))
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return int32(0)
														} else {
															v105 = v102
															m.G0 = v13 + int32(80)
															return v105
														}
													}
												} else {
													v93 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
													v94 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
													*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v94
													v96 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
													*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v96
													v102 = F_bounds_adjacent(m, v93, v13+int32(24), v13+int32(16))
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														v105 = v102
														m.G0 = v13 + int32(80)
														return v105
													}
												}
											}
										}
									}
								}
							}
						} else {
							v35 = F_lookup_type_cache(m, v29, int32(_a_F_multirange_adjacent_multirange_0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+296))
								if v37 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13))) = v29
										F_errmsg_internal(m, int32(_a_F_multirange_adjacent_multirange_1), v13)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_multirange_adjacent_multirange_2), int32(558), int32(_a_F_multirange_adjacent_multirange_3))
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v35
									v42 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
									v44 = v35
									v45 = v43
									v46 = v42
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
									v51 = v13 + int32(72)
									v53 = v13 - int32(-64)
									F_multirange_get_bounds(m, v47, v16, v45-int32(1), v51, v53)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										v56 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
										F_multirange_get_bounds(m, v56, v21, int32(0), v13+int32(56), v13+int32(48))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
											v65 = *(*int64)(unsafe.Add(mBase, uint32(v13)+64))
											*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v65
											v67 = *(*int64)(unsafe.Add(mBase, uint32(v13)+56))
											*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v67
											v74 = F_bounds_adjacent(m, v64, v13+int32(40), v13+int32(32))
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return int32(0)
											} else {
												if v74 != 0 {
													v105 = int32(1)
													m.G0 = v13 + int32(80)
													return v105
												} else {
													if int32(2) <= v45 {
														v78 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
														F_multirange_get_bounds(m, v78, v16, int32(0), v51, v53)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															if int32(2) <= v46 {
																v84 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
																F_multirange_get_bounds(m, v84, v21, v46-int32(1), v13+int32(56), v13+int32(48))
																mBase = m.M
																v92 = m.ExcPending
																if v92 != 0 {
																	return int32(0)
																} else {
																	v93 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
																	v94 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
																	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v94
																	v96 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
																	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v96
																	v102 = F_bounds_adjacent(m, v93, v13+int32(24), v13+int32(16))
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return int32(0)
																	} else {
																		v105 = v102
																		m.G0 = v13 + int32(80)
																		return v105
																	}
																}
															} else {
																v93 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
																v94 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
																*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v94
																v96 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
																*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v96
																v102 = F_bounds_adjacent(m, v93, v13+int32(24), v13+int32(16))
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return int32(0)
																} else {
																	v105 = v102
																	m.G0 = v13 + int32(80)
																	return v105
																}
															}
														}
													} else {
														if int32(2) <= v46 {
															v84 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
															F_multirange_get_bounds(m, v84, v21, v46-int32(1), v13+int32(56), v13+int32(48))
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return int32(0)
															} else {
																v93 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
																v94 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
																*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v94
																v96 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
																*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v96
																v102 = F_bounds_adjacent(m, v93, v13+int32(24), v13+int32(16))
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return int32(0)
																} else {
																	v105 = v102
																	m.G0 = v13 + int32(80)
																	return v105
																}
															}
														} else {
															v93 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
															v94 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
															*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v94
															v96 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
															*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v96
															v102 = F_bounds_adjacent(m, v93, v13+int32(24), v13+int32(16))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return int32(0)
															} else {
																v105 = v102
																m.G0 = v13 + int32(80)
																return v105
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
						v35 = F_lookup_type_cache(m, v29, int32(_a_F_multirange_adjacent_multirange_0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+296))
							if v37 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13))) = v29
									F_errmsg_internal(m, int32(_a_F_multirange_adjacent_multirange_1), v13)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_multirange_adjacent_multirange_2), int32(558), int32(_a_F_multirange_adjacent_multirange_3))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v35
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
								v44 = v35
								v45 = v43
								v46 = v42
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
								v51 = v13 + int32(72)
								v53 = v13 - int32(-64)
								F_multirange_get_bounds(m, v47, v16, v45-int32(1), v51, v53)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
									F_multirange_get_bounds(m, v56, v21, int32(0), v13+int32(56), v13+int32(48))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
										v65 = *(*int64)(unsafe.Add(mBase, uint32(v13)+64))
										*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v65
										v67 = *(*int64)(unsafe.Add(mBase, uint32(v13)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v67
										v74 = F_bounds_adjacent(m, v64, v13+int32(40), v13+int32(32))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											if v74 != 0 {
												v105 = int32(1)
												m.G0 = v13 + int32(80)
												return v105
											} else {
												if int32(2) <= v45 {
													v78 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
													F_multirange_get_bounds(m, v78, v16, int32(0), v51, v53)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														if int32(2) <= v46 {
															v84 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
															F_multirange_get_bounds(m, v84, v21, v46-int32(1), v13+int32(56), v13+int32(48))
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return int32(0)
															} else {
																v93 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
																v94 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
																*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v94
																v96 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
																*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v96
																v102 = F_bounds_adjacent(m, v93, v13+int32(24), v13+int32(16))
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return int32(0)
																} else {
																	v105 = v102
																	m.G0 = v13 + int32(80)
																	return v105
																}
															}
														} else {
															v93 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
															v94 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
															*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v94
															v96 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
															*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v96
															v102 = F_bounds_adjacent(m, v93, v13+int32(24), v13+int32(16))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return int32(0)
															} else {
																v105 = v102
																m.G0 = v13 + int32(80)
																return v105
															}
														}
													}
												} else {
													if int32(2) <= v46 {
														v84 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
														F_multirange_get_bounds(m, v84, v21, v46-int32(1), v13+int32(56), v13+int32(48))
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return int32(0)
														} else {
															v93 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
															v94 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
															*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v94
															v96 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
															*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v96
															v102 = F_bounds_adjacent(m, v93, v13+int32(24), v13+int32(16))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return int32(0)
															} else {
																v105 = v102
																m.G0 = v13 + int32(80)
																return v105
															}
														}
													} else {
														v93 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
														v94 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
														*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v94
														v96 = *(*int64)(unsafe.Add(mBase, uint32(v13)+72))
														*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v96
														v102 = F_bounds_adjacent(m, v93, v13+int32(24), v13+int32(16))
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return int32(0)
														} else {
															v105 = v102
															m.G0 = v13 + int32(80)
															return v105
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
func F_multirange_contained_by_range(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					v34 = F_range_contains_multirange_internal(m, v33, v17, v12)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v34
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_contained_by_range_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_multirange_contained_by_range_1), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_contained_by_range_2), int32(558), int32(_a_F_multirange_contained_by_range_3))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = F_range_contains_multirange_internal(m, v33, v17, v12)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_contained_by_range_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_multirange_contained_by_range_1), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_contained_by_range_2), int32(558), int32(_a_F_multirange_contained_by_range_3))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = F_range_contains_multirange_internal(m, v33, v17, v12)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v34
						}
					}
				}
			}
		}
	}
}
func F_multirange_contains_range(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					v34 = F_multirange_contains_range_internal(m, v33, v12, v17)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v34
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_contains_range_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_multirange_contains_range_1), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_contains_range_2), int32(558), int32(_a_F_multirange_contains_range_3))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = F_multirange_contains_range_internal(m, v33, v12, v17)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_contains_range_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_multirange_contains_range_1), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_contains_range_2), int32(558), int32(_a_F_multirange_contains_range_3))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = F_multirange_contains_range_internal(m, v33, v12, v17)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v34
						}
					}
				}
			}
		}
	}
}
func F_multirange_contains_range_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v98 int32
	_ = v98
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v20 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v14)>>(uint(int32(2))%32))-v13))))
	goto L2
L1:
	;
	m.G0 = v11 + int32(48)
	return v98
L2:
	;
	if v20&int32(1) != 0 {
		v98 = v13
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v23 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v98 = int32(0)
	goto L1
L5:
	;
	v27 = v11 + int32(16)
	v29 = v27 | int32(8)
	F_range_deserialize(m, l0, l2, v27, v29, v11+int32(15))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v36 = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v37 == v36 {
		v98 = v36
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v42 = v37
	v43 = v36
	goto L9
L9:
	;
	v50 = int32(base.Ui32(v42+v43) >> (uint(int32(1)) % 32))
	v52 = v11 + int32(40)
	F_multirange_get_bounds(m, l0, l1, v50, v52, v11+int32(32))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L11
	}
L10:
	;
	goto L4
L11:
	;
	v57 = F_range_cmp_bounds(m, l0, v29, v52)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L13
	}
L12:
	;
	if base.Ui32(v83) < base.Ui32(v82) {
		v42 = v82
		v43 = v83
		goto L9
	} else {
		goto L25
	}
L13:
	;
	if v57 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v82 = v50
	v83 = v43
	goto L12
L15:
	;
	goto L16
L16:
	;
	v62 = v11 + int32(16)
	v64 = v11 + int32(32)
	v65 = F_range_cmp_bounds(m, l0, v62, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	if v65 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v71 = F_range_cmp_bounds(m, l0, v11+int32(40), v62)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v82 = v42
	v83 = v50 + int32(1)
	goto L12
L21:
	;
	if int32(0) < v71 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v76 = F_range_cmp_bounds(m, l0, v64, v29)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	if v76 < int32(0) {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v98 = int32(1)
	goto L1
L25:
	;
	goto L10
}
func F_multirange_intersect_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	v7 = int32(0)
	if l4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = l2
	goto L3
L2:
	;
	v14 = v7
	goto L3
L3:
	;
	if v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = int32(0)
	v19 = F_make_multirange(m, l0, l1, v17, v17)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v27 = F_palloc0(m, (l2+l4)<<(uint(int32(2))%32))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	return v19
L9:
	;
	if l2 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v32 = F_make_multirange(m, l0, l1, int32(0), v27)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v42 = v35
	v43 = v7
	v46 = v7
	v47 = v7
	goto L14
L13:
	;
	return v32
L14:
	;
	if v42 == int32(0) {
		v124 = v46
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v126 = F_make_multirange(m, l0, l1, v124, v27)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L37
	}
L16:
	;
	goto L15
L17:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l3+v47<<(uint(int32(2))%32))))
	v60 = v42
	v61 = v43
	goto L18
L18:
	;
	v66 = F_range_before_internal(m, l1, v60, v53)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L20
	}
L19:
	;
	v81 = v60
	v82 = v61
	v85 = v46
	goto L26
L20:
	;
	if v66 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v69 = v61 + int32(1)
	if l4 <= v69 {
		v124 = v46
		goto L16
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L19
L24:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l5+v69<<(uint(int32(2))%32))))
	if v74 != 0 {
		v60 = v74
		v61 = v69
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v124 = v46
	goto L16
L26:
	;
	v87 = F_range_overlaps_internal(m, l1, v53, v81)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L29
	}
L27:
	;
	v112 = v47 + int32(1)
	if v112 != l2 {
		v42 = v81
		v43 = v82
		v46 = v110
		v47 = v112
		goto L14
	} else {
		goto L36
	}
L28:
	;
	goto L27
L29:
	;
	if v87 == int32(0) {
		v110 = v85
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v94 = F_range_intersect_internal(m, l1, v53, v81)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27+v85<<(uint(int32(2))%32)))) = v94
	v98 = v85 + int32(1)
	v99 = F_range_overleft_internal(m, l1, v81, v53)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	if v99 == int32(0) {
		v110 = v98
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v104 = v82 + int32(1)
	if l4 <= v104 {
		v124 = v98
		goto L16
	} else {
		goto L34
	}
L34:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l5+v104<<(uint(int32(2))%32))))
	if v109 != 0 {
		v81 = v109
		v82 = v104
		v85 = v98
		goto L26
	} else {
		goto L35
	}
L35:
	;
	v124 = v98
	goto L16
L36:
	;
	v124 = v110
	goto L16
L37:
	;
	return v126
}
func F_multirange_overlaps_multirange(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
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
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L43
	}
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+296))
	v39 = int32(0)
	v40 = m.G0
	v42 = v40 - int32(32)
	m.G0 = v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v44 == v39 {
		v153 = v39
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v27 == v24 {
		v37 = v26
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v30 = F_lookup_type_cache(m, v24, int32(_a_F_multirange_overlaps_multirange_0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+296))
	if v32 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v30
	v37 = v30
	goto L5
L12:
	;
	m.G0 = v42 + int32(32)
	m.G0 = v14 + int32(16)
	return v153
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v47 == int32(0) {
		v153 = v39
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_multirange_get_bounds(m, v38, v17, int32(0), v42+int32(24), v42+int32(16))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v47 <= int32(0) {
		v153 = v39
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v61 = v39
	v68 = v2
	v69 = v2
	goto L17
L17:
	;
	v71 = v42 + int32(8)
	F_multirange_get_bounds(m, v38, v22, v68, v71, v42)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	v153 = v147
	goto L12
L19:
	;
	v76 = F_range_cmp_bounds(m, v38, v42+int32(16), v71)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v76 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v90 = v69
	goto L24
L22:
	;
	v116 = v69
	goto L23
L23:
	;
	v118 = v42 + int32(24)
	v121 = F_range_cmp_bounds(m, v38, v118, v42+int32(8))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L30
	}
L24:
	;
	v92 = v90 + int32(1)
	if v44 <= v92 {
		v153 = v61
		goto L12
	} else {
		goto L26
	}
L25:
	;
	v116 = v92
	goto L23
L26:
	;
	v97 = v42 + int32(16)
	F_multirange_get_bounds(m, v38, v17, v92, v42+int32(24), v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v102 = F_range_cmp_bounds(m, v38, v97, v42+int32(8))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v102 < int32(0) {
		v90 = v92
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	if int32(0) <= v121 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v126 = F_range_cmp_bounds(m, v38, v118, v42)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v132 = v42 + int32(8)
	v135 = F_range_cmp_bounds(m, v38, v132, v42+int32(24))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	if v126 <= int32(0) {
		v153 = int32(1)
		goto L12
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	if int32(0) <= v135 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v142 = F_range_cmp_bounds(m, v38, v132, v42+int32(16))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v147 = int32(0)
	v149 = v68 + int32(1)
	if v149 != v47 {
		v61 = v147
		v68 = v149
		v69 = v116
		goto L17
	} else {
		goto L42
	}
L40:
	;
	if v142 <= int32(0) {
		v153 = int32(1)
		goto L12
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L18
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v24
	F_errmsg_internal(m, int32(_a_F_multirange_overlaps_multirange_1), v14)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_multirange_overlaps_multirange_2), int32(558), int32(_a_F_multirange_overlaps_multirange_3))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
