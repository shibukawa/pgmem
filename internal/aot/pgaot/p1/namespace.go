package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_NamespaceCreate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v35 int64
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v10 = m.G0
	v12 = v10 - int32(128)
	m.G0 = v12
	if l0 != 0 {
		v16 = int64(0)
		v19 = F_SearchSysCacheExists(m, int32(37), base.I64_extend_i32_u(l0), v16, v16, v16)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if v19 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(100794500))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
						F_errmsg(m, int32(_a_F_NamespaceCreate_0), v12)
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_NamespaceCreate_1), int32(64), int32(_a_F_NamespaceCreate_2))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				if l2 == int32(0) {
					v27 = F_get_user_default_acl(m, int32(37), l1, int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = v27
						v32 = F_table_open(m, int32(2615), int32(3))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
							v35 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v12)+104)) = v35
							v43 = F_GetNewOidWithIndex(m, v32, int32(2685), int32(1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = base.I64_extend_i32_u(v43)
								v48 = v12 + int32(16)
								v50 = F_strncpy(m, v48, l0, int32(64))
								mBase = m.M
								v51 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v50)+63)) = uint8(v51)
								*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = base.I64_extend_i32_u(l1)
								*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = base.I64_extend_i32_u(v48)
								if v29 != 0 {
									*(*int64)(unsafe.Add(mBase, uint32(v12)+104)) = base.I64_extend_i32_u(v29)
								} else {
									v59 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)) = uint8(v59)
								}
								v65 = F_heap_form_tuple(m, v34, v12+int32(80), v12+int32(124))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									F_CatalogTupleInsert(m, v32, v65)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										F_relation_close(m, v32, int32(3))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v43
											v75 = int32(2615)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v75
											F_recordDependencyOnOwner(m, v75, v43, l1)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												F_recordDependencyOnNewAcl(m, int32(2615), v43, l1, v29)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int32(0)
												} else {
													if l2 == int32(0) {
														F_recordDependencyOnCurrentExtension(m, v12+int32(4), int32(0))
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return int32(0)
														} else {
															v91 = *(*int32)(unsafe.Add(mBase, _c_F_NamespaceCreate[0]))
															if v91 != 0 {
																v93 = int32(0)
																F_RunObjectPostCreateHook(m, int32(2615), v43, v93, v93)
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v12 + int32(128)
																	return v43
																}
															} else {
																m.G0 = v12 + int32(128)
																return v43
															}
														}
													} else {
														v91 = *(*int32)(unsafe.Add(mBase, _c_F_NamespaceCreate[0]))
														if v91 != 0 {
															v93 = int32(0)
															F_RunObjectPostCreateHook(m, int32(2615), v43, v93, v93)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return int32(0)
															} else {
																m.G0 = v12 + int32(128)
																return v43
															}
														} else {
															m.G0 = v12 + int32(128)
															return v43
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
					v29 = int32(0)
					v32 = F_table_open(m, int32(2615), int32(3))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
						v35 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+104)) = v35
						v43 = F_GetNewOidWithIndex(m, v32, int32(2685), int32(1))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = base.I64_extend_i32_u(v43)
							v48 = v12 + int32(16)
							v50 = F_strncpy(m, v48, l0, int32(64))
							mBase = m.M
							v51 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v50)+63)) = uint8(v51)
							*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = base.I64_extend_i32_u(l1)
							*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = base.I64_extend_i32_u(v48)
							if v29 != 0 {
								*(*int64)(unsafe.Add(mBase, uint32(v12)+104)) = base.I64_extend_i32_u(v29)
							} else {
								v59 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)) = uint8(v59)
							}
							v65 = F_heap_form_tuple(m, v34, v12+int32(80), v12+int32(124))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_CatalogTupleInsert(m, v32, v65)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									F_relation_close(m, v32, int32(3))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v43
										v75 = int32(2615)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v75
										F_recordDependencyOnOwner(m, v75, v43, l1)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											F_recordDependencyOnNewAcl(m, int32(2615), v43, l1, v29)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												if l2 == int32(0) {
													F_recordDependencyOnCurrentExtension(m, v12+int32(4), int32(0))
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int32(0)
													} else {
														v91 = *(*int32)(unsafe.Add(mBase, _c_F_NamespaceCreate[0]))
														if v91 != 0 {
															v93 = int32(0)
															F_RunObjectPostCreateHook(m, int32(2615), v43, v93, v93)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return int32(0)
															} else {
																m.G0 = v12 + int32(128)
																return v43
															}
														} else {
															m.G0 = v12 + int32(128)
															return v43
														}
													}
												} else {
													v91 = *(*int32)(unsafe.Add(mBase, _c_F_NamespaceCreate[0]))
													if v91 != 0 {
														v93 = int32(0)
														F_RunObjectPostCreateHook(m, int32(2615), v43, v93, v93)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return int32(0)
														} else {
															m.G0 = v12 + int32(128)
															return v43
														}
													} else {
														m.G0 = v12 + int32(128)
														return v43
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
		v104 = m.ExcPending
		if v104 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_NamespaceCreate_3), int32(0))
			mBase = m.M
			v108 = m.ExcPending
			if v108 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_NamespaceCreate_1), int32(58), int32(_a_F_NamespaceCreate_2))
				mBase = m.M
				v113 = m.ExcPending
				if v113 != 0 {
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
