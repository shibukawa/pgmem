package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int2_avg_accum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v77 int64
	_ = v77
	v4 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+28)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 == int32(0) {
		v35 = int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		switch v10 - int32(429) {
		case 0:
			v35 = int32(1)
		case 1:
			v35 = int32(2)
		default:
			v35 = int32(0)
		}
	}
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v35 != 0 {
		v37 = F_pg_detoast_datum(m, v36)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			v43 = v37
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
			if v44 == int32(0) {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				if v47&int32(-4) == int32(160) {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
					v72 = v43 + (v65<<(uint(int32(3))%32)+int32(23))&int32(-8)
					v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
					*(*int64)(unsafe.Add(mBase, uint32(v72))) = v73 + int64(1)
					v77 = *(*int64)(unsafe.Add(mBase, uint32(v72)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v72)+8)) = v77 + base.I64_extend16_s(v4)
					return v43
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(24269), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(476611), int32(6794), int32(273191))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
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
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(24269), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(476611), int32(6794), int32(273191))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		v41 = F_pg_detoast_datum_copy(m, v36)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			v43 = v41
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
			if v44 == int32(0) {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				if v47&int32(-4) == int32(160) {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
					v72 = v43 + (v65<<(uint(int32(3))%32)+int32(23))&int32(-8)
					v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
					*(*int64)(unsafe.Add(mBase, uint32(v72))) = v73 + int64(1)
					v77 = *(*int64)(unsafe.Add(mBase, uint32(v72)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v72)+8)) = v77 + base.I64_extend16_s(v4)
					return v43
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(24269), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(476611), int32(6794), int32(273191))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
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
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(24269), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(476611), int32(6794), int32(273191))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
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
func F_int2_avg_accum_inv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v77 int64
	_ = v77
	v4 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+28)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 == int32(0) {
		v35 = int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		switch v10 - int32(429) {
		case 0:
			v35 = int32(1)
		case 1:
			v35 = int32(2)
		default:
			v35 = int32(0)
		}
	}
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v35 != 0 {
		v37 = F_pg_detoast_datum(m, v36)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			v43 = v37
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
			if v44 == int32(0) {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				if v47&int32(-4) == int32(160) {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
					v72 = v43 + (v65<<(uint(int32(3))%32)+int32(23))&int32(-8)
					v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
					*(*int64)(unsafe.Add(mBase, uint32(v72))) = v73 - int64(1)
					v77 = *(*int64)(unsafe.Add(mBase, uint32(v72)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v72)+8)) = v77 - base.I64_extend16_s(v4)
					return v43
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(24269), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(476611), int32(6881), int32(30388))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
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
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(24269), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(476611), int32(6881), int32(30388))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		v41 = F_pg_detoast_datum_copy(m, v36)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			v43 = v41
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
			if v44 == int32(0) {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				if v47&int32(-4) == int32(160) {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
					v72 = v43 + (v65<<(uint(int32(3))%32)+int32(23))&int32(-8)
					v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
					*(*int64)(unsafe.Add(mBase, uint32(v72))) = v73 - int64(1)
					v77 = *(*int64)(unsafe.Add(mBase, uint32(v72)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v72)+8)) = v77 - base.I64_extend16_s(v4)
					return v43
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(24269), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(476611), int32(6881), int32(30388))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
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
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(24269), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(476611), int32(6881), int32(30388))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
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
func F_int2_mul_cash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v33 int64
	_ = v33
	var v40 int64
	_ = v40
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = int64(63)
	v12 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v19 = int64(32)
	v20 = int64(base.Ui64(v12) >> (uint(v19) % 64))
	v22 = int64(base.Ui64(v9) >> (uint(v19) % 64))
	v25 = int64(4294967295)
	v26 = v12 & v25
	v28 = v9 & v25
	v29 = v26 * v28
	v33 = int64(base.Ui64(v29)>>(uint(v19)%64)) + v26*v22
	v40 = v28*v20 + v33&v25
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v9*(v12>>(uint(v10)%64)) + v9>>(uint(v10)%64)*v12 + v20*v22 + int64(base.Ui64(v33)>>(uint(v19)%64)) + int64(base.Ui64(v40)>>(uint(v19)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v29&v25 | v40<<(uint(v19)%64)
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v51 != v52>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(382745), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(474583), int32(150), int32(527072))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
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
		v74 = F_Int64GetDatum(m, v52)
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return v74
		}
	}
}
