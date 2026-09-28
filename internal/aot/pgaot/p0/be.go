package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_be_lo_create(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_PreventCommandIfReadOnly(m, int32(_a_F_be_lo_create_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v9 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_be_lo_create[0])) = uint8(v9)
		v11 = F_inv_create(m, v2)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v11)
		}
	}
}
func F_be_lo_from_bytea(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		F_PreventCommandIfReadOnly(m, int32(_a_F_be_lo_from_bytea_0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v16 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_be_lo_from_bytea[0])) = uint8(v16)
			v18 = F_inv_create(m, v6)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_from_bytea[1]))
				v23 = F_inv_open(m, v18, int32(_a_F_be_lo_from_bytea_1), v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
					if v25 == int32(1) {
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
						if v31 == int32(18) {
							v34 = int32(16)
						} else {
							v34 = int32(0)
						}
						if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v41 = int32(4)
						} else {
							v41 = v34
						}
						v54 = v41
					} else {
						v42 = int32(1)
						if v25&v42 != 0 {
							v54 = int32(base.Ui32(v25)>>(uint(v42)%32)) - v42
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v55 = int32(1)
					if v25&v55 != 0 {
						v59 = v55
					} else {
						v59 = int32(4)
					}
					v61 = F_inv_write(m, v23, v8+v59, v54)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int64(0)
					} else {
						F_pfree(m, v23)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v18)
						}
					}
				}
			}
		}
	}
}
func F_be_lo_get(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_lo_get_fragment_internal(m, v2, int64(0), int32(-1))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v5)
	}
}
func F_be_lo_lseek(m *base.Module, l0 int32) int64 {
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
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v10 < int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(67137668))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
				F_errmsg(m, int32(_a_F_be_lo_lseek_0), v8)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_be_lo_lseek_1), int32(216), int32(_a_F_be_lo_lseek_2))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
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
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_lseek[0]))
		if v14 <= v10 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
					F_errmsg(m, int32(_a_F_be_lo_lseek_0), v8)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_be_lo_lseek_1), int32(216), int32(_a_F_be_lo_lseek_2))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
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
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_lseek[1]))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+v10<<(uint(int32(2))%32))))
			if v21 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
						F_errmsg(m, int32(_a_F_be_lo_lseek_0), v8)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_be_lo_lseek_1), int32(216), int32(_a_F_be_lo_lseek_2))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
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
				v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v27 = F_inv_seek(m, v21, base.I64_extend32_s(v24), v26)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					if base.Ui64(int64(4294967296)) <= base.Ui64(v27+int64(2147483648)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
								F_errmsg(m, int32(_a_F_be_lo_lseek_3), v8+int32(16))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_be_lo_lseek_1), int32(225), int32(_a_F_be_lo_lseek_2))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
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
						m.G0 = v8 + int32(32)
						return v27
					}
				}
			}
		}
	}
}
func F_be_lo_truncate(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v3 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_PreventCommandIfReadOnly(m, int32(_a_F_be_lo_truncate_0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		F_lo_truncate_internal(m, v4, v3)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			return int64(0)
		}
	}
}
func F_be_lo_truncate64(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_PreventCommandIfReadOnly(m, int32(_a_F_be_lo_truncate64_0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		F_lo_truncate_internal(m, v4, v3)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			return int64(0)
		}
	}
}
