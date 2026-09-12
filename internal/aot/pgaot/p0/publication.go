package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetPublicationByName(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = F_get_publication_oid(m, l0, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			return int32(0)
		} else {
			v11 = F_GetPublication(m, v3)
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
func F_publicationListToArray(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v12 = F_AllocSetContextCreateInternal(m, v7, int32(25116), v2, int32(8192), int32(8388608))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(4515488)
		v17 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v12
		if l0 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v23 = v20 << (uint(int32(2)) % 32)
		} else {
			v23 = v2
		}
		v24 = F_palloc(m, v23)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			F_check_duplicates_in_publist(m, l0, v24)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
				if l0 != 0 {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v31 = v30
				} else {
					v31 = v2
				}
				v33 = F_construct_array_builtin(m, v24, v31, int32(25))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_MemoryContextDelete(m, v12)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						return v33
					}
				}
			}
		}
	}
}
func F_publication_add_schema(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = F_GetPublication(m, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v16 = F_table_open(m, int32(6237), int32(3))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v19 = int32(0)
			v21 = F_SearchSysCacheExists(m, int32(50), l2, l1, v19, v19)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				if v21 != 0 {
					F_sequence_close(m, v16, int32(3))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						if l3 != 0 {
							v27 = *(*int64)(unsafe.Add(mBase, _consts[224]))
							*(*int64)(unsafe.Add(mBase, uint32(l0))) = v27
							v30 = *(*int32)(unsafe.Add(mBase, _consts[225]))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v30
							m.G0 = v10 + int32(80)
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								F_errcode(m, int32(290948))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									v39 = F_get_namespace_name(m, l2)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v41
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v39
										F_errmsg(m, int32(709417), v10)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											F_errfinish(m, int32(496515), int32(701), int32(506365))
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return
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
				} else {
					if l2 == int32(11) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return
							} else {
								v139 = F_get_namespace_name(m, l2)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v139
									F_errmsg(m, int32(266743), v10+int32(16))
									mBase = m.M
									v146 = m.ExcPending
									if v146 != 0 {
										return
									} else {
										F_errdetail(m, int32(602833), int32(0))
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return
										} else {
											F_errfinish(m, int32(496515), int32(103), int32(506359))
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return
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
					} else {
						if l2 != int32(99) {
							v58 = F_isTempToastNamespace(m, l2)
							mBase = m.M
							v59 = v58
						} else {
							v59 = int32(1)
						}
						if v59 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return
								} else {
									v139 = F_get_namespace_name(m, l2)
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v139
										F_errmsg(m, int32(266743), v10+int32(16))
										mBase = m.M
										v146 = m.ExcPending
										if v146 != 0 {
											return
										} else {
											F_errdetail(m, int32(602833), int32(0))
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return
											} else {
												F_errfinish(m, int32(496515), int32(103), int32(506359))
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return
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
						} else {
							v60 = F_isAnyTempNamespace(m, l2)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								if v60 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
										return
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v162 = m.ExcPending
										if v162 != 0 {
											return
										} else {
											v163 = F_get_namespace_name(m, l2)
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v163
												F_errmsg(m, int32(266743), v10+int32(32))
												mBase = m.M
												v170 = m.ExcPending
												if v170 != 0 {
													return
												} else {
													F_errdetail(m, int32(648448), int32(0))
													mBase = m.M
													v174 = m.ExcPending
													if v174 != 0 {
														return
													} else {
														F_errfinish(m, int32(496515), int32(111), int32(506359))
														mBase = m.M
														v179 = m.ExcPending
														if v179 != 0 {
															return
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
								} else {
									v62 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+66)) = uint8(v62)
									*(*uint16)(unsafe.Add(mBase, uint32(v10)+64)) = uint16(v62)
									v68 = F_GetNewOidWithIndex(m, v16, int32(6238), int32(1))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+76)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v68
										v73 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
										v78 = F_heap_form_tuple(m, v73, v10+int32(68), v10-int32(-64))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return
										} else {
											F_CatalogTupleInsert(m, v16, v78)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return
											} else {
												F_pfree(m, v78)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													v85 = v10 + int32(60)
													v86 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v85))) = v86
													*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v68
													*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = int32(6237)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v86
													*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(6104)
													F_recordDependencyOn(m, v10+int32(52), v10+int32(40), int32(97))
													mBase = m.M
													v102 = m.ExcPending
													if v102 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(2615)
														F_recordDependencyOn(m, v10+int32(52), v10+int32(40), int32(97))
														mBase = m.M
														v114 = m.ExcPending
														if v114 != 0 {
															return
														} else {
															F_sequence_close(m, v16, int32(3))
															mBase = m.M
															v117 = m.ExcPending
															if v117 != 0 {
																return
															} else {
																v119 = F_GetSchemaPublicationRelations(m, l2, int32(2))
																mBase = m.M
																v120 = m.ExcPending
																if v120 != 0 {
																	return
																} else {
																	F_InvalidatePublicationRels(m, v119)
																	mBase = m.M
																	v122 = m.ExcPending
																	if v122 != 0 {
																		return
																	} else {
																		v123 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v123
																		v125 = *(*int64)(unsafe.Add(mBase, uint32(v10)+52))
																		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v125
																		m.G0 = v10 + int32(80)
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
	}
}
