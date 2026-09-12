package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_tablespace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l0
	if l0 == v2 {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[171]))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v13
	} else {
	}
	v16 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
	if v16 != 0 {
		v43 = v16
		v46 = int32(0)
		v48 = F_hash_search(m, v43, v5+int32(-52), v46, v46)
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int32(0)
		} else {
			if v48 == int32(0) {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v54 = F_SearchSysCache1(m, int32(69), v53)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					if v54 != 0 {
						v60 = F_SysCacheGetAttr(m, int32(69), v54, int32(5), v5+int32(-48))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)))
							if v62 == int32(0) {
								v66 = F_tablespace_reloptions(m, v60, int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, _consts[299]))
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
									v73 = F_MemoryContextAlloc(m, v69, int32(base.Ui32(v70)>>(uint(int32(2))%32)))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
										v77 = int32(base.Ui32(v75) >> (uint(int32(2)) % 32))
										if v77 != 0 {
											v78 = F__emscripten_memcpy_bulkmem(m, v73, v66, v77)
											mBase = m.M
										} else {
										}
										v81 = v73
										F_ReleaseCatCache(m, v54)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int32(0)
										} else {
											v85 = v81
											v87 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
											v92 = F_hash_search(m, v87, v5+int32(-52), int32(1), int32(0))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v85
												v95 = v92
												m.G0 = v7 - int32(-64)
												return v95
											}
										}
									}
								}
							} else {
								v81 = v2
								F_ReleaseCatCache(m, v54)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int32(0)
								} else {
									v85 = v81
									v87 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
									v92 = F_hash_search(m, v87, v5+int32(-52), int32(1), int32(0))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v85
										v95 = v92
										m.G0 = v7 - int32(-64)
										return v95
									}
								}
							}
						}
					} else {
						v85 = v2
						v87 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
						v92 = F_hash_search(m, v87, v5+int32(-52), int32(1), int32(0))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v85
							v95 = v92
							m.G0 = v7 - int32(-64)
							return v95
						}
					}
				}
			} else {
				v95 = v48
				m.G0 = v7 - int32(-64)
				return v95
			}
		}
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = int64(34359738372)
		v25 = F_hash_create(m, int32(400697), int32(16), v5+int32(-48), int32(40))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[1167])) = v25
			v31 = *(*int32)(unsafe.Add(mBase, _consts[299]))
			if v31 == int32(0) {
				F_CreateCacheMemoryContext(m)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					F_CacheRegisterSyscacheCallback(m, int32(69), int32(1611), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
						v43 = v42
						v46 = int32(0)
						v48 = F_hash_search(m, v43, v5+int32(-52), v46, v46)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							if v48 == int32(0) {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
								v54 = F_SearchSysCache1(m, int32(69), v53)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									if v54 != 0 {
										v60 = F_SysCacheGetAttr(m, int32(69), v54, int32(5), v5+int32(-48))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)))
											if v62 == int32(0) {
												v66 = F_tablespace_reloptions(m, v60, int32(0))
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return int32(0)
												} else {
													v69 = *(*int32)(unsafe.Add(mBase, _consts[299]))
													v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
													v73 = F_MemoryContextAlloc(m, v69, int32(base.Ui32(v70)>>(uint(int32(2))%32)))
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return int32(0)
													} else {
														v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
														v77 = int32(base.Ui32(v75) >> (uint(int32(2)) % 32))
														if v77 != 0 {
															v78 = F__emscripten_memcpy_bulkmem(m, v73, v66, v77)
															mBase = m.M
														} else {
														}
														v81 = v73
														F_ReleaseCatCache(m, v54)
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
															return int32(0)
														} else {
															v85 = v81
															v87 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
															v92 = F_hash_search(m, v87, v5+int32(-52), int32(1), int32(0))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v85
																v95 = v92
																m.G0 = v7 - int32(-64)
																return v95
															}
														}
													}
												}
											} else {
												v81 = v2
												F_ReleaseCatCache(m, v54)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													v85 = v81
													v87 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
													v92 = F_hash_search(m, v87, v5+int32(-52), int32(1), int32(0))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v85
														v95 = v92
														m.G0 = v7 - int32(-64)
														return v95
													}
												}
											}
										}
									} else {
										v85 = v2
										v87 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
										v92 = F_hash_search(m, v87, v5+int32(-52), int32(1), int32(0))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v85
											v95 = v92
											m.G0 = v7 - int32(-64)
											return v95
										}
									}
								}
							} else {
								v95 = v48
								m.G0 = v7 - int32(-64)
								return v95
							}
						}
					}
				}
			} else {
				F_CacheRegisterSyscacheCallback(m, int32(69), int32(1611), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
					v43 = v42
					v46 = int32(0)
					v48 = F_hash_search(m, v43, v5+int32(-52), v46, v46)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						if v48 == int32(0) {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
							v54 = F_SearchSysCache1(m, int32(69), v53)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								if v54 != 0 {
									v60 = F_SysCacheGetAttr(m, int32(69), v54, int32(5), v5+int32(-48))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)))
										if v62 == int32(0) {
											v66 = F_tablespace_reloptions(m, v60, int32(0))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												v69 = *(*int32)(unsafe.Add(mBase, _consts[299]))
												v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
												v73 = F_MemoryContextAlloc(m, v69, int32(base.Ui32(v70)>>(uint(int32(2))%32)))
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
													v77 = int32(base.Ui32(v75) >> (uint(int32(2)) % 32))
													if v77 != 0 {
														v78 = F__emscripten_memcpy_bulkmem(m, v73, v66, v77)
														mBase = m.M
													} else {
													}
													v81 = v73
													F_ReleaseCatCache(m, v54)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int32(0)
													} else {
														v85 = v81
														v87 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
														v92 = F_hash_search(m, v87, v5+int32(-52), int32(1), int32(0))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v85
															v95 = v92
															m.G0 = v7 - int32(-64)
															return v95
														}
													}
												}
											}
										} else {
											v81 = v2
											F_ReleaseCatCache(m, v54)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												v85 = v81
												v87 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
												v92 = F_hash_search(m, v87, v5+int32(-52), int32(1), int32(0))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v85
													v95 = v92
													m.G0 = v7 - int32(-64)
													return v95
												}
											}
										}
									}
								} else {
									v85 = v2
									v87 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
									v92 = F_hash_search(m, v87, v5+int32(-52), int32(1), int32(0))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v85
										v95 = v92
										m.G0 = v7 - int32(-64)
										return v95
									}
								}
							}
						} else {
							v95 = v48
							m.G0 = v7 - int32(-64)
							return v95
						}
					}
				}
			}
		}
	}
}
func F_has_tablespace_privilege_id_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v18)
		v22 = F_convert_any_priv_string(m, v14, int32(1661424))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = F_object_aclcheck_ext(m, int32(1213), v11, v12, v22, v9+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
				if v28 == int32(1) {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
					v35 = int32(0)
				} else {
					v35 = base.B2i32(v26 == int32(0))
				}
				m.G0 = v9 + int32(16)
				return v35
			}
		}
	}
}
