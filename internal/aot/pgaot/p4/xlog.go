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
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v61 int32
	_ = v61
	var v68 int64
	_ = v68
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v99 int64
	_ = v99
	var v102 int64
	_ = v102
	var v108 int32
	_ = v108
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
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int64
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
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
			v32 = int32(_a_F_xlog_desc_0)
		} else {
			v32 = int32(_a_F_xlog_desc_1)
		}
		v33 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
		if base.Ui32(v40) <= base.Ui32(int32(2)) {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v40<<(uint(int32(2))%32))+uint32(_c_F_xlog_desc[0])))
			v46 = v45
		} else {
			v46 = int32(_a_F_xlog_desc_2)
		}
		v47 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
		v48 = *(*int64)(unsafe.Add(mBase, uint32(v26)+40))
		v49 = *(*int64)(unsafe.Add(mBase, uint32(v26)+48))
		v50 = *(*int64)(unsafe.Add(mBase, uint32(v26)+76))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
		v52 = *(*int32)(unsafe.Add(mBase, uint32(v26)+56))
		v53 = *(*int64)(unsafe.Add(mBase, uint32(v26)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v32
		*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = v46
		*(*uint32)(unsafe.Add(mBase, uint32(v20)+124)) = uint32(v53)
		*(*int32)(unsafe.Add(mBase, uint32(v20)+152)) = v52
		*(*int32)(unsafe.Add(mBase, uint32(v20)+156)) = v51
		if v25 != 0 {
			v61 = int32(_a_F_xlog_desc_3)
		} else {
			v61 = int32(_a_F_xlog_desc_4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v20)+168)) = v61
		*(*int64)(unsafe.Add(mBase, uint32(v20)+160)) = v50
		*(*int64)(unsafe.Add(mBase, uint32(v20)+144)) = v49
		*(*int64)(unsafe.Add(mBase, uint32(v20)+136)) = v48
		*(*int64)(unsafe.Add(mBase, uint32(v20)+128)) = v47
		v68 = int64(base.Ui64(v53) >> (uint(int64(32)) % 64))
		*(*uint32)(unsafe.Add(mBase, uint32(v20)+120)) = uint32(v68)
		*(*uint32)(unsafe.Add(mBase, uint32(v20)+100)) = uint32(v33)
		*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v38
		*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v37
		*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = base.I32_wrap_i64(int64(base.Ui64(v33) >> (uint(int64(32)) % 64)))
		F_appendStringInfo(m, l0, int32(_a_F_xlog_desc_5), v20+int32(96))
		mBase = m.M
		v78 = m.ExcPending
		if v78 != 0 {
			return
		} else {
			m.G0 = v20 + int32(192)
			return
		}
	default:
		v90 = base.I32_extend8_s(v23)
		if v90&int32(-32) == int32(-96) {
			m.G0 = v20 + int32(192)
			return
		} else {
			v96 = v25 - int32(80)
			if v96 != 0 {
				if v96 == int32(16) {
					v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+25)))
					v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)))
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
					v113 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
					v114 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
					v115 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
					v116 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
					if base.Ui32(v116) <= base.Ui32(int32(2)) {
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v116<<(uint(int32(2))%32))+uint32(_c_F_xlog_desc[0])))
						v123 = v121
					} else {
						v123 = int32(_a_F_xlog_desc_2)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v123
					*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v111
					if v109&int32(1) != 0 {
						v130 = int32(_a_F_xlog_desc_6)
					} else {
						v130 = int32(_a_F_xlog_desc_7)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v130
					if v110&int32(1) != 0 {
						v136 = int32(_a_F_xlog_desc_6)
					} else {
						v136 = int32(_a_F_xlog_desc_7)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = v136
					*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v112
					*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v113
					*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v114
					*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v115
					F_appendStringInfo(m, l0, int32(_a_F_xlog_desc_8), v20-int32(-64))
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return
					} else {
						m.G0 = v20 + int32(192)
						return
					}
				} else {
					if v90 <= int32(-113) {
						v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
						if v151&int32(1) != 0 {
							v154 = int32(_a_F_xlog_desc_0)
						} else {
							v154 = int32(_a_F_xlog_desc_1)
						}
						F_appendStringInfoString(m, l0, v154)
						mBase = m.M
						v156 = m.ExcPending
						if v156 != 0 {
							return
						} else {
							m.G0 = v20 + int32(192)
							return
						}
					} else {
						switch v25 - int32(208) {
						case 0:
							v181 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
							v182 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
							v183 = F_timestamptz_to_str(m, v182)
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v183
								*(*uint32)(unsafe.Add(mBase, uint32(v20)+20)) = uint32(v181)
								v188 = int64(base.Ui64(v181) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v20)+16)) = uint32(v188)
								F_appendStringInfo(m, l0, int32(_a_F_xlog_desc_9), v20+int32(16))
								mBase = m.M
								v194 = m.ExcPending
								if v194 != 0 {
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
							v195 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
							if base.Ui32(v195) <= base.Ui32(int32(2)) {
								v200 = *(*int32)(unsafe.Add(mBase, uint32(v195<<(uint(int32(2))%32))+uint32(_c_F_xlog_desc[0])))
								v202 = v200
							} else {
								v202 = int32(_a_F_xlog_desc_2)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v202
							F_appendStringInfo(m, l0, int32(_a_F_xlog_desc_10), v20+int32(32))
							mBase = m.M
							v208 = m.ExcPending
							if v208 != 0 {
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
								v161 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
								v162 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
								v163 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
								v164 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
								v165 = F_timestamptz_to_str(m, v164)
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return
								} else {
									if base.Ui32(v161) <= base.Ui32(int32(2)) {
										v171 = *(*int32)(unsafe.Add(mBase, uint32(v161<<(uint(int32(2))%32))+uint32(_c_F_xlog_desc[0])))
										v173 = v171
									} else {
										v173 = int32(_a_F_xlog_desc_2)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v173
									*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v165
									*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v162
									*(*int32)(unsafe.Add(mBase, uint32(v20))) = v163
									F_appendStringInfo(m, l0, int32(_a_F_xlog_desc_11), v20)
									mBase = m.M
									v180 = m.ExcPending
									if v180 != 0 {
										return
									} else {
										m.G0 = v20 + int32(192)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v99 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
				*(*uint32)(unsafe.Add(mBase, uint32(v20)+52)) = uint32(v99)
				v102 = int64(base.Ui64(v99) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v20)+48)) = uint32(v102)
				F_appendStringInfo(m, l0, int32(_a_F_xlog_desc_12), v20+int32(48))
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
					return
				} else {
					m.G0 = v20 + int32(192)
					return
				}
			}
		}
	case 3:
		v79 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = v79
		F_appendStringInfo(m, l0, int32(_a_F_xlog_desc_13), v20+int32(176))
		mBase = m.M
		v85 = m.ExcPending
		if v85 != 0 {
			return
		} else {
			m.G0 = v20 + int32(192)
			return
		}
	case 7:
		F_appendStringInfoString(m, l0, v26+int32(8))
		mBase = m.M
		v89 = m.ExcPending
		if v89 != 0 {
			return
		} else {
			m.G0 = v20 + int32(192)
			return
		}
	}
}
