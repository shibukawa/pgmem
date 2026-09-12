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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = F_pg_detoast_datum(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
			if v21 == int32(0) {
				v108 = v2
				m.G0 = v11 + int32(80)
				return v108
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
				if v24 == int32(0) {
					v108 = v2
					m.G0 = v11 + int32(80)
					return v108
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
					if v29 != 0 {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
						if v30 == v27 {
							v42 = v29
							v43 = v21
							v44 = v24
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
							F_multirange_get_bounds(m, v45, v14, v43-int32(1), v11+int32(72), v11-int32(-64))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
								F_multirange_get_bounds(m, v54, v19, int32(0), v11+int32(56), v11+int32(48))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
									v63 = *(*int64)(unsafe.Add(mBase, uint32(v11)+64))
									*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v63
									v65 = *(*int64)(unsafe.Add(mBase, uint32(v11)+56))
									*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
									v72 = F_bounds_adjacent(m, v62, v11+int32(40), v11+int32(32))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										if v72 != 0 {
											v108 = int32(1)
											m.G0 = v11 + int32(80)
											return v108
										} else {
											if int32(2) <= v43 {
												v76 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
												F_multirange_get_bounds(m, v76, v14, int32(0), v11+int32(72), v11-int32(-64))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													if int32(2) <= v44 {
														v86 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
														F_multirange_get_bounds(m, v86, v19, v44-int32(1), v11+int32(56), v11+int32(48))
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int32(0)
														} else {
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
															v96 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
															*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v96
															v98 = *(*int64)(unsafe.Add(mBase, uint32(v11)+72))
															*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v98
															v104 = F_bounds_adjacent(m, v95, v11+int32(24), v11+int32(16))
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return int32(0)
															} else {
																v108 = v104
																m.G0 = v11 + int32(80)
																return v108
															}
														}
													} else {
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
														v96 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
														*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v96
														v98 = *(*int64)(unsafe.Add(mBase, uint32(v11)+72))
														*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v98
														v104 = F_bounds_adjacent(m, v95, v11+int32(24), v11+int32(16))
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return int32(0)
														} else {
															v108 = v104
															m.G0 = v11 + int32(80)
															return v108
														}
													}
												}
											} else {
												if int32(2) <= v44 {
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
													F_multirange_get_bounds(m, v86, v19, v44-int32(1), v11+int32(56), v11+int32(48))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return int32(0)
													} else {
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
														v96 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
														*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v96
														v98 = *(*int64)(unsafe.Add(mBase, uint32(v11)+72))
														*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v98
														v104 = F_bounds_adjacent(m, v95, v11+int32(24), v11+int32(16))
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return int32(0)
														} else {
															v108 = v104
															m.G0 = v11 + int32(80)
															return v108
														}
													}
												} else {
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
													v96 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
													*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v96
													v98 = *(*int64)(unsafe.Add(mBase, uint32(v11)+72))
													*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v98
													v104 = F_bounds_adjacent(m, v95, v11+int32(24), v11+int32(16))
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return int32(0)
													} else {
														v108 = v104
														m.G0 = v11 + int32(80)
														return v108
													}
												}
											}
										}
									}
								}
							}
						} else {
							v33 = F_lookup_type_cache(m, v27, int32(65536))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
								if v35 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v27
										F_errmsg_internal(m, int32(389826), v11)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(519712), int32(558), int32(419268))
											mBase = m.M
											v127 = m.ExcPending
											if v127 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v33
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
									v42 = v33
									v43 = v41
									v44 = v40
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
									F_multirange_get_bounds(m, v45, v14, v43-int32(1), v11+int32(72), v11-int32(-64))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
										F_multirange_get_bounds(m, v54, v19, int32(0), v11+int32(56), v11+int32(48))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
											v63 = *(*int64)(unsafe.Add(mBase, uint32(v11)+64))
											*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v63
											v65 = *(*int64)(unsafe.Add(mBase, uint32(v11)+56))
											*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
											v72 = F_bounds_adjacent(m, v62, v11+int32(40), v11+int32(32))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												if v72 != 0 {
													v108 = int32(1)
													m.G0 = v11 + int32(80)
													return v108
												} else {
													if int32(2) <= v43 {
														v76 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
														F_multirange_get_bounds(m, v76, v14, int32(0), v11+int32(72), v11-int32(-64))
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
															return int32(0)
														} else {
															if int32(2) <= v44 {
																v86 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
																F_multirange_get_bounds(m, v86, v19, v44-int32(1), v11+int32(56), v11+int32(48))
																mBase = m.M
																v94 = m.ExcPending
																if v94 != 0 {
																	return int32(0)
																} else {
																	v95 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
																	v96 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
																	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v96
																	v98 = *(*int64)(unsafe.Add(mBase, uint32(v11)+72))
																	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v98
																	v104 = F_bounds_adjacent(m, v95, v11+int32(24), v11+int32(16))
																	mBase = m.M
																	v105 = m.ExcPending
																	if v105 != 0 {
																		return int32(0)
																	} else {
																		v108 = v104
																		m.G0 = v11 + int32(80)
																		return v108
																	}
																}
															} else {
																v95 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
																v96 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
																*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v96
																v98 = *(*int64)(unsafe.Add(mBase, uint32(v11)+72))
																*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v98
																v104 = F_bounds_adjacent(m, v95, v11+int32(24), v11+int32(16))
																mBase = m.M
																v105 = m.ExcPending
																if v105 != 0 {
																	return int32(0)
																} else {
																	v108 = v104
																	m.G0 = v11 + int32(80)
																	return v108
																}
															}
														}
													} else {
														if int32(2) <= v44 {
															v86 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
															F_multirange_get_bounds(m, v86, v19, v44-int32(1), v11+int32(56), v11+int32(48))
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return int32(0)
															} else {
																v95 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
																v96 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
																*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v96
																v98 = *(*int64)(unsafe.Add(mBase, uint32(v11)+72))
																*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v98
																v104 = F_bounds_adjacent(m, v95, v11+int32(24), v11+int32(16))
																mBase = m.M
																v105 = m.ExcPending
																if v105 != 0 {
																	return int32(0)
																} else {
																	v108 = v104
																	m.G0 = v11 + int32(80)
																	return v108
																}
															}
														} else {
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
															v96 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
															*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v96
															v98 = *(*int64)(unsafe.Add(mBase, uint32(v11)+72))
															*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v98
															v104 = F_bounds_adjacent(m, v95, v11+int32(24), v11+int32(16))
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return int32(0)
															} else {
																v108 = v104
																m.G0 = v11 + int32(80)
																return v108
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
						v33 = F_lookup_type_cache(m, v27, int32(65536))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
							if v35 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v27
									F_errmsg_internal(m, int32(389826), v11)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(519712), int32(558), int32(419268))
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v33
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
								v42 = v33
								v43 = v41
								v44 = v40
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
								F_multirange_get_bounds(m, v45, v14, v43-int32(1), v11+int32(72), v11-int32(-64))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
									F_multirange_get_bounds(m, v54, v19, int32(0), v11+int32(56), v11+int32(48))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
										v63 = *(*int64)(unsafe.Add(mBase, uint32(v11)+64))
										*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v63
										v65 = *(*int64)(unsafe.Add(mBase, uint32(v11)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
										v72 = F_bounds_adjacent(m, v62, v11+int32(40), v11+int32(32))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											if v72 != 0 {
												v108 = int32(1)
												m.G0 = v11 + int32(80)
												return v108
											} else {
												if int32(2) <= v43 {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
													F_multirange_get_bounds(m, v76, v14, int32(0), v11+int32(72), v11-int32(-64))
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int32(0)
													} else {
														if int32(2) <= v44 {
															v86 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
															F_multirange_get_bounds(m, v86, v19, v44-int32(1), v11+int32(56), v11+int32(48))
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return int32(0)
															} else {
																v95 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
																v96 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
																*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v96
																v98 = *(*int64)(unsafe.Add(mBase, uint32(v11)+72))
																*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v98
																v104 = F_bounds_adjacent(m, v95, v11+int32(24), v11+int32(16))
																mBase = m.M
																v105 = m.ExcPending
																if v105 != 0 {
																	return int32(0)
																} else {
																	v108 = v104
																	m.G0 = v11 + int32(80)
																	return v108
																}
															}
														} else {
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
															v96 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
															*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v96
															v98 = *(*int64)(unsafe.Add(mBase, uint32(v11)+72))
															*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v98
															v104 = F_bounds_adjacent(m, v95, v11+int32(24), v11+int32(16))
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return int32(0)
															} else {
																v108 = v104
																m.G0 = v11 + int32(80)
																return v108
															}
														}
													}
												} else {
													if int32(2) <= v44 {
														v86 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
														F_multirange_get_bounds(m, v86, v19, v44-int32(1), v11+int32(56), v11+int32(48))
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int32(0)
														} else {
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
															v96 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
															*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v96
															v98 = *(*int64)(unsafe.Add(mBase, uint32(v11)+72))
															*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v98
															v104 = F_bounds_adjacent(m, v95, v11+int32(24), v11+int32(16))
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return int32(0)
															} else {
																v108 = v104
																m.G0 = v11 + int32(80)
																return v108
															}
														}
													} else {
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
														v96 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
														*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v96
														v98 = *(*int64)(unsafe.Add(mBase, uint32(v11)+72))
														*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v98
														v104 = F_bounds_adjacent(m, v95, v11+int32(24), v11+int32(16))
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return int32(0)
														} else {
															v108 = v104
															m.G0 = v11 + int32(80)
															return v108
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
					v25 = F_lookup_type_cache(m, v19, int32(65536))
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
								F_errmsg_internal(m, int32(389826), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(519712), int32(558), int32(419268))
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
				v25 = F_lookup_type_cache(m, v19, int32(65536))
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
							F_errmsg_internal(m, int32(389826), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(519712), int32(558), int32(419268))
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
					v25 = F_lookup_type_cache(m, v19, int32(65536))
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
								F_errmsg_internal(m, int32(389826), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(519712), int32(558), int32(419268))
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
				v25 = F_lookup_type_cache(m, v19, int32(65536))
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
							F_errmsg_internal(m, int32(389826), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(519712), int32(558), int32(419268))
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v19 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v13)>>(uint(int32(2))%32))-v12))))
	goto L2
L1:
	;
	m.G0 = v10 + int32(48)
	return v103
L2:
	;
	if v19&int32(1) != 0 {
		v103 = v12
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v22 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v103 = int32(0)
	goto L1
L5:
	;
	v26 = v10 + int32(16)
	v30 = v26 | int32(8)
	F_range_deserialize(m, l0, l2, v26, v30, v10+int32(15))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v37 = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v38 == v37 {
		v103 = v37
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v43 = v38
	v45 = v37
	goto L9
L9:
	;
	v50 = int32(base.Ui32(v43+v45) >> (uint(int32(1)) % 32))
	F_multirange_get_bounds(m, l0, l1, v50, v10+int32(40), v10+int32(32))
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
	v59 = F_range_cmp_bounds(m, l0, v30, v10+int32(40))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L13
	}
L12:
	;
	if base.Ui32(v89) < base.Ui32(v88) {
		v43 = v88
		v45 = v89
		goto L9
	} else {
		goto L25
	}
L13:
	;
	if v59 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v88 = v50
	v89 = v45
	goto L12
L15:
	;
	goto L16
L16:
	;
	v67 = F_range_cmp_bounds(m, l0, v10+int32(16), v10+int32(32))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	if v67 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v75 = F_range_cmp_bounds(m, l0, v10+int32(40), v10+int32(16))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v88 = v43
	v89 = v50 + int32(1)
	goto L12
L21:
	;
	if int32(0) < v75 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v82 = F_range_cmp_bounds(m, l0, v10+int32(32), v30)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	if v82 < int32(0) {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v103 = int32(1)
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
	var v44 int32
	_ = v44
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
	var v83 int32
	_ = v83
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
	var v122 int32
	_ = v122
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
	v44 = v7
	v47 = v7
	goto L14
L13:
	;
	return v32
L14:
	;
	if v42 == int32(0) {
		v122 = v44
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v126 = F_make_multirange(m, l0, l1, v122, v27)
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
	v83 = v44
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
		v122 = v44
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
	v122 = v44
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
		v44 = v110
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
		v110 = v83
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
	*(*int32)(unsafe.Add(mBase, uint32(v27+v83<<(uint(int32(2))%32)))) = v94
	v98 = v83 + int32(1)
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
		v122 = v98
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
		v83 = v98
		goto L26
	} else {
		goto L35
	}
L35:
	;
	v122 = v98
	goto L16
L36:
	;
	v122 = v110
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L43
	}
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+296))
	v38 = int32(0)
	v39 = m.G0
	v41 = v39 - int32(32)
	m.G0 = v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v43 == v38 {
		v158 = v38
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v26 == v23 {
		v36 = v25
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v29 = F_lookup_type_cache(m, v23, int32(65536))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+296))
	if v31 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v29
	v36 = v29
	goto L5
L12:
	;
	m.G0 = v41 + int32(32)
	m.G0 = v13 + int32(16)
	return v158
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v46 == int32(0) {
		v158 = v38
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_multirange_get_bounds(m, v37, v16, int32(0), v41+int32(24), v41+int32(16))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v46 <= int32(0) {
		v158 = v38
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v61 = v38
	v64 = v2
	v65 = v2
	goto L17
L17:
	;
	F_multirange_get_bounds(m, v37, v21, v64, v41+int32(8), v41)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	v158 = v151
	goto L12
L19:
	;
	v76 = F_range_cmp_bounds(m, v37, v41+int32(16), v41+int32(8))
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
	v87 = v65
	goto L24
L22:
	;
	v114 = v65
	goto L23
L23:
	;
	v121 = F_range_cmp_bounds(m, v37, v41+int32(24), v41+int32(8))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L30
	}
L24:
	;
	v91 = v87 + int32(1)
	if v43 <= v91 {
		v158 = v61
		goto L12
	} else {
		goto L26
	}
L25:
	;
	v114 = v91
	goto L23
L26:
	;
	F_multirange_get_bounds(m, v37, v16, v91, v41+int32(24), v41+int32(16))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v103 = F_range_cmp_bounds(m, v37, v41+int32(16), v41+int32(8))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v103 < int32(0) {
		v87 = v91
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
	v128 = F_range_cmp_bounds(m, v37, v41+int32(24), v41)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v137 = F_range_cmp_bounds(m, v37, v41+int32(8), v41+int32(24))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	if v128 <= int32(0) {
		v158 = int32(1)
		goto L12
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	if int32(0) <= v137 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v146 = F_range_cmp_bounds(m, v37, v41+int32(8), v41+int32(16))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v151 = int32(0)
	v153 = v64 + int32(1)
	if v153 != v46 {
		v61 = v151
		v64 = v153
		v65 = v114
		goto L17
	} else {
		goto L42
	}
L40:
	;
	if v146 <= int32(0) {
		v158 = int32(1)
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
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v23
	F_errmsg_internal(m, int32(389826), v13)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(519712), int32(558), int32(419268))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
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
