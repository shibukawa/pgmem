package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int8_avg_serialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v118 int64
	_ = v118
	var v126 int64
	_ = v126
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v187 int64
	_ = v187
	var v190 int64
	_ = v190
	var v193 int64
	_ = v193
	var v196 int64
	_ = v196
	var v197 int64
	_ = v197
	var v201 int64
	_ = v201
	var v208 int64
	_ = v208
	var v220 int32
	_ = v220
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v226 int64
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 == v2 {
		v47 = int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
		switch v22 - int32(429) {
		case 0:
			v47 = int32(1)
		case 1:
			v47 = int32(2)
		default:
			v47 = int32(0)
		}
	}
	if v47 != 0 {
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = int64(0)
		v52 = v15 - int32(-64)
		F_pq_begintypsend(m, v52)
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int32(0)
		} else {
			v57 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
			F_enlargeStringInfo(m, v52, int32(8))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
				v64 = int64(56)
				v66 = int64(65280)
				v68 = int64(40)
				v71 = int64(16711680)
				v73 = int64(24)
				v75 = int64(4278190080)
				v77 = int64(8)
				*(*int64)(unsafe.Add(mBase, uint32(v61+v62))) = v57<<(uint(v64)%64) | v57&v66<<(uint(v68)%64) | (v57&v71<<(uint(v73)%64) | v57&v75<<(uint(v77)%64)) | (int64(base.Ui64(v57)>>(uint(v77)%64))&v75 | int64(base.Ui64(v57)>>(uint(v73)%64))&v71 | (int64(base.Ui64(v57)>>(uint(v68)%64))&v66 | int64(base.Ui64(v57)>>(uint(v64)%64))))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v61 + int32(8)
				v103 = *(*int64)(unsafe.Add(mBase, uint32(v48)+16))
				v104 = *(*int64)(unsafe.Add(mBase, uint32(v48)+24))
				v106 = F_palloc(m, int32(22))
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v106
					v109 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v106))) = uint16(v109)
					*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v106 + int32(2)
					if v104 < int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = int64(16384)
						v118 = int64(0)
						v145 = v118 - v103
						v146 = v118 - (v104 + base.I64_extend_i32_u(base.B2i32(v103 != v118)))
						v150 = int32(0)
						v153 = v106 + int32(22)
						v158 = v145
						v159 = v146
						for {
							v162 = int32(16)
							v163 = v15 + v162
							v166 = m.G0
							v168 = v166 - v162
							m.G0 = v168
							F___udivmodti4(m, v168, v158, v159, int64(10000), int64(0))
							mBase = m.M
							v172 = *(*int64)(unsafe.Add(mBase, uint32(v168)+8))
							v173 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
							*(*int64)(unsafe.Add(mBase, uint32(v163))) = v173
							*(*int64)(unsafe.Add(mBase, uint32(v163)+8)) = v172
							m.G0 = v168 + v162
							v179 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
							v180 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
							v181 = int64(55536)
							v182 = int64(0)
							v187 = int64(32)
							v190 = int64(base.Ui64(v179) >> (uint(v187) % 64))
							v193 = int64(4294967295)
							v196 = v179 & v193
							v197 = v181 * v196
							v201 = int64(base.Ui64(v197)>>(uint(v187)%64)) + v181*v190
							v208 = v196*v182 + v201&v193
							*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v179*v182 + v180*v181 + v182*v190 + int64(base.Ui64(v201)>>(uint(v187)%64)) + int64(base.Ui64(v208)>>(uint(v187)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v15))) = v197&v193 | v208<<(uint(v187)%64)
							v220 = v153 - int32(2)
							v221 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							v222 = v221 + v158
							*(*uint16)(unsafe.Add(mBase, uint32(v220))) = uint16(v222)
							v226 = int64(0)
							v231 = v150 + int32(1)
							if v159 == v226 {
								v232 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v158))
							} else {
								v232 = base.B2i32(v159 != v226)
							}
							if v232 != 0 {
								v150 = v231
								v153 = v220
								v158 = v179
								v159 = v180
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v220
						v234 = v231
						v241 = v150
					} else {
						v126 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v126
						if v103|v104 != v126 {
							v145 = v103
							v146 = v104
							v150 = int32(0)
							v153 = v106 + int32(22)
							v158 = v145
							v159 = v146
							for {
								v162 = int32(16)
								v163 = v15 + v162
								v166 = m.G0
								v168 = v166 - v162
								m.G0 = v168
								F___udivmodti4(m, v168, v158, v159, int64(10000), int64(0))
								mBase = m.M
								v172 = *(*int64)(unsafe.Add(mBase, uint32(v168)+8))
								v173 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
								*(*int64)(unsafe.Add(mBase, uint32(v163))) = v173
								*(*int64)(unsafe.Add(mBase, uint32(v163)+8)) = v172
								m.G0 = v168 + v162
								v179 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
								v180 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
								v181 = int64(55536)
								v182 = int64(0)
								v187 = int64(32)
								v190 = int64(base.Ui64(v179) >> (uint(v187) % 64))
								v193 = int64(4294967295)
								v196 = v179 & v193
								v197 = v181 * v196
								v201 = int64(base.Ui64(v197)>>(uint(v187)%64)) + v181*v190
								v208 = v196*v182 + v201&v193
								*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v179*v182 + v180*v181 + v182*v190 + int64(base.Ui64(v201)>>(uint(v187)%64)) + int64(base.Ui64(v208)>>(uint(v187)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v15))) = v197&v193 | v208<<(uint(v187)%64)
								v220 = v153 - int32(2)
								v221 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
								v222 = v221 + v158
								*(*uint16)(unsafe.Add(mBase, uint32(v220))) = uint16(v222)
								v226 = int64(0)
								v231 = v150 + int32(1)
								if v159 == v226 {
									v232 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v158))
								} else {
									v232 = base.B2i32(v159 != v226)
								}
								if v232 != 0 {
									v150 = v231
									v153 = v220
									v158 = v179
									v159 = v180
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v220
							v234 = v231
							v241 = v150
						} else {
							v234 = int32(0)
							v241 = v2
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v241
					*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v234
					v249 = v15 - int32(-64)
					F_numericvar_serialize(m, v249, v15+int32(40))
					mBase = m.M
					v253 = m.ExcPending
					if v253 != 0 {
						return int32(0)
					} else {
						v255 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
						v256 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v255))) = v256 << (uint(int32(2)) % 32)
						F_pfree(m, v106)
						mBase = m.M
						v261 = m.ExcPending
						if v261 != 0 {
							return int32(0)
						} else {
							m.G0 = v15 + int32(80)
							return v255
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v135 = m.ExcPending
		if v135 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_int8_avg_serialize_0), int32(0))
			mBase = m.M
			v139 = m.ExcPending
			if v139 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_int8_avg_serialize_1), int32(_a_F_int8_avg_serialize_2), int32(_a_F_int8_avg_serialize_3))
				mBase = m.M
				v144 = m.ExcPending
				if v144 != 0 {
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
func F_int8_decrement(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if v5 == int64(-9223372036854775807-1) {
		v8 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v8)
		return int32(0)
	} else {
		v12 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v12)
		v16 = F_Int64GetDatum(m, v5-int64(1))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			return v16
		}
	}
}
