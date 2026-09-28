package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_be_lo_close(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v9 < int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(67137668))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
				F_errmsg(m, int32(_a_F_be_lo_close_0), v7)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_be_lo_close_1), int32(133), int32(_a_F_be_lo_close_2))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
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
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_close[0]))
		if v13 <= v9 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
					F_errmsg(m, int32(_a_F_be_lo_close_0), v7)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_be_lo_close_1), int32(133), int32(_a_F_be_lo_close_2))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
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
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_close[1]))
			v19 = v16 + v9<<(uint(int32(2))%32)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v20 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
						F_errmsg(m, int32(_a_F_be_lo_close_0), v7)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_be_lo_close_1), int32(133), int32(_a_F_be_lo_close_2))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
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
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0)
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
				if v25 != 0 {
					v27 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_close[2]))
					F_UnregisterSnapshotFromOwner(m, v25, v27)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						F_pfree(m, v20)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							m.G0 = v7 + int32(16)
							return int64(0)
						}
					}
				} else {
					F_pfree(m, v20)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						m.G0 = v7 + int32(16)
						return int64(0)
					}
				}
			}
		}
	}
}
func F_be_lo_lseek64(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v10 < int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(67137668))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
				F_errmsg(m, int32(_a_F_be_lo_lseek64_0), v8)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_be_lo_lseek64_1), int32(241), int32(_a_F_be_lo_lseek64_2))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
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
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_lseek64[0]))
		if v14 <= v10 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
					F_errmsg(m, int32(_a_F_be_lo_lseek64_0), v8)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_be_lo_lseek64_1), int32(241), int32(_a_F_be_lo_lseek64_2))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
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
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_lseek64[1]))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+v10<<(uint(int32(2))%32))))
			if v21 != 0 {
				v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
				v44 = F_inv_seek(m, v21, v41, base.I32_wrap_i64(v42))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int64(0)
				} else {
					m.G0 = v8 + int32(16)
					return v44
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
						F_errmsg(m, int32(_a_F_be_lo_lseek64_0), v8)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_be_lo_lseek64_1), int32(241), int32(_a_F_be_lo_lseek64_2))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
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
