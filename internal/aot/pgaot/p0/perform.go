package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_perform_default_encoding_conversion(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v14 = l2 & int32(1)
	if v14 != 0 {
		v15 = int32(_a_F_perform_default_encoding_conversion_0)
	} else {
		v15 = int32(_a_F_perform_default_encoding_conversion_1)
	}
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v16 == int32(0) {
		v54 = l0
		m.G0 = v9 + int32(32)
		return v54
	} else {
		if base.Ui32(int32(536870911)) <= base.Ui32(l1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_perform_default_encoding_conversion_2), int32(0))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
						v75 = F_errdetail(m, int32(_a_F_perform_default_encoding_conversion_3), v9)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_perform_default_encoding_conversion_4), int32(827), int32(_a_F_perform_default_encoding_conversion_5))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
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
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_perform_default_encoding_conversion[0]))
			v24 = *(*int32)(unsafe.Add(mBase, _c_F_perform_default_encoding_conversion[1]))
			if v14 != 0 {
				v25 = v22
			} else {
				v25 = v24
			}
			v26 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+4)))
			if v14 != 0 {
				v27 = v24
			} else {
				v27 = v22
			}
			v28 = int64(*(*int32)(unsafe.Add(mBase, uint32(v27)+4)))
			v31 = *(*int32)(unsafe.Add(mBase, _c_F_perform_default_encoding_conversion[2]))
			v36 = F_MemoryContextAllocHuge(m, v31, l1<<(uint(int32(2))%32)|int32(1))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v43 = F_FunctionCall6Coll(m, v16, v26, v28, base.I64_extend_i32_u(l0), base.I64_extend_i32_u(v36), base.I64_extend_i32_u(l1), int64(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					if base.Ui32(l1) < base.Ui32(int32(_a_F_perform_default_encoding_conversion_6)) {
						v54 = v36
						m.G0 = v9 + int32(32)
						return v54
					} else {
						v47 = F_strlen(m, v36)
						mBase = m.M
						if base.Ui32(int32(1073741823)) <= base.Ui32(v47) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_perform_default_encoding_conversion_2), int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
										v97 = F_errdetail(m, int32(_a_F_perform_default_encoding_conversion_3), v9+int32(16))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_perform_default_encoding_conversion_4), int32(854), int32(_a_F_perform_default_encoding_conversion_5))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
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
						} else {
							v52 = F_repalloc(m, v36, v47+int32(1))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v54 = v52
								m.G0 = v9 + int32(32)
								return v54
							}
						}
					}
				}
			}
		}
	}
}
