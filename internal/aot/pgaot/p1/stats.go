package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assign_stats_fetch_consistency(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, _consts[925]))
	if l0 != v4 {
		v7 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _consts[932])) = uint8(v7)
	} else {
	}
	return
}
func F_set_stats_slot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	v11 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
	if v24 != 0 {
		v25 = int32(-1)
	} else {
		v25 = v11
	}
	v27 = l3 & int32(65535)
	if v24 == v27 {
		v68 = v25
		v69 = v11
		v71 = v11
	} else {
		v29 = int32(1)
		v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
		if v31 != 0 {
			v32 = v25
		} else {
			v32 = v29
		}
		if v24 != 0 {
			v34 = v32
		} else {
			v34 = int32(0)
		}
		if v27 == v31 {
			v68 = v34
			v69 = v29
			v71 = v11
		} else {
			v36 = int32(2)
			v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)))
			if v38 != 0 {
				v39 = v34
			} else {
				v39 = v36
			}
			if v34 < int32(0) {
				v42 = v39
			} else {
				v42 = v34
			}
			v44 = l3 & int32(65535)
			if v44 == v38 {
				v68 = v42
				v69 = v36
				v71 = v11
			} else {
				v46 = int32(3)
				v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
				if v48 != 0 {
					v49 = v42
				} else {
					v49 = v46
				}
				if v42 < int32(0) {
					v52 = v49
				} else {
					v52 = v42
				}
				if v48 == v44 {
					v68 = v52
					v69 = v46
					v71 = v11
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v57 = int32(65535)
					v58 = v56 & v57
					v61 = base.B2i32(v58 != l3&v57)
					if v58 != l3&v57 {
						v62 = int32(5)
					} else {
						v62 = int32(4)
					}
					if v58 != 0 {
						v64 = v52
					} else {
						v64 = int32(4)
					}
					if v52 < int32(0) {
						v67 = v64
					} else {
						v67 = v52
					}
					v68 = v67
					v69 = v62
					v71 = v61
				}
			}
		}
	}
	if int32(0) <= v68 {
		v76 = v68
	} else {
		v76 = v69
	}
	if v71 != 0 {
		v77 = v76
	} else {
		v77 = v69
	}
	if v77 < int32(5) {
		v81 = v77 << (uint(int32(16)) % 32)
		v83 = v81 + int32(393216)
		v86 = l0 + int32(base.Ui32(v83)>>(uint(int32(14))%32))
		v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86))))
		if v87 != l3&int32(65535) {
			*(*int32)(unsafe.Add(mBase, uint32(v86))) = l3
			v95 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v83)>>(uint(int32(16))%32))))) = uint8(v95)
		} else {
		}
		v98 = v81 + int32(720896)
		v101 = l0 + int32(base.Ui32(v98)>>(uint(int32(14))%32))
		v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
		if l4 != v102 {
			*(*int32)(unsafe.Add(mBase, uint32(v101))) = l4
			v108 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v98)>>(uint(int32(16))%32))))) = uint8(v108)
		} else {
		}
		v111 = v81 - int32(-1048576)
		v114 = l0 + int32(base.Ui32(v111)>>(uint(int32(14))%32))
		v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
		if l5 != v115 {
			*(*int32)(unsafe.Add(mBase, uint32(v114))) = l5
			v121 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v111)>>(uint(int32(16))%32))))) = uint8(v121)
		} else {
		}
		if l7 == int32(0) {
			v126 = v77 + int32(21)
			*(*int32)(unsafe.Add(mBase, uint32(l0+v126<<(uint(int32(2))%32)))) = l6
			v132 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1+v126))) = uint8(v132)
			v135 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2+v126))) = uint8(v135)
		} else {
		}
		if l9 == int32(0) {
			v141 = v77 + int32(26)
			*(*int32)(unsafe.Add(mBase, uint32(l0+v141<<(uint(int32(2))%32)))) = l8
			v147 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1+v141))) = uint8(v147)
			v150 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2+v141))) = uint8(v150)
		} else {
		}
		m.G0 = v20 + int32(16)
		return
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v159 = m.ExcPending
		if v159 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v20))) = v77 + int32(1)
			F_errmsg(m, int32(465056), v20)
			mBase = m.M
			v165 = m.ExcPending
			if v165 != 0 {
				return
			} else {
				F_errfinish(m, int32(470563), int32(782), int32(80351))
				mBase = m.M
				v170 = m.ExcPending
				if v170 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_stats_check_required_arg(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = l2 << (uint(int32(3)) % 32)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v10)+24)))
	if v12 == int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1+v10)))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v23
				F_errmsg(m, int32(288387), v7)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					F_errfinish(m, int32(470855), int32(48), int32(310610))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
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
		m.G0 = v7 + int32(16)
		return
	}
}
