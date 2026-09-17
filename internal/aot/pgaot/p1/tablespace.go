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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l0
	if l0 == v2 {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v14
	} else {
	}
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
	if v17 != 0 {
		v44 = v17
		v47 = int32(0)
		v49 = F_hash_search(m, v44, v6+int32(-52), v47, v47)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			if v49 == int32(0) {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				v55 = F_SearchSysCache1(m, int32(69), v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					if v55 != 0 {
						v61 = F_SysCacheGetAttr(m, int32(69), v55, int32(5), v6+int32(-48))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
							if v63 != 0 {
								v81 = v2
								F_ReleaseCatCache(m, v55)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									v86 = v81
									v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
									v94 = F_hash_search(m, v89, v6+int32(-52), int32(1), int32(0))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
										v97 = v94
										m.G0 = v8 - int32(-64)
										return v97
									}
								}
							} else {
								v65 = F_tablespace_reloptions(m, v61, int32(0))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[2]))
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
									v72 = F_MemoryContextAlloc(m, v68, int32(base.Ui32(v69)>>(uint(int32(2))%32)))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
										v76 = int32(base.Ui32(v74) >> (uint(int32(2)) % 32))
										if v76 == int32(0) {
											v81 = v72
										} else {
											base.MemoryCopy(m, v72, v65, v76)
											v81 = v72
										}
										F_ReleaseCatCache(m, v55)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int32(0)
										} else {
											v86 = v81
											v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
											v94 = F_hash_search(m, v89, v6+int32(-52), int32(1), int32(0))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
												v97 = v94
												m.G0 = v8 - int32(-64)
												return v97
											}
										}
									}
								}
							}
						}
					} else {
						v86 = v2
						v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
						v94 = F_hash_search(m, v89, v6+int32(-52), int32(1), int32(0))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
							v97 = v94
							m.G0 = v8 - int32(-64)
							return v97
						}
					}
				}
			} else {
				v97 = v49
				m.G0 = v8 - int32(-64)
				return v97
			}
		}
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(34359738372)
		v26 = F_hash_create(m, int32(_a_F_get_tablespace_0), int32(16), v6+int32(-48), int32(40))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1])) = v26
			v32 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[2]))
			if v32 == int32(0) {
				F_CreateCacheMemoryContext(m)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					F_CacheRegisterSyscacheCallback(m, int32(69), int32(1595), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
						v44 = v43
						v47 = int32(0)
						v49 = F_hash_search(m, v44, v6+int32(-52), v47, v47)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v49 == int32(0) {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
								v55 = F_SearchSysCache1(m, int32(69), v54)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									if v55 != 0 {
										v61 = F_SysCacheGetAttr(m, int32(69), v55, int32(5), v6+int32(-48))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
											if v63 != 0 {
												v81 = v2
												F_ReleaseCatCache(m, v55)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													v86 = v81
													v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
													v94 = F_hash_search(m, v89, v6+int32(-52), int32(1), int32(0))
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
														v97 = v94
														m.G0 = v8 - int32(-64)
														return v97
													}
												}
											} else {
												v65 = F_tablespace_reloptions(m, v61, int32(0))
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int32(0)
												} else {
													v68 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[2]))
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
													v72 = F_MemoryContextAlloc(m, v68, int32(base.Ui32(v69)>>(uint(int32(2))%32)))
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return int32(0)
													} else {
														v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
														v76 = int32(base.Ui32(v74) >> (uint(int32(2)) % 32))
														if v76 == int32(0) {
															v81 = v72
														} else {
															base.MemoryCopy(m, v72, v65, v76)
															v81 = v72
														}
														F_ReleaseCatCache(m, v55)
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															v86 = v81
															v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
															v94 = F_hash_search(m, v89, v6+int32(-52), int32(1), int32(0))
															mBase = m.M
															v95 = m.ExcPending
															if v95 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
																v97 = v94
																m.G0 = v8 - int32(-64)
																return v97
															}
														}
													}
												}
											}
										}
									} else {
										v86 = v2
										v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
										v94 = F_hash_search(m, v89, v6+int32(-52), int32(1), int32(0))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
											v97 = v94
											m.G0 = v8 - int32(-64)
											return v97
										}
									}
								}
							} else {
								v97 = v49
								m.G0 = v8 - int32(-64)
								return v97
							}
						}
					}
				}
			} else {
				F_CacheRegisterSyscacheCallback(m, int32(69), int32(1595), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
					v44 = v43
					v47 = int32(0)
					v49 = F_hash_search(m, v44, v6+int32(-52), v47, v47)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						if v49 == int32(0) {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
							v55 = F_SearchSysCache1(m, int32(69), v54)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								if v55 != 0 {
									v61 = F_SysCacheGetAttr(m, int32(69), v55, int32(5), v6+int32(-48))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
										if v63 != 0 {
											v81 = v2
											F_ReleaseCatCache(m, v55)
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return int32(0)
											} else {
												v86 = v81
												v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
												v94 = F_hash_search(m, v89, v6+int32(-52), int32(1), int32(0))
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
													v97 = v94
													m.G0 = v8 - int32(-64)
													return v97
												}
											}
										} else {
											v65 = F_tablespace_reloptions(m, v61, int32(0))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v68 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[2]))
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
												v72 = F_MemoryContextAlloc(m, v68, int32(base.Ui32(v69)>>(uint(int32(2))%32)))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
													v76 = int32(base.Ui32(v74) >> (uint(int32(2)) % 32))
													if v76 == int32(0) {
														v81 = v72
													} else {
														base.MemoryCopy(m, v72, v65, v76)
														v81 = v72
													}
													F_ReleaseCatCache(m, v55)
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int32(0)
													} else {
														v86 = v81
														v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
														v94 = F_hash_search(m, v89, v6+int32(-52), int32(1), int32(0))
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
															v97 = v94
															m.G0 = v8 - int32(-64)
															return v97
														}
													}
												}
											}
										}
									}
								} else {
									v86 = v2
									v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
									v94 = F_hash_search(m, v89, v6+int32(-52), int32(1), int32(0))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
										v97 = v94
										m.G0 = v8 - int32(-64)
										return v97
									}
								}
							}
						} else {
							v97 = v49
							m.G0 = v8 - int32(-64)
							return v97
						}
					}
				}
			}
		}
	}
}
func F_has_tablespace_privilege_id_id(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13914(m, l0, int32(_a_F_has_tablespace_privilege_id_id_0), int32(1213))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
