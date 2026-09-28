package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_poly_center(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v16 = F_palloc(m, int32(16))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			F_poly_to_circle(m, v8+int32(8), v11, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v23 == int32(0) {
					v35 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v35
					v37 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v16))) = v37
					v40 = base.I64_extend_i32_u(v16)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
					if v26 != int32(453) {
						v35 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v35
						v37 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v16))) = v37
						v40 = base.I64_extend_i32_u(v16)
					} else {
						v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
						if v29 != int32(1) {
							v35 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v35
							v37 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v16))) = v37
							v40 = base.I64_extend_i32_u(v16)
						} else {
							v32 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
							v40 = int64(0)
						}
					}
				}
				m.G0 = v8 + int32(32)
				return v40
			}
		}
	}
}
func F_poly_overlap(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = F_poly_overlap_internal(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int64(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v13)
							}
						} else {
							return base.I64_extend_i32_u(v13)
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v13)
						}
					} else {
						return base.I64_extend_i32_u(v13)
					}
				}
			}
		}
	}
}
func F_poly_path(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v48 float64
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v17 = v13<<(uint(int32(4))%32) + int32(16)
		v18 = F_palloc(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = v17 << (uint(int32(2)) % 32)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v23
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			if int32(0) < v27 {
				v35 = int32(0)
				for {
					v43 = v35 << (uint(int32(4)) % 32)
					v44 = v18 + int32(16) + v43
					v45 = v43 + (v9 + int32(40))
					v46 = *(*float64)(unsafe.Add(mBase, uint32(v45)))
					*(*float64)(unsafe.Add(mBase, uint32(v44))) = v46
					v48 = *(*float64)(unsafe.Add(mBase, uint32(v45)+8))
					*(*float64)(unsafe.Add(mBase, uint32(v44)+8)) = v48
					v51 = v35 + int32(1)
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					if v51 < v52 {
						v35 = v51
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			return base.I64_extend_i32_u(v18)
		}
	}
}
