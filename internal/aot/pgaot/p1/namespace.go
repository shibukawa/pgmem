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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	if l0 != 0 {
		v15 = int32(0)
		v18 = F_SearchSysCacheExists(m, int32(37), l0, v15, v15, v15)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v18 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v112 = m.ExcPending
				if v112 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(100794500))
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
						F_errmsg(m, int32(_a_F_NamespaceCreate_0), v12)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_NamespaceCreate_1), int32(64), int32(_a_F_NamespaceCreate_2))
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
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
					v26 = F_get_user_default_acl(m, int32(36), l1, int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = v26
						v31 = F_table_open(m, int32(2615), int32(3))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+52))
							v34 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = v34
							v42 = F_GetNewOidWithIndex(m, v31, int32(2685), int32(1))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v42
								v46 = v12 + int32(16)
								v48 = F_strncpy(m, v46, l0, int32(64))
								mBase = m.M
								v49 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v48)+63)) = uint8(v49)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v46
								if v28 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = v28
								} else {
									v54 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v12)+111)) = uint8(v54)
								}
								v60 = F_heap_form_tuple(m, v33, v12+int32(80), v12+int32(108))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									F_CatalogTupleInsert(m, v31, v60)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										F_relation_close(m, v31, int32(3))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v42
											v70 = int32(2615)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v70
											F_recordDependencyOnOwner(m, v70, v42, l1)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												F_recordDependencyOnNewAcl(m, int32(2615), v42, l1, v28)
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return int32(0)
												} else {
													if l2 == int32(0) {
														F_recordDependencyOnCurrentExtension(m, v12+int32(4), int32(0))
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															v86 = *(*int32)(unsafe.Add(mBase, _c_F_NamespaceCreate[0]))
															if v86 != 0 {
																v88 = int32(0)
																F_RunObjectPostCreateHook(m, int32(2615), v42, v88, v88)
																mBase = m.M
																v91 = m.ExcPending
																if v91 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v12 + int32(112)
																	return v42
																}
															} else {
																m.G0 = v12 + int32(112)
																return v42
															}
														}
													} else {
														v86 = *(*int32)(unsafe.Add(mBase, _c_F_NamespaceCreate[0]))
														if v86 != 0 {
															v88 = int32(0)
															F_RunObjectPostCreateHook(m, int32(2615), v42, v88, v88)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int32(0)
															} else {
																m.G0 = v12 + int32(112)
																return v42
															}
														} else {
															m.G0 = v12 + int32(112)
															return v42
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
					v28 = int32(0)
					v31 = F_table_open(m, int32(2615), int32(3))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+52))
						v34 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v34
						*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = v34
						v42 = F_GetNewOidWithIndex(m, v31, int32(2685), int32(1))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v42
							v46 = v12 + int32(16)
							v48 = F_strncpy(m, v46, l0, int32(64))
							mBase = m.M
							v49 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v48)+63)) = uint8(v49)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v46
							if v28 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = v28
							} else {
								v54 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+111)) = uint8(v54)
							}
							v60 = F_heap_form_tuple(m, v33, v12+int32(80), v12+int32(108))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_CatalogTupleInsert(m, v31, v60)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									F_relation_close(m, v31, int32(3))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v42
										v70 = int32(2615)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v70
										F_recordDependencyOnOwner(m, v70, v42, l1)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											F_recordDependencyOnNewAcl(m, int32(2615), v42, l1, v28)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return int32(0)
											} else {
												if l2 == int32(0) {
													F_recordDependencyOnCurrentExtension(m, v12+int32(4), int32(0))
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int32(0)
													} else {
														v86 = *(*int32)(unsafe.Add(mBase, _c_F_NamespaceCreate[0]))
														if v86 != 0 {
															v88 = int32(0)
															F_RunObjectPostCreateHook(m, int32(2615), v42, v88, v88)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int32(0)
															} else {
																m.G0 = v12 + int32(112)
																return v42
															}
														} else {
															m.G0 = v12 + int32(112)
															return v42
														}
													}
												} else {
													v86 = *(*int32)(unsafe.Add(mBase, _c_F_NamespaceCreate[0]))
													if v86 != 0 {
														v88 = int32(0)
														F_RunObjectPostCreateHook(m, int32(2615), v42, v88, v88)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int32(0)
														} else {
															m.G0 = v12 + int32(112)
															return v42
														}
													} else {
														m.G0 = v12 + int32(112)
														return v42
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
		v99 = m.ExcPending
		if v99 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_NamespaceCreate_3), int32(0))
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_NamespaceCreate_1), int32(58), int32(_a_F_NamespaceCreate_2))
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
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
