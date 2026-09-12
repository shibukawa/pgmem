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
	var v35 int32
	_ = v35
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v6 != v7 {
		v35 = v3
		return v35
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v9 == int32(0) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			if v17 != v18 {
				v35 = v3
				return v35
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
				if v20 != v21 {
					v35 = v3
					return v35
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v23 != v24 {
						v35 = v3
						return v35
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						if v26 != v27 {
							v35 = v3
							return v35
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
							v31 = F_equal(m, v29, v30)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v35 = v31
								return v35
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
					v35 = v3
					return v35
				} else {
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
					if v20 != v21 {
						v35 = v3
						return v35
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						if v23 != v24 {
							v35 = v3
							return v35
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							if v26 != v27 {
								v35 = v3
								return v35
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
								v31 = F_equal(m, v29, v30)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									v35 = v31
									return v35
								}
							}
						}
					}
				}
			} else {
				if v9 != v12 {
					v35 = v3
					return v35
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					if v17 != v18 {
						v35 = v3
						return v35
					} else {
						v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
						v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
						if v20 != v21 {
							v35 = v3
							return v35
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							if v23 != v24 {
								v35 = v3
								return v35
							} else {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								if v26 != v27 {
									v35 = v3
									return v35
								} else {
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
									v31 = F_equal(m, v29, v30)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return int32(0)
									} else {
										v35 = v31
										return v35
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = F_SearchSysCache1(m, int32(42), l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 == int32(0) {
			if l2 != 0 {
				m.G0 = v10 + int32(48)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
					F_errmsg_internal(m, int32(39908), v10)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						F_errfinish(m, int32(495158), int32(4190), int32(247534))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
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
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
			v33 = v31 + v32
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
			v35 = F_SearchSysCache1(m, int32(2), v34)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				if v35 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v78
						F_errmsg_internal(m, int32(53551), v10+int32(16))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return
						} else {
							F_errfinish(m, int32(495158), int32(4198), int32(247534))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
					v43 = F_OpfamilyIsVisibleExt(m, l1, int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						if v43 != 0 {
							v49 = int32(0)
							v52 = F_quote_qualified_identifier(m, v49, v33+int32(8))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v39 + v40 + int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v52
								F_appendStringInfo(m, l0, int32(196668), v10+int32(32))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									F_ReleaseCatCache(m, v35)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_ReleaseCatCache(m, v13)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											m.G0 = v10 + int32(48)
											return
										}
									}
								}
							}
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
							v47 = F_get_namespace_name(m, v46)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								v49 = v47
								v52 = F_quote_qualified_identifier(m, v49, v33+int32(8))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v39 + v40 + int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v52
									F_appendStringInfo(m, l0, int32(196668), v10+int32(32))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										F_ReleaseCatCache(m, v35)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											F_ReleaseCatCache(m, v13)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v16 = F_SearchSysCache1(m, int32(42), l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		if v16 == int32(0) {
			if l3 != 0 {
				m.G0 = v13 - int32(-64)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
					F_errmsg_internal(m, int32(39908), v13)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						F_errfinish(m, int32(495158), int32(6066), int32(10763))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
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
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
			v36 = v34 + v35
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
			v38 = F_SearchSysCache1(m, int32(2), v37)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				if v38 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return
					} else {
						v102 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v102
						F_errmsg_internal(m, int32(53551), v11+int32(-48))
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return
						} else {
							F_errfinish(m, int32(495158), int32(6074), int32(10763))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+22)))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v36)+72))
					v45 = F_get_namespace_name_or_temp(m, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						v48 = v36 + int32(8)
						v49 = F_quote_qualified_identifier(m, v45, v48)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							v53 = v42 + v43 + int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v53
							*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v49
							F_appendStringInfo(m, l0, int32(198684), v11+int32(-32))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								if l2 != 0 {
									v61 = F_pstrdup(m, v53)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v61
										v64 = F_pstrdup(m, v45)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v64
											v67 = F_pstrdup(m, v48)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v67
												*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v67
												v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
												*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v71
												v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
												*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v73
												v81 = F_list_make3_impl(m, v11+int32(-36), v11+int32(-40), v11+int32(-44))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = v81
													F_ReleaseCatCache(m, v38)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														F_ReleaseCatCache(m, v16)
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
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
									F_ReleaseCatCache(m, v38)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										F_ReleaseCatCache(m, v16)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
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
func F_get_op_opfamily_sortfamily(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache3(m, int32(3), l0, int32(111), l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14)+28))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v16
			}
		}
	}
}
