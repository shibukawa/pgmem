package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__crypt_gensalt_blowfish_rn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	if base.B2i32(l2 < int32(16))|base.B2i32(l4 < int32(30)) == int32(0) {
		if l0 != 0 {
			if base.Ui32(l0-int32(32)) < base.Ui32(int32(-28)) {
				v243 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v243)
				return v243
			} else {
				v22 = l0
				v23 = int32(36)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v23)
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(610349604)
				v29 = int32(10)
				v30 = base.I32_div_u_s(v22&int32(255), v29)
				v31 = int32(48)
				v32 = v30 + v31
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v32)
				v38 = v22 - v30*v29 | v31
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v38)
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				v41 = int32(2)
				v43 = int32(_a_F__crypt_gensalt_blowfish_rn_0)
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v40)>>(uint(v41)%32)))+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)) = uint8(v45)
				v47 = int32(4)
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
				v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40<<(uint(v47)%32)&v31+v43+int32(base.Ui32(v53)>>(uint(v47)%32))))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v57)
				v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
				v60 = int32(63)
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59&v60)+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+10)) = uint8(v64)
				v68 = int32(60)
				v72 = int32(6)
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53<<(uint(v41)%32)&v68+v43+int32(base.Ui32(v59)>>(uint(v72)%32))))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v75)
				v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
				v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v77)>>(uint(v41)%32)))+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)) = uint8(v82)
				v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
				v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77<<(uint(v47)%32)&v31+v43+int32(base.Ui32(v90)>>(uint(v47)%32))))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)) = uint8(v94)
				v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
				v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96&v60)+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)) = uint8(v101)
				v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90<<(uint(v41)%32)&v68+v43+int32(base.Ui32(v96)>>(uint(v72)%32))))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)) = uint8(v112)
				v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
				v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v114)>>(uint(v41)%32)))+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+15)) = uint8(v119)
				v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
				v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114<<(uint(v47)%32)&v31+v43+int32(base.Ui32(v127)>>(uint(v47)%32))))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)) = uint8(v131)
				v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
				v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133&v60)+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+18)) = uint8(v138)
				v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127<<(uint(v41)%32)&v68+v43+int32(base.Ui32(v133)>>(uint(v72)%32))))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+17)) = uint8(v149)
				v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
				v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v151)>>(uint(v41)%32)))+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+19)) = uint8(v156)
				v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
				v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151<<(uint(v47)%32)&v31+v43+int32(base.Ui32(v164)>>(uint(v47)%32))))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)) = uint8(v168)
				v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
				v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170&v60)+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+22)) = uint8(v175)
				v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164<<(uint(v41)%32)&v68+v43+int32(base.Ui32(v170)>>(uint(v72)%32))))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)) = uint8(v186)
				v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
				v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v188)>>(uint(v41)%32)))+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+23)) = uint8(v193)
				v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
				v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188<<(uint(v47)%32)&v31+v43+int32(base.Ui32(v201)>>(uint(v47)%32))))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+24)) = uint8(v205)
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
				v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207&v60)+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+26)) = uint8(v212)
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201<<(uint(v41)%32)&v68+v43+int32(base.Ui32(v207)>>(uint(v72)%32))))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+25)) = uint8(v223)
				v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)))
				v226 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v226)
				v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v225)>>(uint(v41)%32)))+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+27)) = uint8(v232)
				v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225<<(uint(v47)%32)&v31)+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)) = uint8(v240)
				return l3
			}
		} else {
			v22 = int32(5)
			v23 = int32(36)
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v23)
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(610349604)
			v29 = int32(10)
			v30 = base.I32_div_u_s(v22&int32(255), v29)
			v31 = int32(48)
			v32 = v30 + v31
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v32)
			v38 = v22 - v30*v29 | v31
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v38)
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v41 = int32(2)
			v43 = int32(_a_F__crypt_gensalt_blowfish_rn_0)
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v40)>>(uint(v41)%32)))+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)) = uint8(v45)
			v47 = int32(4)
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40<<(uint(v47)%32)&v31+v43+int32(base.Ui32(v53)>>(uint(v47)%32))))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v57)
			v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			v60 = int32(63)
			v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59&v60)+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+10)) = uint8(v64)
			v68 = int32(60)
			v72 = int32(6)
			v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53<<(uint(v41)%32)&v68+v43+int32(base.Ui32(v59)>>(uint(v72)%32))))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v75)
			v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v77)>>(uint(v41)%32)))+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)) = uint8(v82)
			v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
			v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77<<(uint(v47)%32)&v31+v43+int32(base.Ui32(v90)>>(uint(v47)%32))))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)) = uint8(v94)
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96&v60)+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)) = uint8(v101)
			v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90<<(uint(v41)%32)&v68+v43+int32(base.Ui32(v96)>>(uint(v72)%32))))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)) = uint8(v112)
			v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
			v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v114)>>(uint(v41)%32)))+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+15)) = uint8(v119)
			v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
			v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114<<(uint(v47)%32)&v31+v43+int32(base.Ui32(v127)>>(uint(v47)%32))))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)) = uint8(v131)
			v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
			v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133&v60)+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+18)) = uint8(v138)
			v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127<<(uint(v41)%32)&v68+v43+int32(base.Ui32(v133)>>(uint(v72)%32))))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+17)) = uint8(v149)
			v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
			v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v151)>>(uint(v41)%32)))+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+19)) = uint8(v156)
			v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
			v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151<<(uint(v47)%32)&v31+v43+int32(base.Ui32(v164)>>(uint(v47)%32))))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)) = uint8(v168)
			v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
			v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170&v60)+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+22)) = uint8(v175)
			v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164<<(uint(v41)%32)&v68+v43+int32(base.Ui32(v170)>>(uint(v72)%32))))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)) = uint8(v186)
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v188)>>(uint(v41)%32)))+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+23)) = uint8(v193)
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
			v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188<<(uint(v47)%32)&v31+v43+int32(base.Ui32(v201)>>(uint(v47)%32))))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+24)) = uint8(v205)
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207&v60)+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+26)) = uint8(v212)
			v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201<<(uint(v41)%32)&v68+v43+int32(base.Ui32(v207)>>(uint(v72)%32))))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+25)) = uint8(v223)
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)))
			v226 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v226)
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v225)>>(uint(v41)%32)))+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+27)) = uint8(v232)
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225<<(uint(v47)%32)&v31)+uint32(_c_F__crypt_gensalt_blowfish_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)) = uint8(v240)
			return l3
		}
	} else {
		if int32(0) < l4 {
			v243 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v243)
			return v243
		} else {
			return int32(0)
		}
	}
}
func F__crypt_gensalt_md5_rn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	if base.B2i32(l2 < int32(3))|base.B2i32(l4 < int32(8)) == int32(0) {
		if base.B2i32(l0 == int32(0))|base.B2i32(l0 == int32(1000)) != 0 {
			v24 = int32(36)
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)) = uint8(v24)
			v26 = int32(_a_F__crypt_gensalt_md5_rn_0)
			*(*uint16)(unsafe.Add(mBase, uint32(l3))) = uint16(v26)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			v31 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)) = uint8(v31)
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v30)>>(uint(int32(2))%32)))+uint32(_c_F__crypt_gensalt_md5_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v37)
			v39 = int32(63)
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29&v39)+uint32(_c_F__crypt_gensalt_md5_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+3)) = uint8(v43)
			v46 = v28 << (uint(int32(8)) % 32)
			v50 = int32(12)
			v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v46|v30<<(uint(int32(16))%32))>>(uint(v50)%32))&v39)+uint32(_c_F__crypt_gensalt_md5_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v56)
			v59 = int32(6)
			v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v29|v46)>>(uint(v59)%32))&v39)+uint32(_c_F__crypt_gensalt_md5_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v65)
			if base.B2i32(base.Ui32(l2) < base.Ui32(v59))|base.B2i32(base.Ui32(l4) < base.Ui32(v50)) == v31 {
				v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
				v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
				v77 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)) = uint8(v77)
				v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v76)>>(uint(int32(2))%32)))+uint32(_c_F__crypt_gensalt_md5_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+10)) = uint8(v83)
				v85 = int32(63)
				v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75&v85)+uint32(_c_F__crypt_gensalt_md5_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)) = uint8(v89)
				v92 = v74 << (uint(int32(8)) % 32)
				v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v92|v76<<(uint(int32(16))%32))>>(uint(int32(12))%32))&v85)+uint32(_c_F__crypt_gensalt_md5_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v102)
				v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v75|v92)>>(uint(int32(6))%32))&v85)+uint32(_c_F__crypt_gensalt_md5_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v111)
			} else {
			}
			return l3
		} else {
			v117 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v117)
			return v117
		}
	} else {
		if int32(0) < l4 {
			v117 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v117)
			return v117
		} else {
			return int32(0)
		}
	}
}
