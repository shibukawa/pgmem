package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AssignPostmasterChildSlot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = l0 << (uint(int32(4)) % 32)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_AssignPostmasterChildSlot[0])))
	if v16 != 0 {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_AssignPostmasterChildSlot[1])))
		if base.B2i32(v19 == int32(0))|base.B2i32(v19 == v15+int32(_a_F_AssignPostmasterChildSlot_0)) != 0 {
			v121 = int32(0)
			m.G0 = v12 + int32(48)
			return v121
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v27
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			*(*int32)(unsafe.Add(mBase, uint32(v27))) = v29
			v33 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v19-int32(4)))) = uint8(v33)
			v37 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v19-int32(8)))) = v37
			v40 = v19 - int32(12)
			*(*int32)(unsafe.Add(mBase, uint32(v40))) = l0
			v43 = v19 - int32(20)
			*(*int32)(unsafe.Add(mBase, uint32(v43))) = v37
			v47 = v19 - int32(16)
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_AssignPostmasterChildSlot[2])))
			if v48 < v49 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v144 = m.ExcPending
				if v144 != 0 {
					return int32(0)
				} else {
					v145 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v145
					F_errmsg_internal(m, int32(_a_F_AssignPostmasterChildSlot_1), v12+int32(16))
					mBase = m.M
					v151 = m.ExcPending
					if v151 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_AssignPostmasterChildSlot_2), int32(188), int32(_a_F_AssignPostmasterChildSlot_3))
						mBase = m.M
						v156 = m.ExcPending
						if v156 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_AssignPostmasterChildSlot[0])))
				if v51+v49 <= v48 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v144 = m.ExcPending
					if v144 != 0 {
						return int32(0)
					} else {
						v145 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v145
						F_errmsg_internal(m, int32(_a_F_AssignPostmasterChildSlot_1), v12+int32(16))
						mBase = m.M
						v151 = m.ExcPending
						if v151 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_AssignPostmasterChildSlot_2), int32(188), int32(_a_F_AssignPostmasterChildSlot_3))
							mBase = m.M
							v156 = m.ExcPending
							if v156 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, _c_F_AssignPostmasterChildSlot[3]))
					if v55 == int32(0) {
						v58 = int32(_a_F_AssignPostmasterChildSlot_4)
						*(*int32)(unsafe.Add(mBase, _c_F_AssignPostmasterChildSlot[4])) = v58
						v62 = v58
					} else {
						v62 = v55
					}
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(_a_F_AssignPostmasterChildSlot_4)
					*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v62
					*(*int32)(unsafe.Add(mBase, uint32(v62))) = v19
					*(*int32)(unsafe.Add(mBase, _c_F_AssignPostmasterChildSlot[3])) = v19
					v70 = *(*int32)(unsafe.Add(mBase, _c_F_AssignPostmasterChildSlot[5]))
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
					v76 = v70 + v71<<(uint(int32(2))%32) + int32(44)
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
					if v77 != 0 {
						F_errstart_cold(m, int32(22), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_AssignPostmasterChildSlot_5), int32(0))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_AssignPostmasterChildSlot_6), int32(236), int32(_a_F_AssignPostmasterChildSlot_7))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v76))) = int32(1)
						v97 = F_errstart(m, int32(13), int32(0))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int32(0)
						} else {
							if v97 == int32(0) {
								v121 = v43
								m.G0 = v12 + int32(48)
								return v121
							} else {
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
								v104 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(12))+uint32(_c_F_AssignPostmasterChildSlot[6])))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v104
								*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v101
								F_errmsg_internal(m, int32(_a_F_AssignPostmasterChildSlot_8), v12+int32(32))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_AssignPostmasterChildSlot_2), int32(197), int32(_a_F_AssignPostmasterChildSlot_3))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int32(0)
									} else {
										v121 = v43
										m.G0 = v12 + int32(48)
										return v121
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
		v131 = m.ExcPending
		if v131 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
			F_errmsg_internal(m, int32(_a_F_AssignPostmasterChildSlot_9), v12)
			mBase = m.M
			v135 = m.ExcPending
			if v135 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_AssignPostmasterChildSlot_2), int32(168), int32(_a_F_AssignPostmasterChildSlot_3))
				mBase = m.M
				v140 = m.ExcPending
				if v140 != 0 {
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
