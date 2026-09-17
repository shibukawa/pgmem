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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
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
		v51 = l0
		m.G0 = v9 + int32(32)
		return v51
	} else {
		if base.Ui32(int32(536870911)) <= base.Ui32(l1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_perform_default_encoding_conversion_2), int32(0))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
						F_errdetail(m, int32(_a_F_perform_default_encoding_conversion_3), v9)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_perform_default_encoding_conversion_4), int32(825), int32(_a_F_perform_default_encoding_conversion_5))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
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
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
			if v14 != 0 {
				v27 = v24
			} else {
				v27 = v22
			}
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_perform_default_encoding_conversion[2]))
			v35 = F_MemoryContextAllocHuge(m, v30, l1<<(uint(int32(2))%32)|int32(1))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v40 = F_FunctionCall6Coll(m, v16, v26, v28, l0, v35, l1, int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					if base.Ui32(l1) < base.Ui32(int32(_a_F_perform_default_encoding_conversion_6)) {
						v51 = v35
						m.G0 = v9 + int32(32)
						return v51
					} else {
						v44 = F_strlen(m, v35)
						mBase = m.M
						if base.Ui32(int32(1073741823)) <= base.Ui32(v44) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_perform_default_encoding_conversion_2), int32(0))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
										F_errdetail(m, int32(_a_F_perform_default_encoding_conversion_3), v9+int32(16))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_perform_default_encoding_conversion_4), int32(852), int32(_a_F_perform_default_encoding_conversion_5))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
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
							v49 = F_repalloc(m, v35, v44+int32(1))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v51 = v49
								m.G0 = v9 + int32(32)
								return v51
							}
						}
					}
				}
			}
		}
	}
}
