package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InvalidateCatalogSnapshot(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateCatalogSnapshot[0]))
	if v5 == int32(0) {
		return
	} else {
		F_pairingheap_remove(m, int32(_a_F_InvalidateCatalogSnapshot_0), v5+int32(52))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_InvalidateCatalogSnapshot[0])) = int32(0)
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateCatalogSnapshot[1]))
			if v17 != 0 {
				return
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateCatalogSnapshot[2]))
				if v19 == int32(0) {
					v23 = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_InvalidateCatalogSnapshot[3])) = v23
					v26 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateCatalogSnapshot[4]))
					*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v23
					return
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateCatalogSnapshot[4]))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+52))
					v32 = int32(3)
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateCatalogSnapshot[2]))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v35-int32(48))))
					if base.B2i32(base.Ui32(v31) < base.Ui32(v32))|base.B2i32(base.Ui32(v38) < base.Ui32(v32)) == int32(0) {
						if v31-v38 < int32(0) {
							*(*int32)(unsafe.Add(mBase, _c_F_InvalidateCatalogSnapshot[3])) = v38
							*(*int32)(unsafe.Add(mBase, uint32(v30)+52)) = v38
						} else {
						}
					} else {
						if base.Ui32(v38) <= base.Ui32(v31) {
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_InvalidateCatalogSnapshot[3])) = v38
							*(*int32)(unsafe.Add(mBase, uint32(v30)+52)) = v38
						}
					}
					return
				}
			}
		}
	}
}
func F_InvalidateLocalBuffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = l0 + int32(36)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v14 != int32(-1) {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v17
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
		*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v19
		F_pgaio_wref_wait(m, v9+int32(24))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v25 = int64(0)
			v28 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v25, v25)
			if l1 != 0 {
				v32 = (int32(-2) - v11) << (uint(int32(2)) % 32)
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateLocalBuffer[0]))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v34)))
				if v36|base.B2i32(v28&int64(262143) != int64(0)) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v69 = v9 + int32(24)
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v74 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateLocalBuffer[1]))
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						F_GetRelationPath(m, v69, v70, v71, v72, v74, v75)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateLocalBuffer[0]))
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v79+v32)))
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v67
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v81
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v69
							F_errmsg_internal(m, int32(_a_F_InvalidateLocalBuffer_0), v9)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_InvalidateLocalBuffer_1), int32(663), int32(_a_F_InvalidateLocalBuffer_2))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
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
					v44 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateLocalBuffer[2]))
					v47 = F_hash_search(m, v44, l0, int32(2), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						if v47 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_InvalidateLocalBuffer_3), int32(0))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_InvalidateLocalBuffer_1), int32(669), int32(_a_F_InvalidateLocalBuffer_2))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-4294967296)
							*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v28 & int64(-17179607041)
							m.G0 = v9 + int32(96)
							return
						}
					}
				}
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateLocalBuffer[2]))
				v47 = F_hash_search(m, v44, l0, int32(2), int32(0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					if v47 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_InvalidateLocalBuffer_3), int32(0))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_InvalidateLocalBuffer_1), int32(669), int32(_a_F_InvalidateLocalBuffer_2))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(-1)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-4294967296)
						*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v28 & int64(-17179607041)
						m.G0 = v9 + int32(96)
						return
					}
				}
			}
		}
	} else {
		v25 = int64(0)
		v28 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v25, v25)
		if l1 != 0 {
			v32 = (int32(-2) - v11) << (uint(int32(2)) % 32)
			v34 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateLocalBuffer[0]))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v34)))
			if v36|base.B2i32(v28&int64(262143) != int64(0)) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v69 = v9 + int32(24)
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v74 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateLocalBuffer[1]))
					v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					F_GetRelationPath(m, v69, v70, v71, v72, v74, v75)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateLocalBuffer[0]))
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v79+v32)))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v67
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v81
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v69
						F_errmsg_internal(m, int32(_a_F_InvalidateLocalBuffer_0), v9)
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_InvalidateLocalBuffer_1), int32(663), int32(_a_F_InvalidateLocalBuffer_2))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
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
				v44 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateLocalBuffer[2]))
				v47 = F_hash_search(m, v44, l0, int32(2), int32(0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					if v47 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_InvalidateLocalBuffer_3), int32(0))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_InvalidateLocalBuffer_1), int32(669), int32(_a_F_InvalidateLocalBuffer_2))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(-1)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-4294967296)
						*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v28 & int64(-17179607041)
						m.G0 = v9 + int32(96)
						return
					}
				}
			}
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateLocalBuffer[2]))
			v47 = F_hash_search(m, v44, l0, int32(2), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				if v47 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_InvalidateLocalBuffer_3), int32(0))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_InvalidateLocalBuffer_1), int32(669), int32(_a_F_InvalidateLocalBuffer_2))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(-1)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-4294967296)
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v28 & int64(-17179607041)
					m.G0 = v9 + int32(96)
					return
				}
			}
		}
	}
}
func F_InvalidateSystemCachesExtended(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateSystemCachesExtended[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v8 = v7
	goto L6
L4:
	;
	goto L5
L5:
	;
	F_RelationCacheInvalidate(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	F_ResetCatalogCache(m, v8-int32(100))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v14 != 0 {
		v8 = v14
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateSystemCachesExtended[1]))
	if int32(0) < v20 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v24 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v40 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateSystemCachesExtended[2]))
	if v40 < v42 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v26 = v24 << (uint(int32(4)) % 32)
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_InvalidateSystemCachesExtended[3])))
	v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_InvalidateSystemCachesExtended[4]))))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_InvalidateSystemCachesExtended[5])))
	m.T0[v30].(func(*base.Module, int64, int32, int32))(m, v27, v28, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	v34 = v24 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateSystemCachesExtended[1]))
	if v34 < v36 {
		v24 = v34
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v46 = v40
	goto L21
L19:
	;
	goto L20
L20:
	;
	v61 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateSystemCachesExtended[6]))
	if v61 < v63 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v48 = v46 << (uint(int32(4)) % 32)
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_InvalidateSystemCachesExtended[7])))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_InvalidateSystemCachesExtended[8])))
	m.T0[v51].(func(*base.Module, int64, int32))(m, v49, int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	goto L20
L23:
	;
	v55 = v46 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateSystemCachesExtended[2]))
	if v55 < v57 {
		v46 = v55
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v67 = v61
	goto L28
L26:
	;
	goto L27
L27:
	;
	return
L28:
	;
	v69 = v67 << (uint(int32(4)) % 32)
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v69)+uint32(_c_F_InvalidateSystemCachesExtended[9])))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)+uint32(_c_F_InvalidateSystemCachesExtended[10])))
	m.T0[v72].(func(*base.Module, int64, int32))(m, v70, int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L27
L30:
	;
	v76 = v67 + int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateSystemCachesExtended[6]))
	if v76 < v78 {
		v67 = v76
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
}
