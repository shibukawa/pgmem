package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_NamespaceCreate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v33 int64
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	if l0 != 0 {
		v14 = int32(0)
		v17 = F_SearchSysCacheExists(m, int32(37), l0, v14, v14, v14)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v17 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v113 = m.ExcPending
				if v113 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(100794500))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
						F_errmsg(m, int32(115188), v11)
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(490955), int32(64), int32(349743))
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
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
					v25 = F_get_user_default_acl(m, int32(36), l1, int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = v25
						v30 = F_table_open(m, int32(2615), int32(3))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+52))
							v33 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = v33
							*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = v33
							v41 = F_GetNewOidWithIndex(m, v30, int32(2685), int32(1))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v41
								v47 = F_strncpy(m, v11+int32(16), l0, int32(64))
								mBase = m.M
								v48 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v47)+63)) = uint8(v48)
								*(*int32)(unsafe.Add(mBase, uint32(v11)+88)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v11 + int32(16)
								if v27 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = v27
								} else {
									v55 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v11)+111)) = uint8(v55)
								}
								v61 = F_heap_form_tuple(m, v32, v11+int32(80), v11+int32(108))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_CatalogTupleInsert(m, v30, v61)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										F_sequence_close(m, v30, int32(3))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v41
											v71 = int32(2615)
											*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v71
											F_recordDependencyOnOwner(m, v71, v41, l1)
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return int32(0)
											} else {
												F_recordDependencyOnNewAcl(m, int32(2615), v41, l1, v27)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													if l2 == int32(0) {
														F_recordDependencyOnCurrentExtension(m, v11+int32(4), int32(0))
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int32(0)
														} else {
															v87 = *(*int32)(unsafe.Add(mBase, _consts[297]))
															if v87 != 0 {
																v89 = int32(0)
																F_RunObjectPostCreateHook(m, int32(2615), v41, v89, v89)
																mBase = m.M
																v92 = m.ExcPending
																if v92 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v11 + int32(112)
																	return v41
																}
															} else {
																m.G0 = v11 + int32(112)
																return v41
															}
														}
													} else {
														v87 = *(*int32)(unsafe.Add(mBase, _consts[297]))
														if v87 != 0 {
															v89 = int32(0)
															F_RunObjectPostCreateHook(m, int32(2615), v41, v89, v89)
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return int32(0)
															} else {
																m.G0 = v11 + int32(112)
																return v41
															}
														} else {
															m.G0 = v11 + int32(112)
															return v41
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
					v27 = int32(0)
					v30 = F_table_open(m, int32(2615), int32(3))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+52))
						v33 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = v33
						*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = v33
						v41 = F_GetNewOidWithIndex(m, v30, int32(2685), int32(1))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v41
							v47 = F_strncpy(m, v11+int32(16), l0, int32(64))
							mBase = m.M
							v48 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v47)+63)) = uint8(v48)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+88)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v11 + int32(16)
							if v27 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = v27
							} else {
								v55 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v11)+111)) = uint8(v55)
							}
							v61 = F_heap_form_tuple(m, v32, v11+int32(80), v11+int32(108))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_CatalogTupleInsert(m, v30, v61)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									F_sequence_close(m, v30, int32(3))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v41
										v71 = int32(2615)
										*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v71
										F_recordDependencyOnOwner(m, v71, v41, l1)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											F_recordDependencyOnNewAcl(m, int32(2615), v41, l1, v27)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												if l2 == int32(0) {
													F_recordDependencyOnCurrentExtension(m, v11+int32(4), int32(0))
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return int32(0)
													} else {
														v87 = *(*int32)(unsafe.Add(mBase, _consts[297]))
														if v87 != 0 {
															v89 = int32(0)
															F_RunObjectPostCreateHook(m, int32(2615), v41, v89, v89)
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return int32(0)
															} else {
																m.G0 = v11 + int32(112)
																return v41
															}
														} else {
															m.G0 = v11 + int32(112)
															return v41
														}
													}
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, _consts[297]))
													if v87 != 0 {
														v89 = int32(0)
														F_RunObjectPostCreateHook(m, int32(2615), v41, v89, v89)
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return int32(0)
														} else {
															m.G0 = v11 + int32(112)
															return v41
														}
													} else {
														m.G0 = v11 + int32(112)
														return v41
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
		v100 = m.ExcPending
		if v100 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(448974), int32(0))
			mBase = m.M
			v104 = m.ExcPending
			if v104 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(490955), int32(58), int32(349743))
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
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
