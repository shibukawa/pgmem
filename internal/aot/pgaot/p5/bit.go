package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bit_and(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v15 == v16 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v21 = F_palloc(m, int32(base.Ui32(v18)>>(uint(int32(2))%32)))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v15
					v24 = int32(-4)
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v18 & v24
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					if v27&v24 != int32(32) {
						v32 = int32(8)
						v38 = v8 + v32
						v39 = v13 + v32
						v41 = v21 + v32
						v43 = int32(0)
						for {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
							v46 = v44 & v45
							*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v46)
							v48 = int32(1)
							v55 = v43 + v48
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							if base.Ui32(v55) < base.Ui32(int32(base.Ui32(v56)>>(uint(int32(2))%32))-int32(8)) {
								v38 = v38 + v48
								v39 = v39 + v48
								v41 = v41 + v48
								v43 = v55
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					return v21
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(101187714))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(157628), int32(0))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(494687), int32(1261), int32(430542))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
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
}
func F_bit_or(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v15 == v16 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v21 = F_palloc(m, int32(base.Ui32(v18)>>(uint(int32(2))%32)))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v15
					v24 = int32(-4)
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v18 & v24
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					if v27&v24 != int32(32) {
						v32 = int32(8)
						v38 = v8 + v32
						v39 = v13 + v32
						v41 = v21 + v32
						v43 = int32(0)
						for {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
							v46 = v44 | v45
							*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v46)
							v48 = int32(1)
							v55 = v43 + v48
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							if base.Ui32(v55) < base.Ui32(int32(base.Ui32(v56)>>(uint(int32(2))%32))-int32(8)) {
								v38 = v38 + v48
								v39 = v39 + v48
								v41 = v41 + v48
								v43 = v55
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					return v21
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(101187714))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(157587), int32(0))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(494687), int32(1302), int32(213476))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
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
}
