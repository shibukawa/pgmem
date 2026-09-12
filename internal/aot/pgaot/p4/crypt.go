package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__crypt_gensalt_blowfish_rn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	if l2 < int32(16) {
		if int32(0) < l4 {
			v219 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v219)
			return v219
		} else {
			return int32(0)
		}
	} else {
		if l4 < int32(30) {
			if int32(0) < l4 {
				v219 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v219)
				return v219
			} else {
				return int32(0)
			}
		} else {
			if l0 != 0 {
				if base.Ui32(l0-int32(32)) < base.Ui32(int32(-28)) {
					v219 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v219)
					return v219
				} else {
					v19 = l0
					v20 = int32(36)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v20)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(610349604)
					v26 = int32(10)
					v27 = base.I32_div_u_s(v19&int32(255), v26)
					v28 = int32(48)
					v29 = v27 + v28
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v29)
					v35 = v19 - v27*v26 | v28
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v35)
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
					v39 = int32(2)
					v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v38)>>(uint(v39)%32)))+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)) = uint8(v42)
					v44 = int32(4)
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38<<(uint(v44)%32)&v28|int32(base.Ui32(v48)>>(uint(v44)%32)))+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v53)
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
					v56 = int32(63)
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55&v56)+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+10)) = uint8(v59)
					v63 = int32(60)
					v65 = int32(6)
					v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48<<(uint(v39)%32)&v63|int32(base.Ui32(v55)>>(uint(v65)%32)))+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v69)
					v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
					v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v71)>>(uint(v39)%32)))+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)) = uint8(v75)
					v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
					v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71<<(uint(v44)%32)&v28|int32(base.Ui32(v81)>>(uint(v44)%32)))+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)) = uint8(v86)
					v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
					v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88&v56)+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)) = uint8(v92)
					v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81<<(uint(v39)%32)&v63|int32(base.Ui32(v88)>>(uint(v65)%32)))+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)) = uint8(v102)
					v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
					v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v104)>>(uint(v39)%32)))+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+15)) = uint8(v108)
					v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
					v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104<<(uint(v44)%32)&v28|int32(base.Ui32(v114)>>(uint(v44)%32)))+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)) = uint8(v119)
					v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
					v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121&v56)+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+18)) = uint8(v125)
					v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114<<(uint(v39)%32)&v63|int32(base.Ui32(v121)>>(uint(v65)%32)))+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+17)) = uint8(v135)
					v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
					v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v137)>>(uint(v39)%32)))+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+19)) = uint8(v141)
					v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
					v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137<<(uint(v44)%32)&v28|int32(base.Ui32(v147)>>(uint(v44)%32)))+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)) = uint8(v152)
					v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
					v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154&v56)+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+22)) = uint8(v158)
					v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147<<(uint(v39)%32)&v63|int32(base.Ui32(v154)>>(uint(v65)%32)))+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)) = uint8(v168)
					v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
					v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v170)>>(uint(v39)%32)))+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+23)) = uint8(v174)
					v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
					v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170<<(uint(v44)%32)&v28|int32(base.Ui32(v180)>>(uint(v44)%32)))+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+24)) = uint8(v185)
					v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
					v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187&v56)+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+26)) = uint8(v191)
					v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180<<(uint(v39)%32)&v63|int32(base.Ui32(v187)>>(uint(v65)%32)))+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+25)) = uint8(v201)
					v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)))
					v204 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v204)
					v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v203)>>(uint(v39)%32)))+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+27)) = uint8(v209)
					v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203<<(uint(v44)%32)&v28)+uint32(_consts[1528]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)) = uint8(v216)
					return l3
				}
			} else {
				v19 = int32(5)
				v20 = int32(36)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v20)
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(610349604)
				v26 = int32(10)
				v27 = base.I32_div_u_s(v19&int32(255), v26)
				v28 = int32(48)
				v29 = v27 + v28
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v29)
				v35 = v19 - v27*v26 | v28
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v35)
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				v39 = int32(2)
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v38)>>(uint(v39)%32)))+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)) = uint8(v42)
				v44 = int32(4)
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38<<(uint(v44)%32)&v28|int32(base.Ui32(v48)>>(uint(v44)%32)))+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v53)
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
				v56 = int32(63)
				v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55&v56)+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+10)) = uint8(v59)
				v63 = int32(60)
				v65 = int32(6)
				v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48<<(uint(v39)%32)&v63|int32(base.Ui32(v55)>>(uint(v65)%32)))+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v69)
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v71)>>(uint(v39)%32)))+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)) = uint8(v75)
				v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
				v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71<<(uint(v44)%32)&v28|int32(base.Ui32(v81)>>(uint(v44)%32)))+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)) = uint8(v86)
				v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
				v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88&v56)+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)) = uint8(v92)
				v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81<<(uint(v39)%32)&v63|int32(base.Ui32(v88)>>(uint(v65)%32)))+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)) = uint8(v102)
				v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
				v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v104)>>(uint(v39)%32)))+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+15)) = uint8(v108)
				v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
				v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104<<(uint(v44)%32)&v28|int32(base.Ui32(v114)>>(uint(v44)%32)))+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)) = uint8(v119)
				v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
				v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121&v56)+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+18)) = uint8(v125)
				v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114<<(uint(v39)%32)&v63|int32(base.Ui32(v121)>>(uint(v65)%32)))+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+17)) = uint8(v135)
				v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
				v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v137)>>(uint(v39)%32)))+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+19)) = uint8(v141)
				v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
				v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137<<(uint(v44)%32)&v28|int32(base.Ui32(v147)>>(uint(v44)%32)))+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)) = uint8(v152)
				v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
				v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154&v56)+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+22)) = uint8(v158)
				v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147<<(uint(v39)%32)&v63|int32(base.Ui32(v154)>>(uint(v65)%32)))+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)) = uint8(v168)
				v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
				v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v170)>>(uint(v39)%32)))+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+23)) = uint8(v174)
				v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
				v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170<<(uint(v44)%32)&v28|int32(base.Ui32(v180)>>(uint(v44)%32)))+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+24)) = uint8(v185)
				v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
				v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187&v56)+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+26)) = uint8(v191)
				v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180<<(uint(v39)%32)&v63|int32(base.Ui32(v187)>>(uint(v65)%32)))+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+25)) = uint8(v201)
				v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)))
				v204 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v204)
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v203)>>(uint(v39)%32)))+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+27)) = uint8(v209)
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203<<(uint(v44)%32)&v28)+uint32(_consts[1528]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)) = uint8(v216)
				return l3
			}
		}
	}
}
func F__crypt_gensalt_md5_rn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	if l2 < int32(3) {
		if int32(0) < l4 {
			v106 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v106)
			return v106
		} else {
			return int32(0)
		}
	} else {
		if l4 < int32(8) {
			if int32(0) < l4 {
				v106 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v106)
				return v106
			} else {
				return int32(0)
			}
		} else {
			if l0 == int32(0) {
				v21 = int32(36)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)) = uint8(v21)
				v23 = int32(12580)
				*(*uint16)(unsafe.Add(mBase, uint32(l3))) = uint16(v23)
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
				v28 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)) = uint8(v28)
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v27)>>(uint(int32(2))%32)))+uint32(_consts[1527]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v34)
				v36 = int32(63)
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26&v36)+uint32(_consts[1527]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+3)) = uint8(v39)
				v44 = v25 << (uint(int32(8)) % 32)
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v27<<(uint(int32(16))%32)|v44)>>(uint(int32(12))%32))&v36)+uint32(_consts[1527]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v51)
				v54 = int32(6)
				v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v44|v26)>>(uint(v54)%32))&v36)+uint32(_consts[1527]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v59)
				if base.Ui32(l2) < base.Ui32(v54) {
				} else {
					if base.Ui32(l4) < base.Ui32(int32(12)) {
					} else {
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
						v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
						v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
						v68 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)) = uint8(v68)
						v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v67)>>(uint(int32(2))%32)))+uint32(_consts[1527]))))
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+10)) = uint8(v74)
						v76 = int32(63)
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66&v76)+uint32(_consts[1527]))))
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)) = uint8(v79)
						v84 = v65 << (uint(int32(8)) % 32)
						v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v67<<(uint(int32(16))%32)|v84)>>(uint(int32(12))%32))&v76)+uint32(_consts[1527]))))
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v91)
						v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v84|v66)>>(uint(int32(6))%32))&v76)+uint32(_consts[1527]))))
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v99)
					}
				}
				return l3
			} else {
				if l0 == int32(1000) {
					v21 = int32(36)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)) = uint8(v21)
					v23 = int32(12580)
					*(*uint16)(unsafe.Add(mBase, uint32(l3))) = uint16(v23)
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
					v28 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)) = uint8(v28)
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v27)>>(uint(int32(2))%32)))+uint32(_consts[1527]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v34)
					v36 = int32(63)
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26&v36)+uint32(_consts[1527]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+3)) = uint8(v39)
					v44 = v25 << (uint(int32(8)) % 32)
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v27<<(uint(int32(16))%32)|v44)>>(uint(int32(12))%32))&v36)+uint32(_consts[1527]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v51)
					v54 = int32(6)
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v44|v26)>>(uint(v54)%32))&v36)+uint32(_consts[1527]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v59)
					if base.Ui32(l2) < base.Ui32(v54) {
					} else {
						if base.Ui32(l4) < base.Ui32(int32(12)) {
						} else {
							v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
							v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
							v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
							v68 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)) = uint8(v68)
							v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v67)>>(uint(int32(2))%32)))+uint32(_consts[1527]))))
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+10)) = uint8(v74)
							v76 = int32(63)
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66&v76)+uint32(_consts[1527]))))
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)) = uint8(v79)
							v84 = v65 << (uint(int32(8)) % 32)
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v67<<(uint(int32(16))%32)|v84)>>(uint(int32(12))%32))&v76)+uint32(_consts[1527]))))
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v91)
							v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v84|v66)>>(uint(int32(6))%32))&v76)+uint32(_consts[1527]))))
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v99)
						}
					}
					return l3
				} else {
					v106 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v106)
					return v106
				}
			}
		}
	}
}
