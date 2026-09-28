package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_be_lo_import_with_oid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_lo_import_internal(m, v3, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v8)
		}
	}
}
func F_be_lo_put(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		F_PreventCommandIfReadOnly(m, int32(_a_F_be_lo_put_0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_be_lo_put[0])) = uint8(v17)
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_put[1]))
			v22 = F_inv_open(m, v7, int32(_a_F_be_lo_put_1), v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v25 = F_inv_seek(m, v22, v6, int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
					if v27 == int32(1) {
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
						if v33 == int32(18) {
							v36 = int32(16)
						} else {
							v36 = int32(0)
						}
						if base.Ui32((v33-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v43 = int32(4)
						} else {
							v43 = v36
						}
						v56 = v43
					} else {
						v44 = int32(1)
						if v27&v44 != 0 {
							v56 = int32(base.Ui32(v27)>>(uint(v44)%32)) - v44
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
							v56 = int32(base.Ui32(v50)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v57 = int32(1)
					if v27&v57 != 0 {
						v61 = v57
					} else {
						v61 = int32(4)
					}
					v63 = F_inv_write(m, v22, v9+v61, v56)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int64(0)
					} else {
						F_pfree(m, v22)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int64(0)
						} else {
							return int64(0)
						}
					}
				}
			}
		}
	}
}
func F_be_lo_tell64(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v9 < int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(67137668))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
				F_errmsg(m, int32(_a_F_be_lo_tell64_0), v7)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_be_lo_tell64_1), int32(306), int32(_a_F_be_lo_tell64_2))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
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
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_tell64[0]))
		if v13 <= v9 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
					F_errmsg(m, int32(_a_F_be_lo_tell64_0), v7)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_be_lo_tell64_1), int32(306), int32(_a_F_be_lo_tell64_2))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
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
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_be_lo_tell64[1]))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+v9<<(uint(int32(2))%32))))
			if v20 != 0 {
				v40 = *(*int64)(unsafe.Add(mBase, uint32(v20)+16))
				m.G0 = v7 + int32(16)
				return v40
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
						F_errmsg(m, int32(_a_F_be_lo_tell64_0), v7)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_be_lo_tell64_1), int32(306), int32(_a_F_be_lo_tell64_2))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
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
