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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_assign_stats_fetch_consistency[0]))
	if l0 != v4 {
		v7 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_assign_stats_fetch_consistency[1])) = uint8(v7)
	} else {
	}
	return
}
func F_set_stats_slot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	v11 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
	if v21 == l3&int32(_a_F_set_stats_slot_0) {
		if v21 != 0 {
			v27 = int32(-1)
		} else {
			v27 = int32(0)
		}
		v84 = v27
		v85 = v11
		v86 = v11
	} else {
		v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
		if v21 != 0 {
			if l3&int32(_a_F_set_stats_slot_0) != v28 {
				if v28 != 0 {
					v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)))
					if v49 != 0 {
						v50 = int32(-1)
					} else {
						v50 = int32(2)
					}
					v51 = v50
					v52 = v49
				} else {
					v45 = int32(1)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v51 = v45
					v52 = v46
				}
				v53 = int32(_a_F_set_stats_slot_0)
				v54 = l3 & v53
				if v54 == v52&v53 {
					v84 = v51
					v85 = int32(0)
					v86 = int32(2)
				} else {
					v60 = int32(3)
					v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
					if v64 != 0 {
						v65 = int32(-1)
					} else {
						v65 = v60
					}
					if v51 < int32(0) {
						v68 = v65
					} else {
						v68 = v51
					}
					if v54 == v64 {
						v84 = v68
						v85 = int32(0)
						v86 = v60
					} else {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v73 = int32(_a_F_set_stats_slot_0)
						v74 = v72 & v73
						v77 = base.B2i32(v74 != l3&v73)
						if v74 != l3&v73 {
							v78 = int32(5)
						} else {
							v78 = int32(4)
						}
						if v74 != 0 {
							v80 = v68
						} else {
							v80 = int32(4)
						}
						if v68 < int32(0) {
							v83 = v80
						} else {
							v83 = v68
						}
						v84 = v83
						v85 = v77
						v86 = v78
					}
				}
			} else {
				v32 = int32(1)
				if v28 != 0 {
					v35 = int32(-1)
				} else {
					v35 = v32
				}
				v84 = v35
				v85 = int32(0)
				v86 = v32
			}
		} else {
			v37 = int32(0)
			if l3&int32(_a_F_set_stats_slot_0) != v28 {
				v45 = v37
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v51 = v45
				v52 = v46
				v53 = int32(_a_F_set_stats_slot_0)
				v54 = l3 & v53
				if v54 == v52&v53 {
					v84 = v51
					v85 = int32(0)
					v86 = int32(2)
				} else {
					v60 = int32(3)
					v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
					if v64 != 0 {
						v65 = int32(-1)
					} else {
						v65 = v60
					}
					if v51 < int32(0) {
						v68 = v65
					} else {
						v68 = v51
					}
					if v54 == v64 {
						v84 = v68
						v85 = int32(0)
						v86 = v60
					} else {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v73 = int32(_a_F_set_stats_slot_0)
						v74 = v72 & v73
						v77 = base.B2i32(v74 != l3&v73)
						if v74 != l3&v73 {
							v78 = int32(5)
						} else {
							v78 = int32(4)
						}
						if v74 != 0 {
							v80 = v68
						} else {
							v80 = int32(4)
						}
						if v68 < int32(0) {
							v83 = v80
						} else {
							v83 = v68
						}
						v84 = v83
						v85 = v77
						v86 = v78
					}
				}
			} else {
				v84 = v37
				v85 = int32(0)
				v86 = int32(1)
			}
		}
	}
	if int32(0) <= v84 {
		v91 = v84
	} else {
		v91 = v86
	}
	if v85 != 0 {
		v92 = v91
	} else {
		v92 = v86
	}
	if v92 < int32(5) {
		v96 = v92 << (uint(int32(16)) % 32)
		v98 = v96 + int32(_a_F_set_stats_slot_1)
		v101 = l0 + int32(base.Ui32(v98)>>(uint(int32(14))%32))
		v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101))))
		if v102 != l3&int32(_a_F_set_stats_slot_0) {
			*(*int32)(unsafe.Add(mBase, uint32(v101))) = l3
			v110 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v98)>>(uint(int32(16))%32))))) = uint8(v110)
		} else {
		}
		v113 = v96 + int32(_a_F_set_stats_slot_2)
		v116 = l0 + int32(base.Ui32(v113)>>(uint(int32(14))%32))
		v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
		if l4 != v117 {
			*(*int32)(unsafe.Add(mBase, uint32(v116))) = l4
			v123 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v113)>>(uint(int32(16))%32))))) = uint8(v123)
		} else {
		}
		v126 = v96 - int32(-1048576)
		v129 = l0 + int32(base.Ui32(v126)>>(uint(int32(14))%32))
		v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
		if l5 != v130 {
			*(*int32)(unsafe.Add(mBase, uint32(v129))) = l5
			v136 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v126)>>(uint(int32(16))%32))))) = uint8(v136)
		} else {
		}
		if l7 == int32(0) {
			v141 = v92 + int32(21)
			*(*int32)(unsafe.Add(mBase, uint32(l0+v141<<(uint(int32(2))%32)))) = l6
			v147 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1+v141))) = uint8(v147)
			v150 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2+v141))) = uint8(v150)
		} else {
		}
		if l9 == int32(0) {
			v156 = v92 + int32(26)
			*(*int32)(unsafe.Add(mBase, uint32(l0+v156<<(uint(int32(2))%32)))) = l8
			v162 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1+v156))) = uint8(v162)
			v165 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2+v156))) = uint8(v165)
		} else {
		}
		m.G0 = v19 + int32(16)
		return
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v174 = m.ExcPending
		if v174 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = v92 + int32(1)
			F_errmsg(m, int32(_a_F_set_stats_slot_3), v19)
			mBase = m.M
			v180 = m.ExcPending
			if v180 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_set_stats_slot_4), int32(782), int32(_a_F_set_stats_slot_5))
				mBase = m.M
				v185 = m.ExcPending
				if v185 != 0 {
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
				F_errmsg(m, int32(_a_F_stats_check_required_arg_0), v7)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_stats_check_required_arg_1), int32(48), int32(_a_F_stats_check_required_arg_2))
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
