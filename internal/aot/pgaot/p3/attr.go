package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_StoreAttrDefault(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	v10 = m.G0
	v12 = v10 - int32(224)
	m.G0 = v12
	v19 = F__emscripten_memset_bulkmem(m, v12+int32(96), base.I32_extend8_s(int32(0)), int32(100))
	mBase = m.M
	v20 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+88)) = uint8(v20)
	v22 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+56)) = uint8(v20)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v22
	v38 = F_table_open(m, int32(2604), int32(3))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		return int32(0)
	} else {
		v42 = F_nodeToString(m, l2)
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			v46 = F_GetNewOidWithIndex(m, v38, int32(2657), int32(1))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+208)) = v46
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				*(*int32)(unsafe.Add(mBase, uint32(v12)+216)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v12)+212)) = v49
				v52 = F_cstring_to_text(m, v42)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+220)) = v52
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
					v59 = F_heap_form_tuple(m, v55, v12+int32(208), int32(4372956))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						F_CatalogTupleInsert(m, v38, v59)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v46
							*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(2604)
							F_sequence_close(m, v38, int32(3))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v12)+220))
								F_pfree(m, v71)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v59)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v42)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											v80 = F_table_open(m, int32(1249), int32(3))
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
												v84 = F_SearchSysCacheCopy(m, int32(7), v83, l1)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int32(0)
												} else {
													if v84 != 0 {
														v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
														v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+22)))
														v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+v87)+90)))
														v90 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)) = uint8(v90)
														*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v90
														v94 = *(*int32)(unsafe.Add(mBase, uint32(v80)+52))
														v101 = F_heap_modify_tuple(m, v84, v94, v12+int32(96), v12-int32(-64), v12+int32(32))
														mBase = m.M
														v102 = m.ExcPending
														if v102 != 0 {
															return int32(0)
														} else {
															F_CatalogTupleUpdate(m, v80, v101+int32(4), v101)
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return int32(0)
															} else {
																F_sequence_close(m, v80, int32(3))
																mBase = m.M
																v109 = m.ExcPending
																if v109 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v101)
																	mBase = m.M
																	v111 = m.ExcPending
																	if v111 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(1259)
																		v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l1
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v114
																		if v89 != 0 {
																			v123 = int32(105)
																		} else {
																			v123 = int32(97)
																		}
																		F_recordDependencyOn(m, v12+int32(8), v12+int32(20), v123)
																		mBase = m.M
																		v125 = m.ExcPending
																		if v125 != 0 {
																			return int32(0)
																		} else {
																			v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																			F_recordDependencyOnSingleRelExpr(m, v12+int32(8), l2, v128, int32(110), int32(0))
																			mBase = m.M
																			v132 = m.ExcPending
																			if v132 != 0 {
																				return int32(0)
																			} else {
																				v134 = *(*int32)(unsafe.Add(mBase, _consts[433]))
																				if v134 != 0 {
																					v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																					F_RunObjectPostCreateHook(m, int32(2604), v136, l1, l3)
																					mBase = m.M
																					v138 = m.ExcPending
																					if v138 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v12 + int32(224)
																						return v46
																					}
																				} else {
																					m.G0 = v12 + int32(224)
																					return v46
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
														v146 = m.ExcPending
														if v146 != 0 {
															return int32(0)
														} else {
															v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
															*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v147
															*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
															F_errmsg_internal(m, int32(46203), v12)
															mBase = m.M
															v152 = m.ExcPending
															if v152 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(492835), int32(95), int32(97771))
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
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
