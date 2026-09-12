package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_window_first_value(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = int32(1)
	v14 = F_WinGetFuncArgInFrame(m, v8, int32(0), v10, v10, v6+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v24 = int32(0)
		} else {
			v24 = v14
		}
		m.G0 = v6 + int32(16)
		return v24
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
	var v52 int64
	_ = v52
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v67 int64
	_ = v67
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
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
			v58 = v20
			v61 = v21
			v64 = v22 + int64(1)
			*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v64
			if v64 <= v61 {
				v98 = v58
			} else {
				v67 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
				if v67 == base.I64_extend_i32_s(v58) {
					*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v61 - int64(1)
				} else {
				}
				*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(1)
				v78 = v58 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v78
				v98 = v78
			}
			m.G0 = v12 + int32(16)
			return v98
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
						v96 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
						v98 = v27
						m.G0 = v12 + int32(16)
						return v98
					} else {
						if v31 <= int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67371138))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(251146), int32(0))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(516710), int32(443), int32(403163))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
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
									v55 = v39
								} else {
									v52 = v39 + int64(1)
									*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v52
									v55 = v52
								}
							} else {
								v52 = int64(1)
								*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v52
								v55 = v52
							}
							v58 = int32(1)
							v61 = v55
							v64 = int64(1)
							*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v64
							if v64 <= v61 {
								v98 = v58
							} else {
								v67 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
								if v67 == base.I64_extend_i32_s(v58) {
									*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v61 - int64(1)
								} else {
								}
								*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(1)
								v78 = v58 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v16))) = v78
								v98 = v78
							}
							m.G0 = v12 + int32(16)
							return v98
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
