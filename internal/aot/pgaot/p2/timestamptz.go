package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timestamptz_lt_date(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_lt_date[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_lt_date[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_date2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v41 = base.B2i32(v9 < v17) - base.B2i32(v17 < v9)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v41 = base.B2i32(v9 < v17) - base.B2i32(v17 < v9)
				} else {
					if v9 == int64(9223372036854775807) {
						v32 = int32(-1)
					} else {
						v32 = int32(1)
					}
					v41 = v32
				}
			} else {
				if v9 == int64(-9223372036854775807-1) {
					v37 = int32(1)
				} else {
					v37 = int32(-1)
				}
				v41 = v37
			}
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(base.B2i32(int32(0) < v41))
	}
}
func F_timestamptz_pl_interval_at_zone(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	v5 = m.G0
	v7 = v5 - int32(256)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		F_text_to_cstring_buffer(m, v12, v7, int32(256))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = F_DecodeTimezoneNameToTz(m, v7)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				v21 = F_timestamptz_pl_interval_internal(m, v10, v9, v19)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					m.G0 = v7 + int32(256)
					return v21
				}
			}
		}
	}
}
func F_timestamptz_timetz(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v74 int64
	_ = v74
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v11-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		v16 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
		v74 = int64(0)
		m.G0 = v9 - int32(-64)
		return v74
	} else {
		v25 = int32(0)
		v27 = F_timestamp2tm(m, v11, v7+int32(-48), v7+int32(-44), v7+int32(-52), v25, v25)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int64(0)
		} else {
			if v27 != 0 {
				v31 = int64(0)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v33 = F_errsave_start(m, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
				} else {
					if v33 == int32(0) {
						v74 = v31
						m.G0 = v9 - int32(-64)
						return v74
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamptz_timetz_0), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, v32, int32(_a_F_timestamptz_timetz_1), int32(2994), int32(_a_F_timestamptz_timetz_2))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									v74 = v31
									m.G0 = v9 - int32(-64)
									return v74
								}
							}
						}
					}
				}
			} else {
				v50 = F_palloc(m, int32(16))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					v52 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+12)))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v56
					v58 = int32(60)
					*(*int64)(unsafe.Add(mBase, uint32(v50))) = v52 + base.I64_extend_i32_s(v53+(v54+v55*v58)*v58)*int64(1000000)
					v74 = base.I64_extend_i32_u(v50)
					m.G0 = v9 - int32(-64)
					return v74
				}
			}
		}
	}
}
