package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AssignPostmasterChildSlot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = l0 << (uint(int32(4)) % 32)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[650])))
	if v18 != 0 {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[651])))
		if v19 == int32(0) {
			v120 = v2
			m.G0 = v12 + int32(48)
			return v120
		} else {
			if v19 == v15+int32(4428424) {
				v120 = v2
				m.G0 = v12 + int32(48)
				return v120
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v26
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				*(*int32)(unsafe.Add(mBase, uint32(v26))) = v28
				v32 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v19-int32(4)))) = uint8(v32)
				v36 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v19-int32(8)))) = v36
				v39 = v19 - int32(12)
				*(*int32)(unsafe.Add(mBase, uint32(v39))) = l0
				v42 = v19 - int32(20)
				*(*int32)(unsafe.Add(mBase, uint32(v42))) = v36
				v46 = v19 - int32(16)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[652])))
				if v47 < v48 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v145 = m.ExcPending
					if v145 != 0 {
						return int32(0)
					} else {
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v146
						F_errmsg_internal(m, int32(83824), v12+int32(16))
						mBase = m.M
						v152 = m.ExcPending
						if v152 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(500902), int32(188), int32(86597))
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[650])))
					if v50+v48 <= v47 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v145 = m.ExcPending
						if v145 != 0 {
							return int32(0)
						} else {
							v146 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v146
							F_errmsg_internal(m, int32(83824), v12+int32(16))
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(500902), int32(188), int32(86597))
								mBase = m.M
								v157 = m.ExcPending
								if v157 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, _consts[653]))
						if v54 == int32(0) {
							v57 = int32(4428704)
							*(*int32)(unsafe.Add(mBase, _consts[654])) = v57
							v61 = v57
						} else {
							v61 = v54
						}
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(4428704)
						*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v61
						*(*int32)(unsafe.Add(mBase, uint32(v61))) = v19
						*(*int32)(unsafe.Add(mBase, _consts[653])) = v19
						v69 = *(*int32)(unsafe.Add(mBase, _consts[655]))
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
						v75 = v69 + v70<<(uint(int32(2))%32) + int32(44)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
						if v76 != 0 {
							F_errstart_cold(m, int32(22), int32(0))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(360879), int32(0))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(498737), int32(236), int32(454286))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v75))) = int32(1)
							v96 = F_errstart(m, int32(13), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								if v96 == int32(0) {
									v120 = v42
									m.G0 = v12 + int32(48)
									return v120
								} else {
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
									v105 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(12))+uint32(_consts[656])))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v105
									*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v100
									F_errmsg_internal(m, int32(182069), v12+int32(32))
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(500902), int32(197), int32(86597))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return int32(0)
										} else {
											v120 = v42
											m.G0 = v12 + int32(48)
											return v120
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v132 = m.ExcPending
		if v132 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
			F_errmsg_internal(m, int32(477706), v12)
			mBase = m.M
			v136 = m.ExcPending
			if v136 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(500902), int32(168), int32(86597))
				mBase = m.M
				v141 = m.ExcPending
				if v141 != 0 {
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
