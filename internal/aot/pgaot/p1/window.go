package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_window_first_value(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14404(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_window_ntile(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v49 int64
	_ = v49
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int64
	_ = v107
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_WinCheckAndInitializeNullTreatment(m, v14, int32(0), l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		v21 = F_WinGetPartitionLocalMemory(m, v14, int32(32))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			if v23 != 0 {
				v24 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
				v25 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
				v60 = v24
				v64 = v23
				v65 = v25 + int64(1)
				*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v65
				if v60 < v65 {
					v68 = *(*int64)(unsafe.Add(mBase, uint32(v21)+24))
					if v68 == base.I64_extend_i32_s(v64) {
						*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v60 - int64(1)
					} else {
					}
					*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = int64(1)
					v79 = v64 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v79
					v81 = v79
				} else {
					v81 = v64
				}
				v107 = base.I64_extend_i32_s(v81)
				m.G0 = v12 + int32(16)
				return v107
			} else {
				v28 = F_WinGetPartitionRowCount(m, v14)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v33 = F_WinGetFuncArgCurrent(m, v14, int32(0), v12+int32(15))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
						if v35 != 0 {
							v99 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v99)
							v107 = int64(0)
							m.G0 = v12 + int32(16)
							return v107
						} else {
							if base.I32_wrap_i64(v33) <= int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(67371138))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_window_ntile_0), int32(0))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_window_ntile_1), int32(449), int32(_a_F_window_ntile_2))
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
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
								*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(1)
								v42 = v33 & int64(2147483647)
								v43 = base.I64_div_s(v28, v42)
								*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v43
								if int64(0) < v43 {
									v49 = v28 - v43*v42
									*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v49
									if v49 == int64(0) {
										v57 = v43
									} else {
										v55 = v43 + int64(1)
										*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v55
										v57 = v55
									}
								} else {
									v55 = int64(1)
									*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v55
									v57 = v55
								}
								v60 = v57
								v64 = int32(1)
								v65 = int64(1)
								*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v65
								if v60 < v65 {
									v68 = *(*int64)(unsafe.Add(mBase, uint32(v21)+24))
									if v68 == base.I64_extend_i32_s(v64) {
										*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v60 - int64(1)
									} else {
									}
									*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = int64(1)
									v79 = v64 + int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v21))) = v79
									v81 = v79
								} else {
									v81 = v64
								}
								v107 = base.I64_extend_i32_s(v81)
								m.G0 = v12 + int32(16)
								return v107
							}
						}
					}
				}
			}
		}
	}
}
func F_window_row_number_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int64
	_ = v17
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = base.I32_wrap_i64(v6)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	switch v8 - int32(470) {
	case 0:
		v12 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v12
		v17 = v6 & int64(4294967295)
	case 1:
		v12 = int32(1061)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v12
		v17 = v6 & int64(4294967295)
	default:
		v17 = int64(0)
	}
	return v17
}
