package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_poly_center(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = F_palloc(m, int32(16))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_poly_to_circle(m, v6+int32(8), v9)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = *(*int64)(unsafe.Add(mBase, uint32(v6)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v20
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v14))) = v22
				m.G0 = v6 + int32(32)
				return v14
			}
		}
	}
}
func F_poly_overlap(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = F_poly_overlap_internal(m, v7, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v18 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						if v22 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								return v16
							}
						} else {
							return v16
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v22 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return v16
						}
					} else {
						return v16
					}
				}
			}
		}
	}
}
func F_poly_path(m *base.Module, l0 int32) int32 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v17 = v13<<(uint(int32(4))%32) + int32(16)
		v18 = F_palloc(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
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
			return v18
		}
	}
}
