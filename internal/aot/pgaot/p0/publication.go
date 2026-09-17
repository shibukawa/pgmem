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
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_publicationListToArray[0]))
	v11 = F_AllocSetContextCreateInternal(m, v6, int32(_a_F_publicationListToArray_0), int32(0), int32(_a_F_publicationListToArray_1), int32(_a_F_publicationListToArray_2))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(_a_F_publicationListToArray_3)
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_publicationListToArray[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_publicationListToArray[0])) = v11
		if l0 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v23 = v19 << (uint(int32(2)) % 32)
		} else {
			v23 = int32(0)
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
				*(*int32)(unsafe.Add(mBase, _c_F_publicationListToArray[0])) = v16
				if l0 != 0 {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v32 = v30
				} else {
					v32 = int32(0)
				}
				v34 = F_construct_array_builtin(m, v24, v32, int32(25))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					F_MemoryContextDelete(m, v11)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						return v34
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
	var v27 int32
	_ = v27
	var v30 int64
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
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = F_GetPublication(m, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v16 = F_table_open(m, int32(_a_F_publication_add_schema_0), int32(3))
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
					F_relation_close(m, v16, int32(3))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						if l3 != 0 {
							v27 = *(*int32)(unsafe.Add(mBase, _c_F_publication_add_schema[0]))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v27
							v30 = *(*int64)(unsafe.Add(mBase, _c_F_publication_add_schema[1]))
							*(*int64)(unsafe.Add(mBase, uint32(l0))) = v30
							m.G0 = v10 + int32(80)
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								F_errcode(m, int32(_a_F_publication_add_schema_1))
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
										F_errmsg(m, int32(_a_F_publication_add_schema_2), v10)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_publication_add_schema_3), int32(701), int32(_a_F_publication_add_schema_4))
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
						v129 = m.ExcPending
						if v129 != 0 {
							return
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return
							} else {
								v133 = F_get_namespace_name(m, l2)
								mBase = m.M
								v134 = m.ExcPending
								if v134 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v133
									F_errmsg(m, int32(_a_F_publication_add_schema_5), v10+int32(16))
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return
									} else {
										F_errdetail(m, int32(_a_F_publication_add_schema_6), int32(0))
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_publication_add_schema_3), int32(103), int32(_a_F_publication_add_schema_7))
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
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
							v56 = F_isTempToastNamespace(m, l2)
							mBase = m.M
							v58 = v56
						} else {
							v58 = int32(1)
						}
						if v58 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return
								} else {
									v133 = F_get_namespace_name(m, l2)
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v133
										F_errmsg(m, int32(_a_F_publication_add_schema_5), v10+int32(16))
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return
										} else {
											F_errdetail(m, int32(_a_F_publication_add_schema_6), int32(0))
											mBase = m.M
											v144 = m.ExcPending
											if v144 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_publication_add_schema_3), int32(103), int32(_a_F_publication_add_schema_7))
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
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
							v59 = F_isAnyTempNamespace(m, l2)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								if v59 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return
										} else {
											v157 = F_get_namespace_name(m, l2)
											mBase = m.M
											v158 = m.ExcPending
											if v158 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v157
												F_errmsg(m, int32(_a_F_publication_add_schema_5), v10+int32(32))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return
												} else {
													F_errdetail(m, int32(_a_F_publication_add_schema_8), int32(0))
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_publication_add_schema_3), int32(111), int32(_a_F_publication_add_schema_7))
														mBase = m.M
														v173 = m.ExcPending
														if v173 != 0 {
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
									v61 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+66)) = uint8(v61)
									*(*uint16)(unsafe.Add(mBase, uint32(v10)+64)) = uint16(v61)
									v67 = F_GetNewOidWithIndex(m, v16, int32(_a_F_publication_add_schema_9), int32(1))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+76)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v67
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
										v77 = F_heap_form_tuple(m, v72, v10+int32(68), v10-int32(-64))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											F_CatalogTupleInsert(m, v16, v77)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												F_pfree(m, v77)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return
												} else {
													v83 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v83
													*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v67
													*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = int32(_a_F_publication_add_schema_0)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v83
													*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(_a_F_publication_add_schema_10)
													v94 = v10 + int32(52)
													v96 = v10 + int32(40)
													F_recordDependencyOn(m, v94, v96, int32(97))
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(2615)
														F_recordDependencyOn(m, v94, v96, int32(97))
														mBase = m.M
														v107 = m.ExcPending
														if v107 != 0 {
															return
														} else {
															F_relation_close(m, v16, int32(3))
															mBase = m.M
															v110 = m.ExcPending
															if v110 != 0 {
																return
															} else {
																v112 = F_GetSchemaPublicationRelations(m, l2, int32(2))
																mBase = m.M
																v113 = m.ExcPending
																if v113 != 0 {
																	return
																} else {
																	F_InvalidatePublicationRels(m, v112)
																	mBase = m.M
																	v115 = m.ExcPending
																	if v115 != 0 {
																		return
																	} else {
																		v116 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v116
																		v118 = *(*int64)(unsafe.Add(mBase, uint32(v10)+52))
																		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v118
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
