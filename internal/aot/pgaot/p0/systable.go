package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_systable_endscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v3 != 0 {
		F_ExecDropSingleTupleTableSlot(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v8 != 0 {
				v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_index_endscan(m, v9)
				mBase = m.M
				v11 = m.ExcPending
				if v11 != 0 {
					return
				} else {
					v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_relation_close(m, v12, int32(1))
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v23 != 0 {
							F_UnregisterSnapshot(m, v23)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								v27 = *(*int32)(unsafe.Add(mBase, _c_F_systable_endscan[0]))
								if v27 != 0 {
									v29 = int32(0)
									*(*uint8)(unsafe.Add(mBase, _c_F_systable_endscan[1])) = uint8(v29)
								} else {
								}
								F_pfree(m, l0)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, _c_F_systable_endscan[0]))
							if v27 != 0 {
								v29 = int32(0)
								*(*uint8)(unsafe.Add(mBase, _c_F_systable_endscan[1])) = uint8(v29)
							} else {
							}
							F_pfree(m, l0)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+188))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
				m.T0[v19].(func(*base.Module, int32))(m, v16)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v23 != 0 {
						F_UnregisterSnapshot(m, v23)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, _c_F_systable_endscan[0]))
							if v27 != 0 {
								v29 = int32(0)
								*(*uint8)(unsafe.Add(mBase, _c_F_systable_endscan[1])) = uint8(v29)
							} else {
							}
							F_pfree(m, l0)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, _c_F_systable_endscan[0]))
						if v27 != 0 {
							v29 = int32(0)
							*(*uint8)(unsafe.Add(mBase, _c_F_systable_endscan[1])) = uint8(v29)
						} else {
						}
						F_pfree(m, l0)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v8 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_index_endscan(m, v9)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_relation_close(m, v12, int32(1))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v23 != 0 {
						F_UnregisterSnapshot(m, v23)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, _c_F_systable_endscan[0]))
							if v27 != 0 {
								v29 = int32(0)
								*(*uint8)(unsafe.Add(mBase, _c_F_systable_endscan[1])) = uint8(v29)
							} else {
							}
							F_pfree(m, l0)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, _c_F_systable_endscan[0]))
						if v27 != 0 {
							v29 = int32(0)
							*(*uint8)(unsafe.Add(mBase, _c_F_systable_endscan[1])) = uint8(v29)
						} else {
						}
						F_pfree(m, l0)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+188))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
			m.T0[v19].(func(*base.Module, int32))(m, v16)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v23 != 0 {
					F_UnregisterSnapshot(m, v23)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, _c_F_systable_endscan[0]))
						if v27 != 0 {
							v29 = int32(0)
							*(*uint8)(unsafe.Add(mBase, _c_F_systable_endscan[1])) = uint8(v29)
						} else {
						}
						F_pfree(m, l0)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, _c_F_systable_endscan[0]))
					if v27 != 0 {
						v29 = int32(0)
						*(*uint8)(unsafe.Add(mBase, _c_F_systable_endscan[1])) = uint8(v29)
					} else {
					}
					F_pfree(m, l0)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
