package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_xlog_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int64
	_ = v187
	var v188 int64
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int64
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	v19 = m.G0
	v21 = v19 - int32(208)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+48)))
	v26 = v24 & int32(240)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	switch int32(base.Ui32(v24) >> (uint(int32(4)) % 32)) {
	case 0, 1:
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+16)))
		if v32 != 0 {
			v33 = int32(_a_F_xlog_desc_0)
		} else {
			v33 = int32(_a_F_xlog_desc_1)
		}
		v34 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
		if base.Ui32(v41) <= base.Ui32(int32(2)) {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v41<<(uint(int32(2))%32))+uint32(_c_F_xlog_desc[0])))
			v47 = v46
		} else {
			v47 = int32(_a_F_xlog_desc_2)
		}
		v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+24)))
		v49 = *(*int64)(unsafe.Add(mBase, uint32(v27)+40))
		v50 = *(*int64)(unsafe.Add(mBase, uint32(v27)+56))
		v51 = *(*int64)(unsafe.Add(mBase, uint32(v27)+64))
		v52 = *(*int64)(unsafe.Add(mBase, uint32(v27)+80))
		v53 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
		v54 = *(*int64)(unsafe.Add(mBase, uint32(v27)+48))
		v55 = *(*int64)(unsafe.Add(mBase, uint32(v27)+32))
		*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v33
		*(*int32)(unsafe.Add(mBase, uint32(v21)+116)) = v47
		*(*uint32)(unsafe.Add(mBase, uint32(v21)+128)) = uint32(v55)
		*(*int64)(unsafe.Add(mBase, uint32(v21)+144)) = v54
		if v26 != 0 {
			v62 = int32(_a_F_xlog_desc_3)
		} else {
			v62 = int32(_a_F_xlog_desc_4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v62
		*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v53
		*(*int64)(unsafe.Add(mBase, uint32(v21)+168)) = v52
		*(*int64)(unsafe.Add(mBase, uint32(v21)+160)) = v51
		*(*int64)(unsafe.Add(mBase, uint32(v21)+152)) = v50
		*(*int64)(unsafe.Add(mBase, uint32(v21)+132)) = v49
		if v48 != 0 {
			v71 = int32(_a_F_xlog_desc_0)
		} else {
			v71 = int32(_a_F_xlog_desc_1)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v21)+120)) = v71
		v74 = int64(base.Ui64(v55) >> (uint(int64(32)) % 64))
		*(*uint32)(unsafe.Add(mBase, uint32(v21)+124)) = uint32(v74)
		*(*uint32)(unsafe.Add(mBase, uint32(v21)+100)) = uint32(v34)
		*(*int32)(unsafe.Add(mBase, uint32(v21)+104)) = v39
		*(*int32)(unsafe.Add(mBase, uint32(v21)+108)) = v38
		*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = base.I32_wrap_i64(int64(base.Ui64(v34) >> (uint(int64(32)) % 64)))
		F_appendStringInfo(m, l0, int32(_a_F_xlog_desc_5), v21+int32(96))
		mBase = m.M
		v84 = m.ExcPending
		if v84 != 0 {
			return
		} else {
			m.G0 = v21 + int32(208)
			return
		}
	default:
		if v24&int32(224) == int32(160) {
			m.G0 = v21 + int32(208)
			return
		} else {
			v101 = v26 - int32(80)
			if v101 != 0 {
				if v101 == int32(16) {
					v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+25)))
					v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+24)))
					v116 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
					v118 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
					v119 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					v121 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
					if base.Ui32(v121) <= base.Ui32(int32(2)) {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v121<<(uint(int32(2))%32))+uint32(_c_F_xlog_desc[0])))
						v128 = v126
					} else {
						v128 = int32(_a_F_xlog_desc_2)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v128
					*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v116
					if v114&int32(1) != 0 {
						v135 = int32(_a_F_xlog_desc_6)
					} else {
						v135 = int32(_a_F_xlog_desc_7)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v21)+92)) = v135
					if v115&int32(1) != 0 {
						v141 = int32(_a_F_xlog_desc_6)
					} else {
						v141 = int32(_a_F_xlog_desc_7)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = v141
					*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v117
					*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v118
					*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v119
					*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v120
					F_appendStringInfo(m, l0, int32(_a_F_xlog_desc_8), v21-int32(-64))
					mBase = m.M
					v151 = m.ExcPending
					if v151 != 0 {
						return
					} else {
						m.G0 = v21 + int32(208)
						return
					}
				} else {
					if base.I32_extend8_s(v24) <= int32(-113) {
						v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
						if v157&int32(1) != 0 {
							v160 = int32(_a_F_xlog_desc_0)
						} else {
							v160 = int32(_a_F_xlog_desc_1)
						}
						F_appendStringInfoString(m, l0, v160)
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return
						} else {
							m.G0 = v21 + int32(208)
							return
						}
					} else {
						switch v26 - int32(208) {
						case 0:
							v187 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
							v188 = *(*int64)(unsafe.Add(mBase, uint32(v27)+8))
							v189 = F_timestamptz_to_str(m, v188)
							mBase = m.M
							v190 = m.ExcPending
							if v190 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v189
								*(*uint32)(unsafe.Add(mBase, uint32(v21)+20)) = uint32(v187)
								v194 = int64(base.Ui64(v187) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v21)+16)) = uint32(v194)
								F_appendStringInfo(m, l0, int32(_a_F_xlog_desc_9), v21+int32(16))
								mBase = m.M
								v200 = m.ExcPending
								if v200 != 0 {
									return
								} else {
									m.G0 = v21 + int32(208)
									return
								}
							}
						case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
							if base.Ui32(v24) < base.Ui32(int32(240)) {
								m.G0 = v21 + int32(208)
								return
							} else {
								v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
								if v219&int32(1) != 0 {
									v222 = int32(_a_F_xlog_desc_0)
								} else {
									v222 = int32(_a_F_xlog_desc_1)
								}
								F_appendStringInfoString(m, l0, v222)
								mBase = m.M
								v224 = m.ExcPending
								if v224 != 0 {
									return
								} else {
									m.G0 = v21 + int32(208)
									return
								}
							}
						case 16:
							v201 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
							if base.Ui32(v201) <= base.Ui32(int32(2)) {
								v206 = *(*int32)(unsafe.Add(mBase, uint32(v201<<(uint(int32(2))%32))+uint32(_c_F_xlog_desc[0])))
								v208 = v206
							} else {
								v208 = int32(_a_F_xlog_desc_2)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v208
							F_appendStringInfo(m, l0, int32(_a_F_xlog_desc_10), v21+int32(32))
							mBase = m.M
							v214 = m.ExcPending
							if v214 != 0 {
								return
							} else {
								m.G0 = v21 + int32(208)
								return
							}
						default:
							if v26 != int32(144) {
								if base.Ui32(v24) < base.Ui32(int32(240)) {
									m.G0 = v21 + int32(208)
									return
								} else {
									v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
									if v219&int32(1) != 0 {
										v222 = int32(_a_F_xlog_desc_0)
									} else {
										v222 = int32(_a_F_xlog_desc_1)
									}
									F_appendStringInfoString(m, l0, v222)
									mBase = m.M
									v224 = m.ExcPending
									if v224 != 0 {
										return
									} else {
										m.G0 = v21 + int32(208)
										return
									}
								}
							} else {
								v167 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
								v168 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
								v169 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
								v170 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
								v171 = F_timestamptz_to_str(m, v170)
								mBase = m.M
								v172 = m.ExcPending
								if v172 != 0 {
									return
								} else {
									if base.Ui32(v167) <= base.Ui32(int32(2)) {
										v177 = *(*int32)(unsafe.Add(mBase, uint32(v167<<(uint(int32(2))%32))+uint32(_c_F_xlog_desc[0])))
										v179 = v177
									} else {
										v179 = int32(_a_F_xlog_desc_2)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v179
									*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v171
									*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v168
									*(*int32)(unsafe.Add(mBase, uint32(v21))) = v169
									F_appendStringInfo(m, l0, int32(_a_F_xlog_desc_11), v21)
									mBase = m.M
									v186 = m.ExcPending
									if v186 != 0 {
										return
									} else {
										m.G0 = v21 + int32(208)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v104 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
				*(*uint32)(unsafe.Add(mBase, uint32(v21)+52)) = uint32(v104)
				v107 = int64(base.Ui64(v104) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v21)+48)) = uint32(v107)
				F_appendStringInfo(m, l0, int32(_a_F_xlog_desc_12), v21+int32(48))
				mBase = m.M
				v113 = m.ExcPending
				if v113 != 0 {
					return
				} else {
					m.G0 = v21 + int32(208)
					return
				}
			}
		}
	case 3:
		v85 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
		*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v85
		F_appendStringInfo(m, l0, int32(_a_F_xlog_desc_13), v21+int32(192))
		mBase = m.M
		v91 = m.ExcPending
		if v91 != 0 {
			return
		} else {
			m.G0 = v21 + int32(208)
			return
		}
	case 7:
		F_appendStringInfoString(m, l0, v27+int32(8))
		mBase = m.M
		v95 = m.ExcPending
		if v95 != 0 {
			return
		} else {
			m.G0 = v21 + int32(208)
			return
		}
	}
}
