package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_StoreAttrDefault(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(224)
	m.G0 = v13
	v16 = v13 + int32(96)
	base.MemoryFill(m, v16, v5, int32(100))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+88)) = uint8(v5)
	v22 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+80)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v22
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+56)) = uint8(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v22
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
				*(*int32)(unsafe.Add(mBase, uint32(v13)+208)) = v46
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				*(*int32)(unsafe.Add(mBase, uint32(v13)+216)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v13)+212)) = v49
				v52 = F_cstring_to_text(m, v42)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+220)) = v52
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
					v59 = F_heap_form_tuple(m, v55, v13+int32(208), int32(_a_F_StoreAttrDefault_0))
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
							*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v46
							*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(2604)
							F_relation_close(m, v38, int32(3))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)+220))
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
														*(*uint8)(unsafe.Add(mBase, uint32(v13)+44)) = uint8(v90)
														*(*int32)(unsafe.Add(mBase, uint32(v13)+144)) = v90
														v94 = *(*int32)(unsafe.Add(mBase, uint32(v80)+52))
														v99 = F_heap_modify_tuple(m, v84, v94, v16, v13-int32(-64), v13+int32(32))
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return int32(0)
														} else {
															F_CatalogTupleUpdate(m, v80, v99+int32(4), v99)
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v80, int32(3))
																mBase = m.M
																v107 = m.ExcPending
																if v107 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v99)
																	mBase = m.M
																	v109 = m.ExcPending
																	if v109 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(1259)
																		v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																		*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = l1
																		*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v112
																		v116 = v13 + int32(8)
																		if v89 != 0 {
																			v121 = int32(105)
																		} else {
																			v121 = int32(97)
																		}
																		F_recordDependencyOn(m, v116, v13+int32(20), v121)
																		mBase = m.M
																		v123 = m.ExcPending
																		if v123 != 0 {
																			return int32(0)
																		} else {
																			v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																			F_recordDependencyOnSingleRelExpr(m, v116, l2, v124, int32(110), int32(0))
																			mBase = m.M
																			v128 = m.ExcPending
																			if v128 != 0 {
																				return int32(0)
																			} else {
																				v130 = *(*int32)(unsafe.Add(mBase, _c_F_StoreAttrDefault[0]))
																				if v130 != 0 {
																					v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																					F_RunObjectPostCreateHook(m, int32(2604), v132, l1, l3)
																					mBase = m.M
																					v134 = m.ExcPending
																					if v134 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v13 + int32(224)
																						return v46
																					}
																				} else {
																					m.G0 = v13 + int32(224)
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
														v142 = m.ExcPending
														if v142 != 0 {
															return int32(0)
														} else {
															v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
															*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v143
															*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
															F_errmsg_internal(m, int32(_a_F_StoreAttrDefault_1), v13)
															mBase = m.M
															v148 = m.ExcPending
															if v148 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_StoreAttrDefault_2), int32(95), int32(_a_F_StoreAttrDefault_3))
																mBase = m.M
																v153 = m.ExcPending
																if v153 != 0 {
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
