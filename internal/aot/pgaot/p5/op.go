package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__equalOpExpr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v6 != v7 {
		v37 = v3
		return v37
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v9 == int32(0) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			if v17 != v18 {
				v37 = v3
				return v37
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
				if v20 != v21 {
					v37 = v3
					return v37
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v23 != v24 {
						v37 = v3
						return v37
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						if v26 != v27 {
							v37 = v3
							return v37
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
							v31 = F_equal(m, v29, v30)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v37 = v31
								return v37
							}
						}
					}
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v12 == int32(0) {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				if v17 != v18 {
					v37 = v3
					return v37
				} else {
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
					if v20 != v21 {
						v37 = v3
						return v37
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						if v23 != v24 {
							v37 = v3
							return v37
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							if v26 != v27 {
								v37 = v3
								return v37
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
								v31 = F_equal(m, v29, v30)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									v37 = v31
									return v37
								}
							}
						}
					}
				}
			} else {
				if v9 != v12 {
					v37 = v3
					return v37
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					if v17 != v18 {
						v37 = v3
						return v37
					} else {
						v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
						v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
						if v20 != v21 {
							v37 = v3
							return v37
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							if v23 != v24 {
								v37 = v3
								return v37
							} else {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								if v26 != v27 {
									v37 = v3
									return v37
								} else {
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
									v31 = F_equal(m, v29, v30)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return int32(0)
									} else {
										v37 = v31
										return v37
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
func F_getOpFamilyDescription(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v14 = F_SearchSysCache1(m, int32(42), base.I64_extend_i32_u(l1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14 == int32(0) {
			if l2 != 0 {
				m.G0 = v10 + int32(48)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
					F_errmsg_internal(m, int32(_a_F_getOpFamilyDescription_0), v10)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_getOpFamilyDescription_1), int32(_a_F_getOpFamilyDescription_2), int32(_a_F_getOpFamilyDescription_3))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
			v34 = v32 + v33
			v35 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v34)+4)))
			v36 = F_SearchSysCache1(m, int32(2), v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				if v36 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v79
						F_errmsg_internal(m, int32(_a_F_getOpFamilyDescription_4), v10+int32(16))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_getOpFamilyDescription_1), int32(_a_F_getOpFamilyDescription_5), int32(_a_F_getOpFamilyDescription_3))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+22)))
					v44 = F_OpfamilyIsVisibleExt(m, l1, int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						if v44 != 0 {
							v50 = int32(0)
							v53 = F_quote_qualified_identifier(m, v50, v34+int32(8))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v40 + v41 + int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v53
								F_appendStringInfo(m, l0, int32(_a_F_getOpFamilyDescription_6), v10+int32(32))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									F_ReleaseCatCache(m, v36)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										F_ReleaseCatCache(m, v14)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											m.G0 = v10 + int32(48)
											return
										}
									}
								}
							}
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
							v48 = F_get_namespace_name(m, v47)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v50 = v48
								v53 = F_quote_qualified_identifier(m, v50, v34+int32(8))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v40 + v41 + int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v53
									F_appendStringInfo(m, l0, int32(_a_F_getOpFamilyDescription_6), v10+int32(32))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										F_ReleaseCatCache(m, v36)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return
											} else {
												m.G0 = v10 + int32(48)
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
func F_getOpFamilyIdentity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v17 = F_SearchSysCache1(m, int32(42), base.I64_extend_i32_u(l1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		if v17 == int32(0) {
			if l3 != 0 {
				m.G0 = v13 - int32(-64)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
					F_errmsg_internal(m, int32(_a_F_getOpFamilyIdentity_0), v13)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_getOpFamilyIdentity_1), int32(_a_F_getOpFamilyIdentity_2), int32(_a_F_getOpFamilyIdentity_3))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
			v37 = v35 + v36
			v38 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v37)+4)))
			v39 = F_SearchSysCache1(m, int32(2), v38)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				if v39 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return
					} else {
						v104 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v104
						F_errmsg_internal(m, int32(_a_F_getOpFamilyIdentity_4), v11+int32(-48))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_getOpFamilyIdentity_1), int32(_a_F_getOpFamilyIdentity_5), int32(_a_F_getOpFamilyIdentity_3))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v37)+72))
					v46 = F_get_namespace_name_or_temp(m, v45)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						v49 = v37 + int32(8)
						v50 = F_quote_qualified_identifier(m, v46, v49)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v54 = v43 + v44 + int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v54
							*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v50
							F_appendStringInfo(m, l0, int32(_a_F_getOpFamilyIdentity_6), v11+int32(-32))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								if l2 != 0 {
									v62 = F_pstrdup(m, v54)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v62
										v65 = F_pstrdup(m, v46)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v65
											v68 = F_pstrdup(m, v49)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v68
												*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v68
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
												*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v72
												v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
												*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v74
												v82 = F_list_make3_impl(m, v11+int32(-36), v11+int32(-40), v11+int32(-44))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = v82
													F_ReleaseCatCache(m, v39)
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return
													} else {
														F_ReleaseCatCache(m, v17)
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return
														} else {
															m.G0 = v13 - int32(-64)
															return
														}
													}
												}
											}
										}
									}
								} else {
									F_ReleaseCatCache(m, v39)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										F_ReleaseCatCache(m, v17)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											m.G0 = v13 - int32(-64)
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
func F_get_op_hash_functions_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	if l2 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	} else {
	}
	if l3 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	} else {
	}
	if l0 <= int32(2987) {
		if l0 == int32(1070) {
			v38 = F_lookup_type_cache(m, l1, int32(16))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
				if v40 == int32(626) {
					v53 = F_get_op_hash_functions(m, l0, l2, l3)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						return v53
					}
				} else {
					return int32(0)
				}
			}
		} else {
			if l0 != int32(2860) {
				v53 = F_get_op_hash_functions(m, l0, l2, l3)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					return v53
				}
			} else {
				v16 = F_lookup_type_cache(m, l1, int32(16))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
					if v20 == int32(_a_F_get_op_hash_functions_ext_0) {
						v53 = F_get_op_hash_functions(m, l0, l2, l3)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							return v53
						}
					} else {
						return int32(0)
					}
				}
			}
		}
	} else {
		if l0 == int32(2988) {
			v46 = F_lookup_type_cache(m, l1, int32(16))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
				if v48 == int32(_a_F_get_op_hash_functions_ext_1) {
					v53 = F_get_op_hash_functions(m, l0, l2, l3)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						return v53
					}
				} else {
					return int32(0)
				}
			}
		} else {
			if l0 != int32(3882) {
				v53 = F_get_op_hash_functions(m, l0, l2, l3)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					return v53
				}
			} else {
				v30 = F_lookup_type_cache(m, l1, int32(16))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+68))
					if v32 == int32(3902) {
						v53 = F_get_op_hash_functions(m, l0, l2, l3)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							return v53
						}
					} else {
						return int32(0)
					}
				}
			}
		}
	}
}
func F_get_op_opfamily_sortfamily(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v7 = F_SearchSysCache3(m, int32(3), base.I64_extend_i32_u(l0), int64(111), base.I64_extend_i32_u(l1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 == int32(0) {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15+v16)+28))
			F_ReleaseCatCache(m, v7)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				return v18
			}
		}
	}
}
