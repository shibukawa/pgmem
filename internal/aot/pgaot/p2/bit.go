package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bit_overlay(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	if int32(0) < l2 {
		v9 = l2 + l3
		if base.B2i32(l3 < int32(0)) != base.B2i32(v9 < l2) {
			F_errstart_cold(m, int32(21), int32(0))
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_bit_overlay_0), int32(0))
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_bit_overlay_1), int32(1195), int32(_a_F_bit_overlay_2))
						v60 = m.ExcPending
						if v60 != 0 {
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
			v12 = int32(1)
			v16 = F_bitsubstring(m, l0, v12, l2-v12, int32(0))
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v22 = F_bitsubstring(m, l0, v9, int32(-1), int32(1))
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = F_bit_catenate(m, v16, l1)
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = F_bit_catenate(m, v24, v22)
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							return v26
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(17039490))
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_bit_overlay_3), int32(0))
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_bit_overlay_1), int32(1191), int32(_a_F_bit_overlay_2))
					v44 = m.ExcPending
					if v44 != 0 {
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
func F_bit_recv(m *base.Module, l0 int32) int64 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pq_getmsgint(m, v12, int32(4))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		if base.Ui32(v14) < base.Ui32(int32(2147483641)) {
			if base.B2i32(v14 != v11)&base.B2i32(int32(0) < v11) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(101187714))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v11
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
						F_errmsg(m, int32(_a_F_bit_recv_0), v9)
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_bit_recv_1), int32(357), int32(_a_F_bit_recv_2))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
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
				v27 = int32(base.Ui32(v14+int32(7)) >> (uint(int32(3)) % 32))
				v29 = v27 + int32(8)
				v30 = F_palloc(m, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v14
					*(*int32)(unsafe.Add(mBase, uint32(v30))) = v29 << (uint(int32(2)) % 32)
					F_pq_copymsgbytes(m, v12, v30+int32(8), v27)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
						v42 = int32(base.Ui32(v40) >> (uint(int32(2)) % 32))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
						v48 = v42<<(uint(int32(3))%32) - v45 + int32(-64)
						if int32(0) < v48 {
							v53 = v30 + v42 - int32(1)
							v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
							v57 = v54 & (int32(255) << (uint(v48) % 32))
							*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v57)
						} else {
						}
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v30)
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_bit_recv_3), int32(0))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_bit_recv_1), int32(347), int32(_a_F_bit_recv_2))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
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
