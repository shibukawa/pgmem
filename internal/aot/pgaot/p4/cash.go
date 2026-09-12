package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cash_dist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = v9 - v5
	v13 = int32(0)
	if base.B2i32(base.B2i32(int64(0) < v5)^base.B2i32(v10 < v9) == v13)&base.B2i32(v10 != int64(-9223372036854775807-1)) == v13 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(401453), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(497306), int32(107), int32(75722))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
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
		v43 = v10 >> (uint(int64(63)) % 64)
		v46 = F_Int64GetDatum(m, v10^v43-v43)
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int32(0)
		} else {
			return v46
		}
	}
}
func F_cash_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	return base.B2i32(v3 == v5)
}
func F_cash_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	return base.B2i32(v5 <= v3)
}
func F_cash_mul_int4(m *base.Module, l0 int32) int32 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = int64(63)
	v12 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
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
				F_errmsg(m, int32(401453), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(497312), int32(150), int32(551223))
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
func F_cash_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pq_getmsgint64(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_Int64GetDatum(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_cash_send(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	F_pq_begintypsend(m, v6)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		F_enlargeStringInfo(m, v6, int32(8))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v20 = int64(56)
			v22 = int64(65280)
			v24 = int64(40)
			v27 = int64(16711680)
			v29 = int64(24)
			v31 = int64(4278190080)
			v33 = int64(8)
			*(*int64)(unsafe.Add(mBase, uint32(v17+v18))) = v9<<(uint(v20)%64) | v9&v22<<(uint(v24)%64) | (v9&v27<<(uint(v29)%64) | v9&v31<<(uint(v33)%64)) | (int64(base.Ui64(v9)>>(uint(v33)%64))&v31 | int64(base.Ui64(v9)>>(uint(v29)%64))&v27 | (int64(base.Ui64(v9)>>(uint(v24)%64))&v22 | int64(base.Ui64(v9)>>(uint(v20)%64))))
			v57 = v17 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v57
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*int32)(unsafe.Add(mBase, uint32(v60))) = v57 << (uint(int32(2)) % 32)
			m.G0 = v6 + int32(16)
			return v60
		}
	}
}
