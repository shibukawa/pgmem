package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_xlog_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v64 int32
	_ = v64
	var v71 int64
	_ = v71
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int64
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int64
	_ = v193
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	v18 = m.G0
	v20 = v18 - int32(192)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+48)))
	v25 = v23 & int32(240)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	switch int32(base.Ui32(v23) >> (uint(int32(4)) % 32)) {
	case 0, 1:
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+16)))
		if v31 != 0 {
			v32 = int32(361051)
		} else {
			v32 = int32(378570)
		}
		v33 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
		if base.Ui32(v41) <= base.Ui32(int32(2)) {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v41<<(uint(int32(2))%32))+uint32(_consts[138])))
			v49 = v48
		} else {
			v49 = int32(571405)
		}
		v50 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
		v51 = *(*int64)(unsafe.Add(mBase, uint32(v26)+40))
		v52 = *(*int64)(unsafe.Add(mBase, uint32(v26)+48))
		v53 = *(*int64)(unsafe.Add(mBase, uint32(v26)+76))
		v54 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v26)+56))
		v56 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v32
		*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = v49
		*(*uint32)(unsafe.Add(mBase, uint32(v20)+124)) = uint32(v56)
		*(*int32)(unsafe.Add(mBase, uint32(v20)+152)) = v55
		*(*int32)(unsafe.Add(mBase, uint32(v20)+156)) = v54
		if v25 != 0 {
			v64 = int32(390850)
		} else {
			v64 = int32(255897)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v20)+168)) = v64
		*(*int64)(unsafe.Add(mBase, uint32(v20)+160)) = v53
		*(*int64)(unsafe.Add(mBase, uint32(v20)+144)) = v52
		*(*int64)(unsafe.Add(mBase, uint32(v20)+136)) = v51
		*(*int64)(unsafe.Add(mBase, uint32(v20)+128)) = v50
		v71 = int64(base.Ui64(v56) >> (uint(int64(32)) % 64))
		*(*uint32)(unsafe.Add(mBase, uint32(v20)+120)) = uint32(v71)
		*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = base.I32_wrap_i64(v33)
		*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v38
		*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v37
		*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = base.I32_wrap_i64(int64(base.Ui64(v33) >> (uint(int64(32)) % 64)))
		F_appendStringInfo(m, l0, int32(208999), v20+int32(96))
		mBase = m.M
		v81 = m.ExcPending
		if v81 != 0 {
			return
		} else {
			m.G0 = v20 + int32(192)
			return
		}
	default:
		v93 = base.I32_extend8_s(v23)
		if v93&int32(-32) == int32(-96) {
			m.G0 = v20 + int32(192)
			return
		} else {
			switch v25 - int32(80) {
			case 0:
				v100 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
				*(*uint32)(unsafe.Add(mBase, uint32(v20)+52)) = uint32(v100)
				v103 = int64(base.Ui64(v100) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v20)+48)) = uint32(v103)
				F_appendStringInfo(m, l0, int32(540555), v20+int32(48))
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return
				} else {
					m.G0 = v20 + int32(192)
					return
				}
			default:
				if v93 <= int32(-113) {
					v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
					if v154&int32(1) != 0 {
						v157 = int32(361051)
					} else {
						v157 = int32(378570)
					}
					F_appendStringInfoString(m, l0, v157)
					mBase = m.M
					v159 = m.ExcPending
					if v159 != 0 {
						return
					} else {
						m.G0 = v20 + int32(192)
						return
					}
				} else {
					switch v25 - int32(208) {
					case 0:
						v186 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
						v187 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
						v188 = F_timestamptz_to_str(m, v187)
						mBase = m.M
						v189 = m.ExcPending
						if v189 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v188
							*(*uint32)(unsafe.Add(mBase, uint32(v20)+20)) = uint32(v186)
							v193 = int64(base.Ui64(v186) >> (uint(int64(32)) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v20)+16)) = uint32(v193)
							F_appendStringInfo(m, l0, int32(204369), v20+int32(16))
							mBase = m.M
							v199 = m.ExcPending
							if v199 != 0 {
								return
							} else {
								m.G0 = v20 + int32(192)
								return
							}
						}
					case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
						m.G0 = v20 + int32(192)
						return
					case 16:
						v201 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
						if base.Ui32(v201) <= base.Ui32(int32(2)) {
							v208 = *(*int32)(unsafe.Add(mBase, uint32(v201<<(uint(int32(2))%32))+uint32(_consts[138])))
							v209 = v208
						} else {
							v209 = int32(571405)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v209
						F_appendStringInfo(m, l0, int32(196145), v20+int32(32))
						mBase = m.M
						v215 = m.ExcPending
						if v215 != 0 {
							return
						} else {
							m.G0 = v20 + int32(192)
							return
						}
					default:
						if v25 != int32(144) {
							m.G0 = v20 + int32(192)
							return
						} else {
							v164 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
							v165 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
							v166 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
							v168 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
							v169 = F_timestamptz_to_str(m, v168)
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
								return
							} else {
								if base.Ui32(v164) <= base.Ui32(int32(2)) {
									v177 = *(*int32)(unsafe.Add(mBase, uint32(v164<<(uint(int32(2))%32))+uint32(_consts[138])))
									v178 = v177
								} else {
									v178 = int32(571405)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v178
								*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v169
								*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v165
								*(*int32)(unsafe.Add(mBase, uint32(v20))) = v166
								F_appendStringInfo(m, l0, int32(196115), v20)
								mBase = m.M
								v185 = m.ExcPending
								if v185 != 0 {
									return
								} else {
									m.G0 = v20 + int32(192)
									return
								}
							}
						}
					}
				}
			case 16:
				v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+25)))
				v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)))
				v112 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
				v114 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
				v115 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				v118 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
				if base.Ui32(v118) <= base.Ui32(int32(2)) {
					v125 = *(*int32)(unsafe.Add(mBase, uint32(v118<<(uint(int32(2))%32))+uint32(_consts[138])))
					v126 = v125
				} else {
					v126 = int32(571405)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v126
				*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v112
				if v110&int32(1) != 0 {
					v133 = int32(285807)
				} else {
					v133 = int32(355204)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v133
				if v111&int32(1) != 0 {
					v139 = int32(285807)
				} else {
					v139 = int32(355204)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = v139
				*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v113
				*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v114
				*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v115
				*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v116
				F_appendStringInfo(m, l0, int32(186735), v20-int32(-64))
				mBase = m.M
				v149 = m.ExcPending
				if v149 != 0 {
					return
				} else {
					m.G0 = v20 + int32(192)
					return
				}
			}
		}
	case 3:
		v82 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = v82
		F_appendStringInfo(m, l0, int32(64823), v20+int32(176))
		mBase = m.M
		v88 = m.ExcPending
		if v88 != 0 {
			return
		} else {
			m.G0 = v20 + int32(192)
			return
		}
	case 7:
		F_appendStringInfoString(m, l0, v26+int32(8))
		mBase = m.M
		v92 = m.ExcPending
		if v92 != 0 {
			return
		} else {
			m.G0 = v20 + int32(192)
			return
		}
	}
}
