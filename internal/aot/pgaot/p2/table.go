package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_table_block_parallelscan_initialize(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v3 = int32(0)
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v7
	v10 = F_RelationGetNumberOfBlocksInFork(m, l0, v3)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v10
		v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[136])))
		if v16 != int32(1) {
			v28 = v3
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+118)))
			if v20 == int32(116) {
				v28 = v3
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _consts[137]))
				v26 = base.I32_div_s(v24, int32(4))
				v28 = base.B2i32(base.Ui32(v26) < base.Ui32(v10))
			}
		}
		*(*int64)(unsafe.Add(mBase, uint32(l1)+32)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(-4294967296)
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v28)
		return int32(40)
	}
}
func F_table_block_parallelscan_nextpage(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v35 int32
	_ = v35
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v8 != 0 {
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		v32 = v8
		v33 = v9 + int64(1)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		if base.Ui32(v12) < base.Ui32(int32(2)) {
			v25 = v12
		} else {
			v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
			if base.Ui64(v15) <= base.Ui64(base.I64_extend_i32_u(v16-v12<<(uint(int32(6))%32))) {
				v25 = v12
			} else {
				v23 = int32(base.Ui32(v12) >> (uint(int32(1)) % 32))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v23
				v25 = v23
			}
		}
		v26 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
		*(*int64)(unsafe.Add(mBase, uint32(l2)+32)) = v26 + base.I64_extend_i32_u(v25)
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v32 = v30
		v33 = v26
	}
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v33
	v35 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v32 - int32(1)
	v39 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+20)))
	if base.Ui64(v33) < base.Ui64(v39) {
		v41 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+28)))
		v43 = base.I64_rem_u_s(v33+v41, v39)
		v44 = base.I32_wrap_i64(v43)
		v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
		if v45 != 0 {
			v54 = v44
			v55 = v44
			F_ss_report_location(m, l0, v54)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v61 = v55
				return v61
			}
		} else {
			v61 = v44
			return v61
		}
	} else {
		if v33 != v39 {
			v61 = v35
			return v61
		} else {
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
			if v47&int32(1) == int32(0) {
				v61 = v35
				return v61
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
				v54 = v52
				v55 = int32(-1)
				F_ss_report_location(m, l0, v54)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v61 = v55
					return v61
				}
			}
		}
	}
}
func F_table_tuple_get_latest_tid(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+188))
	v13 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	if v13 != 0 {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[126])))
		if v15&int32(1) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(336321), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_errfinish(m, int32(496975), int32(247), int32(432783))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
			v21 = m.T0[v20].(func(*base.Module, int32, int32) int32)(m, l0, l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				if v21 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
							v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
							v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v54
							*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v53 + int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v51 | v52<<(uint(int32(16))%32)
							F_errmsg(m, int32(704581), v8)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								F_errfinish(m, int32(496975), int32(259), int32(432783))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
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
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
					m.T0[v25].(func(*base.Module, int32, int32))(m, l0, l1)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
		v21 = m.T0[v20].(func(*base.Module, int32, int32) int32)(m, l0, l1)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			if v21 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
						v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
						v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v54
						*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v53 + int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v51 | v52<<(uint(int32(16))%32)
						F_errmsg(m, int32(704581), v8)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							F_errfinish(m, int32(496975), int32(259), int32(432783))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
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
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
				m.T0[v25].(func(*base.Module, int32, int32))(m, l0, l1)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		}
	}
}
