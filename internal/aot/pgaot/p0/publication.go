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
func F_publicationListToArray(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_publicationListToArray[0]))
	v11 = F_AllocSetContextCreateInternal(m, v6, int32(_a_F_publicationListToArray_0), int32(0), int32(_a_F_publicationListToArray_1), int32(_a_F_publicationListToArray_2))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = int32(_a_F_publicationListToArray_3)
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_publicationListToArray[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_publicationListToArray[0])) = v11
		if l0 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v22 = v20
		} else {
			v22 = int32(0)
		}
		v23 = F_palloc_mul(m, int32(8), v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			F_check_duplicates_in_publist(m, l0, v23)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_publicationListToArray[0])) = v16
				if l0 != 0 {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v31 = v29
				} else {
					v31 = int32(0)
				}
				v33 = F_construct_array_builtin(m, v23, v31, int32(25))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
				} else {
					F_MemoryContextDelete(m, v11)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(v33)
					}
				}
			}
		}
	}
}
func F_publication_add_schema(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int64
	_ = v123
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v14 = F_GetPublication(m, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v18 = F_table_open(m, int32(_a_F_publication_add_schema_0), int32(3))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v21 = base.I64_extend_i32_u(l2)
			v22 = base.I64_extend_i32_u(l1)
			v23 = int64(0)
			v25 = F_SearchSysCacheExists(m, int32(50), v21, v22, v23, v23)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				if v25 != 0 {
					F_relation_close(m, v18, int32(3))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						if l3 != 0 {
							v31 = *(*int32)(unsafe.Add(mBase, _c_F_publication_add_schema[0]))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v31
							v34 = *(*int64)(unsafe.Add(mBase, _c_F_publication_add_schema[1]))
							*(*int64)(unsafe.Add(mBase, uint32(l0))) = v34
							m.G0 = v12 + int32(96)
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								F_errcode(m, int32(_a_F_publication_add_schema_1))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									v43 = F_get_namespace_name(m, l2)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return
									} else {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v45
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = v43
										F_errmsg(m, int32(_a_F_publication_add_schema_2), v12)
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_publication_add_schema_3), int32(839), int32(_a_F_publication_add_schema_4))
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
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
						v134 = m.ExcPending
						if v134 != 0 {
							return
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return
							} else {
								v138 = F_get_namespace_name(m, l2)
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v138
									F_errmsg(m, int32(_a_F_publication_add_schema_5), v12+int32(16))
									mBase = m.M
									v145 = m.ExcPending
									if v145 != 0 {
										return
									} else {
										v148 = F_errdetail(m, int32(_a_F_publication_add_schema_6), int32(0))
										mBase = m.M
										v149 = m.ExcPending
										if v149 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_publication_add_schema_3), int32(137), int32(_a_F_publication_add_schema_7))
											mBase = m.M
											v154 = m.ExcPending
											if v154 != 0 {
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
							v60 = F_isTempToastNamespace(m, l2)
							mBase = m.M
							v62 = v60
						} else {
							v62 = int32(1)
						}
						if v62 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v137 = m.ExcPending
								if v137 != 0 {
									return
								} else {
									v138 = F_get_namespace_name(m, l2)
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v138
										F_errmsg(m, int32(_a_F_publication_add_schema_5), v12+int32(16))
										mBase = m.M
										v145 = m.ExcPending
										if v145 != 0 {
											return
										} else {
											v148 = F_errdetail(m, int32(_a_F_publication_add_schema_6), int32(0))
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_publication_add_schema_3), int32(137), int32(_a_F_publication_add_schema_7))
												mBase = m.M
												v154 = m.ExcPending
												if v154 != 0 {
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
							v63 = F_isAnyTempNamespace(m, l2)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								if v63 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v158 = m.ExcPending
									if v158 != 0 {
										return
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return
										} else {
											v162 = F_get_namespace_name(m, l2)
											mBase = m.M
											v163 = m.ExcPending
											if v163 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v162
												F_errmsg(m, int32(_a_F_publication_add_schema_5), v12+int32(32))
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
													return
												} else {
													v172 = F_errdetail(m, int32(_a_F_publication_add_schema_8), int32(0))
													mBase = m.M
													v173 = m.ExcPending
													if v173 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_publication_add_schema_3), int32(145), int32(_a_F_publication_add_schema_7))
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
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
									v65 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v12)+62)) = uint8(v65)
									*(*uint16)(unsafe.Add(mBase, uint32(v12)+60)) = uint16(v65)
									v71 = F_GetNewOidWithIndex(m, v18, int32(_a_F_publication_add_schema_9), int32(1))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v21
										*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = v22
										*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = base.I64_extend_i32_u(v71)
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
										v82 = F_heap_form_tuple(m, v77, v12-int32(-64), v12+int32(60))
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											F_CatalogTupleInsert(m, v18, v82)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												F_pfree(m, v82)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return
												} else {
													v88 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v88
													*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v71
													*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(_a_F_publication_add_schema_0)
													*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v88
													*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(_a_F_publication_add_schema_10)
													v99 = v12 + int32(48)
													v101 = v12 + int32(36)
													F_recordDependencyOn(m, v99, v101, int32(97))
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(2615)
														F_recordDependencyOn(m, v99, v101, int32(97))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															F_relation_close(m, v18, int32(3))
															mBase = m.M
															v115 = m.ExcPending
															if v115 != 0 {
																return
															} else {
																v117 = F_GetSchemaPublicationRelations(m, l2, int32(2))
																mBase = m.M
																v118 = m.ExcPending
																if v118 != 0 {
																	return
																} else {
																	F_InvalidatePublicationRels(m, v117)
																	mBase = m.M
																	v120 = m.ExcPending
																	if v120 != 0 {
																		return
																	} else {
																		v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v121
																		v123 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
																		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v123
																		m.G0 = v12 + int32(96)
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
