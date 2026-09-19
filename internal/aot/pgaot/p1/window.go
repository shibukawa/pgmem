package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_window_first_value(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14019(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_window_ntile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v45 int64
	_ = v45
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v62 int64
	_ = v62
	var v65 int64
	_ = v65
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = F_WinGetPartitionLocalMemory(m, v14, int32(32))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		if v20 != 0 {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
			v56 = v20
			v58 = v21
			v62 = v22 + int64(1)
			*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v62
			if v62 <= v58 {
				v96 = v56
			} else {
				v65 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
				if v65 == base.I64_extend_i32_s(v56) {
					*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v58 - int64(1)
				} else {
				}
				*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(1)
				v76 = v56 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v76
				v96 = v76
			}
			m.G0 = v12 + int32(16)
			return v96
		} else {
			v25 = F_WinGetPartitionRowCount(m, v14)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = int32(0)
				v31 = F_WinGetFuncArgCurrent(m, v14, v27, v12+int32(15))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
					if v33 != 0 {
						v94 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v94)
						v96 = v27
						m.G0 = v12 + int32(16)
						return v96
					} else {
						if v31 <= int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67371138))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_window_ntile_0), int32(0))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_window_ntile_1), int32(443), int32(_a_F_window_ntile_2))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(1)
							v38 = base.I64_extend_i32_u(v31)
							v39 = base.I64_div_s(v25, v38)
							*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v39
							if int64(0) < v39 {
								v45 = v25 - v39*v38
								*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v45
								if v45 == int64(0) {
									v53 = v39
								} else {
									v51 = v39 + int64(1)
									*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v51
									v53 = v51
								}
							} else {
								v51 = int64(1)
								*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v51
								v53 = v51
							}
							v56 = int32(1)
							v58 = v53
							v62 = int64(1)
							*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v62
							if v62 <= v58 {
								v96 = v56
							} else {
								v65 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
								if v65 == base.I64_extend_i32_s(v56) {
									*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v58 - int64(1)
								} else {
								}
								*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(1)
								v76 = v56 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v16))) = v76
								v96 = v76
							}
							m.G0 = v12 + int32(16)
							return v96
						}
					}
				}
			}
		}
	}
}
func F_window_row_number_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	switch v6 - int32(462) {
	case 0:
		v10 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v10
		v13 = v5
	case 1:
		v10 = int32(1061)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v10
		v13 = v5
	default:
		v13 = int32(0)
	}
	return v13
}
