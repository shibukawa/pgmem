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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	if l1 <= int32(0) {
		v11 = int32(43)
	} else {
		v11 = int32(45)
	}
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v11)
	v14 = l1 >> (uint(int32(31)) % 32)
	v16 = l1 ^ v14 - v14
	v18 = base.I32_div_s(v16, int32(3600))
	v19 = int32(-60)
	v22 = base.I32_div_s(v16, int32(60))
	v23 = v18*v19 + v22
	v25 = l0 + int32(1)
	v28 = v22*v19 + v16
	if v28 != 0 {
		v29 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v18) {
			v43 = F_pg_ultoa_n(m, v18, v25)
			mBase = m.M
			if v29 <= v43 {
				v54 = v25 + v43
			} else {
				v46 = l0 + int32(3)
				v48 = F_memmove(m, v46-v43, v25, v43)
				mBase = m.M
				v51 = F___memset(m, v25, int32(48), v29-v43)
				mBase = m.M
				v54 = v46
			}
		} else {
			v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18<<(uint(int32(1))%32))+uint32(_consts[4]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v25))) = uint16(v39)
			v54 = l0 + int32(3)
		}
		v55 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v55)
		v58 = v54 + int32(1)
		v59 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v23) {
			v73 = F_pg_ultoa_n(m, v23, v58)
			mBase = m.M
			if v59 <= v73 {
				v84 = v58 + v73
			} else {
				v76 = v54 + int32(3)
				v78 = F_memmove(m, v76-v73, v58, v73)
				mBase = m.M
				v81 = F___memset(m, v58, int32(48), v59-v73)
				mBase = m.M
				v84 = v76
			}
		} else {
			v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23<<(uint(int32(1))%32))+uint32(_consts[4]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v58))) = uint16(v69)
			v84 = v54 + int32(3)
		}
		v115 = v84
		v116 = v28
		v117 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v117)
		v120 = v115 + int32(1)
		v121 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v116) {
			v135 = F_pg_ultoa_n(m, v116, v120)
			mBase = m.M
			if v121 <= v135 {
				v146 = v120 + v135
			} else {
				v138 = v115 + int32(3)
				v140 = F_memmove(m, v138-v135, v120, v135)
				mBase = m.M
				v143 = F___memset(m, v120, int32(48), v121-v135)
				mBase = m.M
				v146 = v138
			}
		} else {
			v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116<<(uint(int32(1))%32))+uint32(_consts[4]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v120))) = uint16(v131)
			v146 = v115 + int32(3)
		}
		v147 = v146
	} else {
		v85 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v18) {
			v99 = F_pg_ultoa_n(m, v18, v25)
			mBase = m.M
			if v85 <= v99 {
				v110 = v25 + v99
			} else {
				v102 = l0 + int32(3)
				v104 = F_memmove(m, v102-v99, v25, v99)
				mBase = m.M
				v107 = F___memset(m, v25, int32(48), v85-v99)
				mBase = m.M
				v110 = v102
			}
		} else {
			v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18<<(uint(int32(1))%32))+uint32(_consts[4]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v25))) = uint16(v95)
			v110 = l0 + int32(3)
		}
		if l2 == int32(4) {
			v115 = v110
			v116 = v23
			v117 = int32(58)
			*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v117)
			v120 = v115 + int32(1)
			v121 = int32(2)
			if base.Ui32(int32(99)) < base.Ui32(v116) {
				v135 = F_pg_ultoa_n(m, v116, v120)
				mBase = m.M
				if v121 <= v135 {
					v146 = v120 + v135
				} else {
					v138 = v115 + int32(3)
					v140 = F_memmove(m, v138-v135, v120, v135)
					mBase = m.M
					v143 = F___memset(m, v120, int32(48), v121-v135)
					mBase = m.M
					v146 = v138
				}
			} else {
				v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116<<(uint(int32(1))%32))+uint32(_consts[4]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v120))) = uint16(v131)
				v146 = v115 + int32(3)
			}
			v147 = v146
		} else {
			if v23 == int32(0) {
				v147 = v110
			} else {
				v115 = v110
				v116 = v23
				v117 = int32(58)
				*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v117)
				v120 = v115 + int32(1)
				v121 = int32(2)
				if base.Ui32(int32(99)) < base.Ui32(v116) {
					v135 = F_pg_ultoa_n(m, v116, v120)
					mBase = m.M
					if v121 <= v135 {
						v146 = v120 + v135
					} else {
						v138 = v115 + int32(3)
						v140 = F_memmove(m, v138-v135, v120, v135)
						mBase = m.M
						v143 = F___memset(m, v120, int32(48), v121-v135)
						mBase = m.M
						v146 = v138
					}
				} else {
					v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116<<(uint(int32(1))%32))+uint32(_consts[4]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v120))) = uint16(v131)
					v146 = v115 + int32(3)
				}
				v147 = v146
			}
		}
	}
	return v147
}
func F_assign_timezone(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _consts[329])) = v4
	v10 = F__emscripten_memset_bulkmem(m, int32(4498624), base.I32_extend8_s(int32(0)), int32(500))
	mBase = m.M
	return
}
