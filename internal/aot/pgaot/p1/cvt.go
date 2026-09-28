package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cvt_int8_int2(m *base.Module, l0 int64) int64 {
	var v3 int64
	_ = v3
	var v6 int64
	_ = v6
	var v9 int64
	_ = v9
	v3 = int64(-32768)
	if l0 <= v3 {
		v6 = v3
	} else {
		v6 = l0
	}
	if int64(32767) <= v6 {
		v9 = int64(32767)
	} else {
		v9 = v6
	}
	return v9
}
func F_cvt_int8_int4(m *base.Module, l0 int64) int64 {
	var v3 int64
	_ = v3
	var v6 int64
	_ = v6
	var v9 int64
	_ = v9
	v3 = int64(-2147483648)
	if l0 <= v3 {
		v6 = v3
	} else {
		v6 = l0
	}
	if int64(2147483647) <= v6 {
		v9 = int64(2147483647)
	} else {
		v9 = v6
	}
	return v9
}
func F_cvt_name_text(m *base.Module, l0 int64) int64 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_cstring_to_text(m, base.I32_wrap_i64(l0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v3)
	}
}
func F_cvt_text_name(m *base.Module, l0 int64) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v8 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v13 = F_palloc0(m, int32(64))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = int32(1)
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v16 == v15 {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
				if v22 == int32(18) {
					v25 = int32(16)
				} else {
					v25 = int32(0)
				}
				if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v32 = int32(4)
				} else {
					v32 = v25
				}
				v59 = v32
				v60 = v15
				if v59 != 0 {
					v61 = int32(1)
					if v60&v61 != 0 {
						v65 = v61
					} else {
						v65 = int32(4)
					}
					base.MemoryCopy(m, v13, v8+v65, v59)
				} else {
				}
				return base.I64_extend_i32_u(v13)
			} else {
				if v16&int32(1) != 0 {
					v35 = int32(1)
					v44 = v16
					v45 = int32(base.Ui32(v16)>>(uint(v35)%32)) - v35
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v44 = v39
					v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
				}
				if v45 < int32(64) {
					v59 = v45
					v60 = v44
					if v59 != 0 {
						v61 = int32(1)
						if v60&v61 != 0 {
							v65 = v61
						} else {
							v65 = int32(4)
						}
						base.MemoryCopy(m, v13, v8+v65, v59)
					} else {
					}
					return base.I64_extend_i32_u(v13)
				} else {
					v48 = int32(1)
					if v16&v48 != 0 {
						v52 = v48
					} else {
						v52 = int32(4)
					}
					v55 = F_pg_mbcliplen(m, v8+v52, v45, int32(63))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int64(0)
					} else {
						v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
						v59 = v55
						v60 = v57
						if v59 != 0 {
							v61 = int32(1)
							if v60&v61 != 0 {
								v65 = v61
							} else {
								v65 = int32(4)
							}
							base.MemoryCopy(m, v13, v8+v65, v59)
						} else {
						}
						return base.I64_extend_i32_u(v13)
					}
				}
			}
		}
	}
}
