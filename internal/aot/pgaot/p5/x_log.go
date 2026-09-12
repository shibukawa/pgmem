package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_XLogCheckBufferNeedsBackup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v48 int32
	_ = v48
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v13 = *(*int64)(unsafe.Add(mBase, _consts[115]))
	*(*int64)(unsafe.Add(mBase, uint32(v6+int32(8)))) = v13
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
	*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(7)))) = uint8(v16)
	if l0 < int32(0) {
		v21 = *(*int32)(unsafe.Add(mBase, _consts[11]))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(l0^int32(-1))<<(uint(int32(2))%32))))
		v35 = v27
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, _consts[12]))
		v35 = v29 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	v36 = int32(1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)))
	if v37 == v36 {
		v40 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
		v41 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35)+4)))
		v42 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35))))
		if base.Ui64(v41|v42<<(uint(int64(32))%64)) <= base.Ui64(v40) {
			v48 = v36
		} else {
			v48 = int32(0)
		}
	} else {
		v48 = int32(0)
	}
	m.G0 = v6 + int32(16)
	return v48
}
func F_XLogCheckpointNeeded(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v8 = *(*int64)(unsafe.Add(mBase, _consts[115]))
	v10 = int64(*(*int32)(unsafe.Add(mBase, _consts[116])))
	v11 = base.I64_div_u_s(v8, v10)
	return base.B2i32(base.Ui64(base.I64_extend_i32_s(v3-int32(1))+v11) <= base.Ui64(l0))
}
func F_XLogRecGetBlockData(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v4 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+72))
	if v6 < l1 {
		v29 = v4
		return v29
	} else {
		v12 = v5 + l1*int32(52) + int32(76)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		if v13 != int32(1) {
			v29 = v4
			return v29
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+43)))
			if v16 == int32(0) {
				if l2 == int32(0) {
					v29 = v4
					return v29
				} else {
					v21 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v21
					return v21
				}
			} else {
				if l2 != 0 {
					v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+48)))
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v25
				} else {
				}
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
				v29 = v27
				return v29
			}
		}
	}
}
func F_XLogRecGetBlockTag(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	if v12 < l1 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
			F_errmsg_internal(m, int32(421180), v9)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				F_errfinish(m, int32(495276), int32(2000), int32(337684))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v18 = v11 + l1*int32(52) + int32(76)
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
		if v19 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
				F_errmsg_internal(m, int32(421180), v9)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					F_errfinish(m, int32(495276), int32(2000), int32(337684))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			if l2 != 0 {
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v18)+4))
				*(*int64)(unsafe.Add(mBase, uint32(l2))) = v22
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v24
			} else {
			}
			if l3 != 0 {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v26
			} else {
			}
			if l4 != 0 {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = v28
			} else {
			}
			m.G0 = v9 + int32(16)
			return
		}
	}
}
func F_XLogRegisterData(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	v12 = *(*int32)(unsafe.Add(mBase, _consts[151]))
	if v12 <= v10 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(504878), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, _consts[150]))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v23
				v26 = *(*int32)(unsafe.Add(mBase, _consts[151]))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v26
				F_errdetail_internal(m, int32(631827), v7)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errfinish(m, int32(492746), int32(374), int32(504979))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
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
		*(*int32)(unsafe.Add(mBase, _consts[150])) = v10 + int32(1)
		v41 = *(*int32)(unsafe.Add(mBase, _consts[152]))
		v44 = v41 + v10*int32(12)
		*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = l0
		v47 = int32(4121084)
		v48 = *(*int32)(unsafe.Add(mBase, _consts[153]))
		*(*int32)(unsafe.Add(mBase, uint32(v48))) = v44
		*(*int32)(unsafe.Add(mBase, _consts[153])) = v44
		v52 = int32(4410952)
		v54 = *(*int64)(unsafe.Add(mBase, _consts[154]))
		*(*int64)(unsafe.Add(mBase, _consts[154])) = v54 + base.I64_extend_i32_u(l1)
		m.G0 = v7 + int32(16)
		return
	}
}
func F_XLogWalRcvClose(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	F_XLogWalRcvFlush(m, int32(0), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[567]))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v14
		v17 = *(*int64)(unsafe.Add(mBase, _consts[568]))
		v20 = int64(*(*int32)(unsafe.Add(mBase, _consts[116])))
		v21 = base.I64_div_u_s(int64(4294967296), v20)
		v22 = base.I64_div_u_s(v17, v21)
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+20)) = uint32(v22)
		v25 = v17 - v21*v22
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+24)) = uint32(v25)
		v33 = F_pg_snprintf(m, v8+int32(32), int32(64), int32(509810), v8+int32(16))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, _consts[569]))
			v37 = F_close(m, v36)
			mBase = m.M
			if v37 == int32(0) {
				v41 = *(*int32)(unsafe.Add(mBase, _consts[570]))
				if v41 != int32(2) {
					F_XLogArchiveForceDone(m, v8+int32(32))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[569])) = int32(-1)
						m.G0 = v8 + int32(96)
						return
					}
				} else {
					F_XLogArchiveNotify(m, v8+int32(32))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[569])) = int32(-1)
						m.G0 = v8 + int32(96)
						return
					}
				}
			} else {
				F_errstart_cold(m, int32(23), int32(0))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(32)
						F_errmsg(m, int32(293428), v8)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							F_errfinish(m, int32(494918), int32(1064), int32(360930))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
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
	}
}
