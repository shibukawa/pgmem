package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DecodeTimezoneNameToTz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v11 = F_DecodeTimezoneName(m, l0, v5+int32(8), v5+int32(12))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			v21 = v15
			m.G0 = v5 + int32(16)
			return v21
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			v19 = F_pg_tzset_offset(m, int32(0)-v17)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = v19
				m.G0 = v5 + int32(16)
				return v21
			}
		}
	}
}
func F_EncodeTimezone(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	if l1 <= int32(0) {
		v12 = int32(43)
	} else {
		v12 = int32(45)
	}
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v12)
	v15 = l1 >> (uint(int32(31)) % 32)
	v17 = l1 ^ v15 - v15
	v19 = base.I32_div_s(v17, int32(3600))
	v20 = int32(-60)
	v23 = base.I32_div_s(v17, int32(60))
	v24 = v19*v20 + v23
	v26 = l0 + int32(1)
	v29 = v23*v20 + v17
	if v29 != 0 {
		v30 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v19)) == int32(0) {
			v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19<<(uint(int32(1))%32))+uint32(_c_F_EncodeTimezone[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v41)
			v56 = l0 + int32(3)
		} else {
			v45 = F_pg_ultoa_n(m, v19, v26)
			mBase = m.M
			if v30 <= v45 {
				v56 = v26 + v45
			} else {
				v48 = l0 + int32(3)
				if v45 != 0 {
					base.MemoryCopy(m, v48-v45, v26, v45)
				} else {
				}
				v51 = v30 - v45
				if v51 != 0 {
					base.MemoryFill(m, v26, int32(48), v51)
				} else {
				}
				v56 = v48
			}
		}
		v57 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v57)
		v60 = v56 + int32(1)
		v61 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v24)) == int32(0) {
			v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24<<(uint(int32(1))%32))+uint32(_c_F_EncodeTimezone[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v60))) = uint16(v72)
			v87 = v56 + int32(3)
		} else {
			v76 = F_pg_ultoa_n(m, v24, v60)
			mBase = m.M
			if v61 <= v76 {
				v87 = v60 + v76
			} else {
				v79 = v56 + int32(3)
				if v76 != 0 {
					base.MemoryCopy(m, v79-v76, v60, v76)
				} else {
				}
				v82 = v61 - v76
				if v82 != 0 {
					base.MemoryFill(m, v60, int32(48), v82)
				} else {
				}
				v87 = v79
			}
		}
		v119 = v87
		v120 = v29
		v121 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v121)
		v124 = v119 + int32(1)
		v125 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v120)) == int32(0) {
			v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120<<(uint(int32(1))%32))+uint32(_c_F_EncodeTimezone[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v124))) = uint16(v136)
			v151 = v119 + int32(3)
		} else {
			v140 = F_pg_ultoa_n(m, v120, v124)
			mBase = m.M
			if v125 <= v140 {
				v151 = v124 + v140
			} else {
				v143 = v119 + int32(3)
				if v140 != 0 {
					base.MemoryCopy(m, v143-v140, v124, v140)
				} else {
				}
				v146 = v125 - v140
				if v146 != 0 {
					base.MemoryFill(m, v124, int32(48), v146)
				} else {
				}
				v151 = v143
			}
		}
		v152 = v151
	} else {
		v88 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v19)) == int32(0) {
			v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19<<(uint(int32(1))%32))+uint32(_c_F_EncodeTimezone[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v99)
			v114 = l0 + int32(3)
		} else {
			v103 = F_pg_ultoa_n(m, v19, v26)
			mBase = m.M
			if v88 <= v103 {
				v114 = v26 + v103
			} else {
				v106 = l0 + int32(3)
				if v103 != 0 {
					base.MemoryCopy(m, v106-v103, v26, v103)
				} else {
				}
				v109 = v88 - v103
				if v109 != 0 {
					base.MemoryFill(m, v26, int32(48), v109)
				} else {
				}
				v114 = v106
			}
		}
		if l2 == int32(4) {
			v119 = v114
			v120 = v24
			v121 = int32(58)
			*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v121)
			v124 = v119 + int32(1)
			v125 = int32(2)
			if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v120)) == int32(0) {
				v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120<<(uint(int32(1))%32))+uint32(_c_F_EncodeTimezone[0]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v124))) = uint16(v136)
				v151 = v119 + int32(3)
			} else {
				v140 = F_pg_ultoa_n(m, v120, v124)
				mBase = m.M
				if v125 <= v140 {
					v151 = v124 + v140
				} else {
					v143 = v119 + int32(3)
					if v140 != 0 {
						base.MemoryCopy(m, v143-v140, v124, v140)
					} else {
					}
					v146 = v125 - v140
					if v146 != 0 {
						base.MemoryFill(m, v124, int32(48), v146)
					} else {
					}
					v151 = v143
				}
			}
			v152 = v151
		} else {
			if v24 == int32(0) {
				v152 = v114
			} else {
				v119 = v114
				v120 = v24
				v121 = int32(58)
				*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v121)
				v124 = v119 + int32(1)
				v125 = int32(2)
				if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v120)) == int32(0) {
					v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120<<(uint(int32(1))%32))+uint32(_c_F_EncodeTimezone[0]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v124))) = uint16(v136)
					v151 = v119 + int32(3)
				} else {
					v140 = F_pg_ultoa_n(m, v120, v124)
					mBase = m.M
					if v125 <= v140 {
						v151 = v124 + v140
					} else {
						v143 = v119 + int32(3)
						if v140 != 0 {
							base.MemoryCopy(m, v143-v140, v124, v140)
						} else {
						}
						v146 = v125 - v140
						if v146 != 0 {
							base.MemoryFill(m, v124, int32(48), v146)
						} else {
						}
						v151 = v143
					}
				}
				v152 = v151
			}
		}
	}
	return v152
}
func F_assign_timezone(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_timezone[0])) = v4
	base.MemoryFill(m, int32(_a_F_assign_timezone_0), int32(0), int32(500))
	return
}
