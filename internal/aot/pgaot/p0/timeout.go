package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RegisterTimeout(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	if base.Ui32(l0) < base.Ui32(int32(13)) {
		v53 = l0
		*(*int32)(unsafe.Add(mBase, uint32(v53*int32(40))+uint32(_consts[409]))) = l1
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[970]))
		if v7 == int32(0) {
			v53 = int32(13)
			*(*int32)(unsafe.Add(mBase, uint32(v53*int32(40))+uint32(_consts[409]))) = l1
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _consts[971]))
			if v11 == int32(0) {
				v53 = int32(14)
				*(*int32)(unsafe.Add(mBase, uint32(v53*int32(40))+uint32(_consts[409]))) = l1
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _consts[972]))
				if v16 == int32(0) {
					v53 = int32(15)
					*(*int32)(unsafe.Add(mBase, uint32(v53*int32(40))+uint32(_consts[409]))) = l1
					return
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, _consts[973]))
					if v21 == int32(0) {
						v53 = int32(16)
						*(*int32)(unsafe.Add(mBase, uint32(v53*int32(40))+uint32(_consts[409]))) = l1
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, _consts[974]))
						if v26 == int32(0) {
							v53 = int32(17)
							*(*int32)(unsafe.Add(mBase, uint32(v53*int32(40))+uint32(_consts[409]))) = l1
							return
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, _consts[975]))
							if v31 == int32(0) {
								v53 = int32(18)
								*(*int32)(unsafe.Add(mBase, uint32(v53*int32(40))+uint32(_consts[409]))) = l1
								return
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, _consts[976]))
								if v36 == int32(0) {
									v53 = int32(19)
									*(*int32)(unsafe.Add(mBase, uint32(v53*int32(40))+uint32(_consts[409]))) = l1
									return
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, _consts[977]))
									if v41 == int32(0) {
										v53 = int32(20)
										*(*int32)(unsafe.Add(mBase, uint32(v53*int32(40))+uint32(_consts[409]))) = l1
										return
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, _consts[978]))
										if v46 == int32(0) {
											v53 = int32(21)
											*(*int32)(unsafe.Add(mBase, uint32(v53*int32(40))+uint32(_consts[409]))) = l1
											return
										} else {
											v51 = *(*int32)(unsafe.Add(mBase, _consts[979]))
											if v51 != 0 {
												F_errstart_cold(m, int32(22), int32(0))
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return
												} else {
													F_errcode(m, int32(16581))
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return
													} else {
														F_errmsg(m, int32(136738), int32(0))
														mBase = m.M
														v69 = m.ExcPending
														if v69 != 0 {
															return
														} else {
															F_errfinish(m, int32(491989), int32(520), int32(66116))
															mBase = m.M
															v74 = m.ExcPending
															if v74 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v53 = int32(22)
												*(*int32)(unsafe.Add(mBase, uint32(v53*int32(40))+uint32(_consts[409]))) = l1
												return
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_enable_timeout_after(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v24 int64
	_ = v24
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	*(*int32)(unsafe.Add(mBase, _consts[404])) = int32(0)
	v10 = m.G0
	v11 = int32(16)
	v12 = v10 - v11
	m.G0 = v12
	F___gettimeofday(m, v12)
	mBase = m.M
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v16 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+8)))
	m.G0 = v12 + v11
	v24 = v16 + v15*int64(1000000) - int64(946684800000000)
	F_enable_timeout(m, l0, v24, v24+base.I64_extend_i32_s(l1)*int64(1000), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return
	} else {
		F_schedule_alarm(m, v24)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			return
		}
	}
}
