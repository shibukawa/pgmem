package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bbsink_copystream_begin_manifest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_pq_beginmessage(m, v5, int32(100))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		F_enlargeStringInfo(m, v5, int32(1))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v16 = int32(109)
			*(*uint8)(unsafe.Add(mBase, uint32(v13+v14))) = uint8(v16)
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v13 + int32(1)
			F_pq_endmessage(m, v5)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				m.G0 = v5 + int32(16)
				return
			}
		}
	}
}
func F_bbsink_copystream_end_manifest(m *base.Module, l0 int32) {
	return
}
func F_bbsink_forward_end_manifest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
	m.T0[v4].(func(*base.Module, int32))(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_bbsink_server_archive_contents(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v13
	v20 = F_FileWriteV(m, v12, v9+int32(40), int32(1), v11, int32(167772165))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		if l1 != v20 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				if v20 < int32(0) {
					F_errcode_for_file_access(m)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v71 = *(*int32)(unsafe.Add(mBase, _consts[424]))
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v69*int32(48))+32))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v75
						F_errmsg(m, int32(311063), v9)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return
						} else {
							F_errhint(m, int32(667856), int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_errfinish(m, int32(515146), int32(175), int32(127578))
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
					F_errcode(m, int32(4293))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v34 = *(*int32)(unsafe.Add(mBase, _consts[424]))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v32*int32(48))+32))
						v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
						*(*uint32)(unsafe.Add(mBase, uint32(v9)+28)) = uint32(v39)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v20
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v38
						F_errmsg(m, int32(43758), v9+int32(16))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							F_errhint(m, int32(667856), int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								F_errfinish(m, int32(515146), int32(182), int32(127578))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			}
		} else {
			v58 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v58 + base.I64_extend_i32_s(l1)
			F_bbsink_forward_archive_contents(m, l0, l1)
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				m.G0 = v9 + int32(48)
				return
			}
		}
	}
}
func F_bbsink_server_begin_manifest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v9
	v14 = F_psprintf(m, int32(245374), v7+int32(16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v17 = F_PathNameOpenFile(m, v14, int32(193))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v17
			if v17 <= int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v14
						F_errmsg(m, int32(311157), v7)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_errfinish(m, int32(515146), int32(242), int32(82944))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
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
				F_pfree(m, v14)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_bbsink_forward_begin_manifest(m, l0)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						m.G0 = v7 + int32(32)
						return
					}
				}
			}
		}
	}
}
func F_bbsink_throttle_manifest_contents(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	F_throttle(m, l0, l1)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_bbsink_forward_manifest_contents(m, l0, l1)
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
