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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l0
	if l0 == v2 {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v15
	} else {
	}
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
	if v18 != 0 {
		v45 = v18
		v48 = int32(0)
		v50 = F_hash_search(m, v45, v7+int32(-52), v48, v48)
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			if v50 == int32(0) {
				v55 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)))
				v56 = F_SearchSysCache1(m, int32(69), v55)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					if v56 != 0 {
						v62 = F_SysCacheGetAttr(m, int32(69), v56, int32(5), v7+int32(-48))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
							if v64 != 0 {
								v82 = v2
								F_ReleaseCatCache(m, v56)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									v87 = v82
									v91 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
									v96 = F_hash_search(m, v91, v7+int32(-52), int32(1), int32(0))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v87
										v99 = v96
										m.G0 = v9 - int32(-64)
										return v99
									}
								}
							} else {
								v66 = F_tablespace_reloptions(m, v62, int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[2]))
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
									v73 = F_MemoryContextAlloc(m, v69, int32(base.Ui32(v70)>>(uint(int32(2))%32)))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
										v77 = int32(base.Ui32(v75) >> (uint(int32(2)) % 32))
										if v77 == int32(0) {
											v82 = v73
										} else {
											base.MemoryCopy(m, v73, v66, v77)
											v82 = v73
										}
										F_ReleaseCatCache(m, v56)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return int32(0)
										} else {
											v87 = v82
											v91 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
											v96 = F_hash_search(m, v91, v7+int32(-52), int32(1), int32(0))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v87
												v99 = v96
												m.G0 = v9 - int32(-64)
												return v99
											}
										}
									}
								}
							}
						}
					} else {
						v87 = v2
						v91 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
						v96 = F_hash_search(m, v91, v7+int32(-52), int32(1), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v87
							v99 = v96
							m.G0 = v9 - int32(-64)
							return v99
						}
					}
				}
			} else {
				v99 = v50
				m.G0 = v9 - int32(-64)
				return v99
			}
		}
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(34359738372)
		v27 = F_hash_create(m, int32(_a_F_get_tablespace_0), int64(16), v7+int32(-48), int32(40))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1])) = v27
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[2]))
			if v33 == int32(0) {
				F_CreateCacheMemoryContext(m)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_CacheRegisterSyscacheCallback(m, int32(69), int32(1809), int64(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
						v45 = v44
						v48 = int32(0)
						v50 = F_hash_search(m, v45, v7+int32(-52), v48, v48)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							if v50 == int32(0) {
								v55 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)))
								v56 = F_SearchSysCache1(m, int32(69), v55)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									if v56 != 0 {
										v62 = F_SysCacheGetAttr(m, int32(69), v56, int32(5), v7+int32(-48))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
											if v64 != 0 {
												v82 = v2
												F_ReleaseCatCache(m, v56)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int32(0)
												} else {
													v87 = v82
													v91 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
													v96 = F_hash_search(m, v91, v7+int32(-52), int32(1), int32(0))
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v87
														v99 = v96
														m.G0 = v9 - int32(-64)
														return v99
													}
												}
											} else {
												v66 = F_tablespace_reloptions(m, v62, int32(0))
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return int32(0)
												} else {
													v69 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[2]))
													v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
													v73 = F_MemoryContextAlloc(m, v69, int32(base.Ui32(v70)>>(uint(int32(2))%32)))
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return int32(0)
													} else {
														v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
														v77 = int32(base.Ui32(v75) >> (uint(int32(2)) % 32))
														if v77 == int32(0) {
															v82 = v73
														} else {
															base.MemoryCopy(m, v73, v66, v77)
															v82 = v73
														}
														F_ReleaseCatCache(m, v56)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int32(0)
														} else {
															v87 = v82
															v91 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
															v96 = F_hash_search(m, v91, v7+int32(-52), int32(1), int32(0))
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v87
																v99 = v96
																m.G0 = v9 - int32(-64)
																return v99
															}
														}
													}
												}
											}
										}
									} else {
										v87 = v2
										v91 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
										v96 = F_hash_search(m, v91, v7+int32(-52), int32(1), int32(0))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v87
											v99 = v96
											m.G0 = v9 - int32(-64)
											return v99
										}
									}
								}
							} else {
								v99 = v50
								m.G0 = v9 - int32(-64)
								return v99
							}
						}
					}
				}
			} else {
				F_CacheRegisterSyscacheCallback(m, int32(69), int32(1809), int64(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
					v45 = v44
					v48 = int32(0)
					v50 = F_hash_search(m, v45, v7+int32(-52), v48, v48)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						if v50 == int32(0) {
							v55 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)))
							v56 = F_SearchSysCache1(m, int32(69), v55)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if v56 != 0 {
									v62 = F_SysCacheGetAttr(m, int32(69), v56, int32(5), v7+int32(-48))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
										if v64 != 0 {
											v82 = v2
											F_ReleaseCatCache(m, v56)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return int32(0)
											} else {
												v87 = v82
												v91 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
												v96 = F_hash_search(m, v91, v7+int32(-52), int32(1), int32(0))
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v87
													v99 = v96
													m.G0 = v9 - int32(-64)
													return v99
												}
											}
										} else {
											v66 = F_tablespace_reloptions(m, v62, int32(0))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												v69 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[2]))
												v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
												v73 = F_MemoryContextAlloc(m, v69, int32(base.Ui32(v70)>>(uint(int32(2))%32)))
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
													v77 = int32(base.Ui32(v75) >> (uint(int32(2)) % 32))
													if v77 == int32(0) {
														v82 = v73
													} else {
														base.MemoryCopy(m, v73, v66, v77)
														v82 = v73
													}
													F_ReleaseCatCache(m, v56)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return int32(0)
													} else {
														v87 = v82
														v91 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
														v96 = F_hash_search(m, v91, v7+int32(-52), int32(1), int32(0))
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v87
															v99 = v96
															m.G0 = v9 - int32(-64)
															return v99
														}
													}
												}
											}
										}
									}
								} else {
									v87 = v2
									v91 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace[1]))
									v96 = F_hash_search(m, v91, v7+int32(-52), int32(1), int32(0))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v87
										v99 = v96
										m.G0 = v9 - int32(-64)
										return v99
									}
								}
							}
						} else {
							v99 = v50
							m.G0 = v9 - int32(-64)
							return v99
						}
					}
				}
			}
		}
	}
}
func F_has_tablespace_privilege_id_id(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14303(m, l0, int32(_a_F_has_tablespace_privilege_id_id_0), int32(1213))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
