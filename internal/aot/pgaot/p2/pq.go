package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pq_endmsgread(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_pq_endmsgread[0])) = uint8(v2)
	return
}
func F_pq_getmsgint(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	switch l1 - int32(1) {
	case 0:
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v11-v12 <= int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16908800))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pq_getmsgint_0), int32(0))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pq_getmsgint_1), int32(533), int32(_a_F_pq_getmsgint_2))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
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
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v12))))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12 + int32(1)
			v60 = v18
			m.G0 = v7 + int32(16)
			return v60
		}
	case 1:
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v22-v23 <= int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16908800))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pq_getmsgint_0), int32(0))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pq_getmsgint_1), int32(533), int32(_a_F_pq_getmsgint_2))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
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
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+v23))))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v23 + int32(2)
			v33 = int32(8)
			v60 = (v29<<(uint(v33)%32) | int32(base.Ui32(v29)>>(uint(v33)%32))) & int32(_a_F_pq_getmsgint_3)
			m.G0 = v7 + int32(16)
			return v60
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
			F_errmsg_internal(m, int32(_a_F_pq_getmsgint_4), v7)
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_pq_getmsgint_1), int32(437), int32(_a_F_pq_getmsgint_5))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 3:
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v40-v41 <= int32(3) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16908800))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pq_getmsgint_0), int32(0))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pq_getmsgint_1), int32(533), int32(_a_F_pq_getmsgint_2))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
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
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v45+v41)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v41 + int32(4)
			v53 = int32(16711935)
			v60 = base.I32_rotr(v47, int32(24))&v53 | base.I32_rotr(v47&v53, int32(8))
			m.G0 = v7 + int32(16)
			return v60
		}
	}
}
func F_pq_getmsgint64(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4-v5 <= int32(7) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_pq_getmsgint64_0), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pq_getmsgint64_1), int32(533), int32(_a_F_pq_getmsgint64_2))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v29 = *(*int64)(unsafe.Add(mBase, uint32(v27+v5)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5 + int32(8)
		v33 = int64(56)
		v35 = int64(65280)
		v37 = int64(40)
		v40 = int64(16711680)
		v42 = int64(24)
		v44 = int64(4278190080)
		v46 = int64(8)
		return v29<<(uint(v33)%64) | v29&v35<<(uint(v37)%64) | (v29&v40<<(uint(v42)%64) | v29&v44<<(uint(v46)%64)) | (int64(base.Ui64(v29)>>(uint(v46)%64))&v44 | int64(base.Ui64(v29)>>(uint(v42)%64))&v40 | (int64(base.Ui64(v29)>>(uint(v37)%64))&v35 | int64(base.Ui64(v29)>>(uint(v33)%64))))
	}
}
func F_pq_setkeepalivescount(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	if l1 == int32(0) {
	} else {
		v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
		if v11 == int32(1) {
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+408))
			if l0 == v14 {
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+392))
				if int32(0) < v16 {
					if l0 == int32(0) {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+392))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v46
					} else {
					}
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+408)) = v49
				} else {
					v19 = int32(0)
					v21 = m.G0
					v23 = v21 - int32(16)
					m.G0 = v23
					if l1 == v19 {
						v38 = v19
					} else {
						v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
						if v27 == int32(1) {
							v38 = v19
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+408))
							if v30 != 0 {
								v38 = v30
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+392))
								if v31 != 0 {
									v38 = v31
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(4)
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(392))))
									v38 = v37
								}
							}
						}
					}
					m.G0 = v23 + int32(16)
					if int32(0) <= v38 {
						if l0 == int32(0) {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+392))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v46
						} else {
						}
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+408)) = v49
					} else {
					}
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return
}
