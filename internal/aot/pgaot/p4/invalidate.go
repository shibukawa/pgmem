package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InvalidateCatalogSnapshot(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	v3 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	if v3 == int32(0) {
		return
	} else {
		F_pairingheap_remove(m, int32(4178632), v3+int32(52))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[1458])) = int32(0)
			v15 = *(*int32)(unsafe.Add(mBase, _consts[263]))
			if v15 != 0 {
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _consts[264]))
				if v18 != 0 {
					v20 = *(*int32)(unsafe.Add(mBase, _consts[120]))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
					v23 = *(*int32)(unsafe.Add(mBase, _consts[264]))
					v25 = v23 - int32(48)
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v26))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v21)) == int32(0) {
						v38 = base.B2i32(base.Ui32(v21) < base.Ui32(v26))
					} else {
						v38 = int32(base.Ui32(v21-v26) >> (uint(int32(31)) % 32))
					}
					if v38 == int32(0) {
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
						v42 = v41
						*(*int32)(unsafe.Add(mBase, _consts[83])) = v42
						v46 = *(*int32)(unsafe.Add(mBase, _consts[120]))
						*(*int32)(unsafe.Add(mBase, uint32(v46)+40)) = v42
					}
				} else {
					v42 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[83])) = v42
					v46 = *(*int32)(unsafe.Add(mBase, _consts[120]))
					*(*int32)(unsafe.Add(mBase, uint32(v46)+40)) = v42
				}
			}
			return
		}
	}
}
func F_InvalidateLocalBuffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v18 int64
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = l0 + int32(36)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v13 != int32(-1) {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+104)) = v16
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+96)) = v18
		F_pgaio_wref_wait(m, v8+int32(96))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if l1 != 0 {
				v28 = (int32(-2) - v10) << (uint(int32(2)) % 32)
				v30 = *(*int32)(unsafe.Add(mBase, _consts[942]))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v30)))
				if v32|v24&int32(262143) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v68 = *(*int32)(unsafe.Add(mBase, _consts[166]))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						F_GetRelationPath(m, v8+int32(24), v64, v65, v66, v68, v69)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, _consts[942]))
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v73+v28)))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
							*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v75
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v8 + int32(24)
							F_errmsg_internal(m, int32(682181), v8)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_errfinish(m, int32(499463), int32(637), int32(227101))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
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
					v38 = *(*int32)(unsafe.Add(mBase, _consts[945]))
					v41 = F_hash_search(m, v38, l0, int32(2), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						if v41 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(446093), int32(0))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return
								} else {
									F_errfinish(m, int32(499463), int32(643), int32(227101))
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24 & int32(262143)
							m.G0 = v8 + int32(112)
							return
						}
					}
				}
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, _consts[945]))
				v41 = F_hash_search(m, v38, l0, int32(2), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					if v41 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(446093), int32(0))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return
							} else {
								F_errfinish(m, int32(499463), int32(643), int32(227101))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
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
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24 & int32(262143)
						m.G0 = v8 + int32(112)
						return
					}
				}
			}
		}
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if l1 != 0 {
			v28 = (int32(-2) - v10) << (uint(int32(2)) % 32)
			v30 = *(*int32)(unsafe.Add(mBase, _consts[942]))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v30)))
			if v32|v24&int32(262143) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v68 = *(*int32)(unsafe.Add(mBase, _consts[166]))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					F_GetRelationPath(m, v8+int32(24), v64, v65, v66, v68, v69)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						v73 = *(*int32)(unsafe.Add(mBase, _consts[942]))
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v73+v28)))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
						*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v75
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v8 + int32(24)
						F_errmsg_internal(m, int32(682181), v8)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							F_errfinish(m, int32(499463), int32(637), int32(227101))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
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
				v38 = *(*int32)(unsafe.Add(mBase, _consts[945]))
				v41 = F_hash_search(m, v38, l0, int32(2), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					if v41 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(446093), int32(0))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return
							} else {
								F_errfinish(m, int32(499463), int32(643), int32(227101))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
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
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24 & int32(262143)
						m.G0 = v8 + int32(112)
						return
					}
				}
			}
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, _consts[945]))
			v41 = F_hash_search(m, v38, l0, int32(2), int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				if v41 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(446093), int32(0))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return
						} else {
							F_errfinish(m, int32(499463), int32(643), int32(227101))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v24 & int32(262143)
					m.G0 = v8 + int32(112)
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[1364]))
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
	v20 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
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
	v46 = int32(0)
	v48 = *(*int32)(unsafe.Add(mBase, _consts[1369]))
	if v46 < v48 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v26 = v24 * int32(12)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1370])))
	v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1371]))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1372])))
	m.T0[v36].(func(*base.Module, int32, int32, int32))(m, v29, v32, int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	v40 = v24 + int32(1)
	v42 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	if v40 < v42 {
		v24 = v40
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v52 = v46
	goto L21
L19:
	;
	goto L20
L20:
	;
	v71 = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, _consts[1373]))
	if v71 < v73 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v54 = v52 << (uint(int32(3)) % 32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)+uint32(_consts[1374])))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v54)+uint32(_consts[1375])))
	m.T0[v61].(func(*base.Module, int32, int32))(m, v57, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	goto L20
L23:
	;
	v65 = v52 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, _consts[1369]))
	if v65 < v67 {
		v52 = v65
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v77 = v71
	goto L28
L26:
	;
	goto L27
L27:
	;
	return
L28:
	;
	v79 = v77 << (uint(int32(3)) % 32)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)+uint32(_consts[1376])))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)+uint32(_consts[1377])))
	m.T0[v86].(func(*base.Module, int32, int32))(m, v82, int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L27
L30:
	;
	v90 = v77 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, _consts[1373]))
	if v90 < v92 {
		v77 = v90
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
}
