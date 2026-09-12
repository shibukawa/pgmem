package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WinRowsArePeers(m *base.Module, l0 int32, l1 int64, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
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
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	if v15 == int32(0) {
		v73 = int32(1)
		m.G0 = v11 + int32(32)
		return v73
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+404))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+400))
		v21 = F_window_gettupleslot(m, l0, l1, v20)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			if v21 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = l1
					F_errmsg_internal(m, int32(446989), v11+int32(16))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(518338), int32(3344), int32(142896))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v27 = F_window_gettupleslot(m, l0, l2, v19)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v11))) = l2
							F_errmsg_internal(m, int32(446989), v11)
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(518338), int32(3347), int32(142896))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+96))
						if v32 == int32(0) {
							v62 = int32(1)
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
							m.T0[v65].(func(*base.Module, int32))(m, v20)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
								m.T0[v69].(func(*base.Module, int32))(m, v19)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									v73 = v62
									m.G0 = v11 + int32(32)
									return v73
								}
							}
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v13)+372))
							*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v19
							*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v20
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)+140))
							if v39 != 0 {
								v40 = int32(4548768)
								v41 = *(*int32)(unsafe.Add(mBase, _consts[28]))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
								*(*int32)(unsafe.Add(mBase, _consts[28])) = v43
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
								v48 = m.T0[v47].(func(*base.Module, int32, int32, int32) int32)(m, v39, v36, v11+int32(31))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[28])) = v41
									v57 = base.B2i32(v48 != int32(0))
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
									F_MemoryContextReset(m, v58)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v62 = v57
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
										m.T0[v65].(func(*base.Module, int32))(m, v20)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
											m.T0[v69].(func(*base.Module, int32))(m, v19)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												v73 = v62
												m.G0 = v11 + int32(32)
												return v73
											}
										}
									}
								}
							} else {
								v57 = int32(1)
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
								F_MemoryContextReset(m, v58)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v62 = v57
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
									m.T0[v65].(func(*base.Module, int32))(m, v20)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
										m.T0[v69].(func(*base.Module, int32))(m, v19)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											v73 = v62
											m.G0 = v11 + int32(32)
											return v73
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
func F_win_to_utf8(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v15, v16, v17, int32(-1), int32(6))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		v25 = v15 - int32(18)
		if base.Ui32(int32(16)) <= base.Ui32(v25) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(2600))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v15
					F_errmsg(m, int32(131246), v10)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(516136), int32(114), int32(576708))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			if int32(base.Ui32(int32(63599))>>(uint(v25)%32))&int32(1) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(2600))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v15
						F_errmsg(m, int32(131246), v10)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(516136), int32(114), int32(576708))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v25<<(uint(int32(2))%32))+uint32(_consts[1489])))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				v40 = int32(0)
				v45 = F_LocalToUtf(m, v14, v17, v13, v39, v40, v40, v40, v15, base.B2i32(v12 != v40))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					m.G0 = v10 + int32(16)
					return v45
				}
			}
		}
	}
}
