package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SlruInternalDeleteSegment(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	v2 = l1
	v5 = m.G0
	v7 = v5 - int32(1072)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v9 != int32(5) {
		v12 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = v12
		*(*int64)(unsafe.Add(mBase, uint32(v7)+48)) = v12
		*(*int64)(unsafe.Add(mBase, uint32(v7)+64)) = v2
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+48)) = uint16(v9)
		v22 = F_RegisterSyncRequest(m, v7+int32(48), int32(2), int32(1))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v25 = l0 + int32(16)
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
			if v26 == int32(1) {
				*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v2
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v25
				v37 = F_pg_snprintf(m, v7+int32(48), int32(1024), int32(501981), v7+int32(16))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v51 = F_errstart(m, int32(13), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						if v51 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(48)
							F_errmsg_internal(m, int32(693091), v7)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								F_errfinish(m, int32(485007), int32(1518), int32(94086))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									v66 = F_unlink(m, v7+int32(48))
									mBase = m.M
									m.G0 = v7 + int32(1072)
									return
								}
							}
						} else {
							v66 = F_unlink(m, v7+int32(48))
							mBase = m.M
							m.G0 = v7 + int32(1072)
							return
						}
					}
				}
			} else {
				*(*uint32)(unsafe.Add(mBase, uint32(v7)+36)) = uint32(v2)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v25
				v47 = F_pg_snprintf(m, v7+int32(48), int32(1024), int32(502368), v7+int32(32))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					v51 = F_errstart(m, int32(13), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						if v51 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(48)
							F_errmsg_internal(m, int32(693091), v7)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								F_errfinish(m, int32(485007), int32(1518), int32(94086))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									v66 = F_unlink(m, v7+int32(48))
									mBase = m.M
									m.G0 = v7 + int32(1072)
									return
								}
							}
						} else {
							v66 = F_unlink(m, v7+int32(48))
							mBase = m.M
							m.G0 = v7 + int32(1072)
							return
						}
					}
				}
			}
		}
	} else {
		v25 = l0 + int32(16)
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
		if v26 == int32(1) {
			*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v2
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v25
			v37 = F_pg_snprintf(m, v7+int32(48), int32(1024), int32(501981), v7+int32(16))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				v51 = F_errstart(m, int32(13), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					if v51 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(48)
						F_errmsg_internal(m, int32(693091), v7)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							F_errfinish(m, int32(485007), int32(1518), int32(94086))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								v66 = F_unlink(m, v7+int32(48))
								mBase = m.M
								m.G0 = v7 + int32(1072)
								return
							}
						}
					} else {
						v66 = F_unlink(m, v7+int32(48))
						mBase = m.M
						m.G0 = v7 + int32(1072)
						return
					}
				}
			}
		} else {
			*(*uint32)(unsafe.Add(mBase, uint32(v7)+36)) = uint32(v2)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v25
			v47 = F_pg_snprintf(m, v7+int32(48), int32(1024), int32(502368), v7+int32(32))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				v51 = F_errstart(m, int32(13), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					if v51 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(48)
						F_errmsg_internal(m, int32(693091), v7)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							F_errfinish(m, int32(485007), int32(1518), int32(94086))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								v66 = F_unlink(m, v7+int32(48))
								mBase = m.M
								m.G0 = v7 + int32(1072)
								return
							}
						}
					} else {
						v66 = F_unlink(m, v7+int32(48))
						mBase = m.M
						m.G0 = v7 + int32(1072)
						return
					}
				}
			}
		}
	}
}
func F_SlruScanDirCbDeleteCutoff(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	var v22 int32
	_ = v22
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = m.T0[v7].(func(*base.Module, int64, int64) int32)(m, l2, v6)
	mBase = m.M
	if v8 == int32(0) {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v14 = m.T0[v13].(func(*base.Module, int64, int64) int32)(m, l2+int64(31), v6)
		mBase = m.M
		if v14 == int32(0) {
			return int32(0)
		} else {
			v18 = base.I64_div_s(l2, int64(32))
			F_SlruInternalDeleteSegment(m, l0, v18)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	}
}
