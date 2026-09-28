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
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v40 int32
	_ = v40
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
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	v5 = int32(0)
	v11 = int64(0)
	v13 = m.G0
	v15 = v13 - int32(336)
	m.G0 = v15
	v18 = v15 + int32(96)
	base.MemoryFill(m, v18, v5, int32(200))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+88)) = uint8(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = v11
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+56)) = uint8(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = v11
	v40 = F_table_open(m, int32(2604), int32(3))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		return int32(0)
	} else {
		v44 = F_nodeToString(m, l2)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v48 = F_GetNewOidWithIndex(m, v40, int32(2657), int32(1))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v15)+304)) = base.I64_extend_i32_u(v48)
				v52 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
				v53 = base.I64_extend_i32_s(l1)
				*(*int64)(unsafe.Add(mBase, uint32(v15)+320)) = v53
				*(*int64)(unsafe.Add(mBase, uint32(v15)+312)) = v52
				v56 = F_cstring_to_text(m, v44)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v15)+328)) = base.I64_extend_i32_u(v56)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
					v64 = F_heap_form_tuple(m, v60, v15+int32(304), int32(_a_F_StoreAttrDefault_0))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						F_CatalogTupleInsert(m, v40, v64)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v48
							*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(2604)
							F_relation_close(m, v40, int32(3))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+328))
								F_pfree(m, v76)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v64)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v44)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											v85 = F_table_open(m, int32(1249), int32(3))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												v88 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
												v89 = F_SearchSysCacheCopy(m, int32(7), v88, v53)
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return int32(0)
												} else {
													if v89 != 0 {
														v91 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
														v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+22)))
														v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+v92)+90)))
														v95 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v15)+44)) = uint8(v95)
														*(*int64)(unsafe.Add(mBase, uint32(v15)+192)) = int64(1)
														v99 = *(*int32)(unsafe.Add(mBase, uint32(v85)+52))
														v104 = F_heap_modify_tuple(m, v89, v99, v18, v15-int32(-64), v15+int32(32))
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return int32(0)
														} else {
															F_CatalogTupleUpdate(m, v85, v104+int32(4), v104)
															mBase = m.M
															v109 = m.ExcPending
															if v109 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v85, int32(3))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v104)
																	mBase = m.M
																	v114 = m.ExcPending
																	if v114 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(1259)
																		v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = l1
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v117
																		v121 = v15 + int32(8)
																		if v94 != 0 {
																			v126 = int32(105)
																		} else {
																			v126 = int32(97)
																		}
																		F_recordDependencyOn(m, v121, v15+int32(20), v126)
																		mBase = m.M
																		v128 = m.ExcPending
																		if v128 != 0 {
																			return int32(0)
																		} else {
																			v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																			F_recordDependencyOnSingleRelExpr(m, v121, l2, v129, int32(110), int32(0))
																			mBase = m.M
																			v133 = m.ExcPending
																			if v133 != 0 {
																				return int32(0)
																			} else {
																				v135 = *(*int32)(unsafe.Add(mBase, _c_F_StoreAttrDefault[0]))
																				if v135 != 0 {
																					v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																					F_RunObjectPostCreateHook(m, int32(2604), v137, l1, l3)
																					mBase = m.M
																					v139 = m.ExcPending
																					if v139 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v15 + int32(336)
																						return v48
																					}
																				} else {
																					m.G0 = v15 + int32(336)
																					return v48
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
														v147 = m.ExcPending
														if v147 != 0 {
															return int32(0)
														} else {
															v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
															*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v148
															*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
															F_errmsg_internal(m, int32(_a_F_StoreAttrDefault_1), v15)
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_StoreAttrDefault_2), int32(100), int32(_a_F_StoreAttrDefault_3))
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
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
