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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v121 int64
	_ = v121
	var v129 int64
	_ = v129
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v186 int32
	_ = v186
	var v187 int64
	_ = v187
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v195 int64
	_ = v195
	var v198 int64
	_ = v198
	var v201 int64
	_ = v201
	var v204 int64
	_ = v204
	var v205 int64
	_ = v205
	var v209 int64
	_ = v209
	var v216 int64
	_ = v216
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v230 int64
	_ = v230
	var v234 int64
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 == v2 {
		v48 = int32(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		switch v23 - int32(429) {
		case 0:
			v48 = int32(1)
		case 1:
			v48 = int32(2)
		default:
			v48 = int32(0)
		}
	}
	if v48 != 0 {
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = int64(0)
		F_pq_begintypsend(m, v16-int32(-64))
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int32(0)
		} else {
			v58 = *(*int64)(unsafe.Add(mBase, uint32(v49)+8))
			F_enlargeStringInfo(m, v16-int32(-64), int32(8))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int32(0)
			} else {
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
				v67 = int64(56)
				v69 = int64(65280)
				v71 = int64(40)
				v74 = int64(16711680)
				v76 = int64(24)
				v78 = int64(4278190080)
				v80 = int64(8)
				*(*int64)(unsafe.Add(mBase, uint32(v64+v65))) = v58<<(uint(v67)%64) | v58&v69<<(uint(v71)%64) | (v58&v74<<(uint(v76)%64) | v58&v78<<(uint(v80)%64)) | (int64(base.Ui64(v58)>>(uint(v80)%64))&v78 | int64(base.Ui64(v58)>>(uint(v76)%64))&v74 | (int64(base.Ui64(v58)>>(uint(v71)%64))&v69 | int64(base.Ui64(v58)>>(uint(v67)%64))))
				*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v64 + int32(8)
				v106 = *(*int64)(unsafe.Add(mBase, uint32(v49)+24))
				v107 = *(*int64)(unsafe.Add(mBase, uint32(v49)+16))
				v109 = F_palloc(m, int32(22))
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v109
					v112 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v109))) = uint16(v112)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v109 + int32(2)
					if v106 < int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = int64(16384)
						v121 = int64(0)
						v148 = v121 - (v106 + base.I64_extend_i32_u(base.B2i32(v107 != v121)))
						v149 = v121 - v107
						v155 = int32(0)
						v157 = v109 + int32(22)
						v164 = v148
						v165 = v149
						for {
							v169 = v16 + int32(24)
							v172 = m.G0
							v173 = int32(16)
							v174 = v172 - v173
							m.G0 = v174
							F___udivmodti4(m, v174, v165, v164, int64(10000), int64(0))
							mBase = m.M
							v178 = *(*int64)(unsafe.Add(mBase, uint32(v174)))
							v179 = *(*int64)(unsafe.Add(mBase, uint32(v174)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v169)+8)) = v179
							*(*int64)(unsafe.Add(mBase, uint32(v169))) = v178
							m.G0 = v174 + v173
							v186 = v16 + int32(8)
							v187 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
							v188 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(32))))
							v189 = int64(55536)
							v190 = int64(0)
							v195 = int64(32)
							v198 = int64(base.Ui64(v187) >> (uint(v195) % 64))
							v201 = int64(4294967295)
							v204 = v187 & v201
							v205 = v189 * v204
							v209 = int64(base.Ui64(v205)>>(uint(v195)%64)) + v189*v198
							v216 = v204*v190 + v209&v201
							*(*int64)(unsafe.Add(mBase, uint32(v186)+8)) = v187*v190 + v188*v189 + v190*v198 + int64(base.Ui64(v209)>>(uint(v195)%64)) + int64(base.Ui64(v216)>>(uint(v195)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v186))) = v205&v201 | v216<<(uint(v195)%64)
							v228 = v157 - int32(2)
							v229 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
							v230 = v229 + v165
							*(*uint16)(unsafe.Add(mBase, uint32(v228))) = uint16(v230)
							v234 = int64(0)
							v239 = v155 + int32(1)
							if v164 == v234 {
								v240 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v165))
							} else {
								v240 = base.B2i32(v164 != v234)
							}
							if v240 != 0 {
								v155 = v239
								v157 = v228
								v164 = v188
								v165 = v187
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v228
						v242 = v239
						v246 = v155
					} else {
						v129 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v129
						if v106|v107 != v129 {
							v148 = v106
							v149 = v107
							v155 = int32(0)
							v157 = v109 + int32(22)
							v164 = v148
							v165 = v149
							for {
								v169 = v16 + int32(24)
								v172 = m.G0
								v173 = int32(16)
								v174 = v172 - v173
								m.G0 = v174
								F___udivmodti4(m, v174, v165, v164, int64(10000), int64(0))
								mBase = m.M
								v178 = *(*int64)(unsafe.Add(mBase, uint32(v174)))
								v179 = *(*int64)(unsafe.Add(mBase, uint32(v174)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v169)+8)) = v179
								*(*int64)(unsafe.Add(mBase, uint32(v169))) = v178
								m.G0 = v174 + v173
								v186 = v16 + int32(8)
								v187 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
								v188 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(32))))
								v189 = int64(55536)
								v190 = int64(0)
								v195 = int64(32)
								v198 = int64(base.Ui64(v187) >> (uint(v195) % 64))
								v201 = int64(4294967295)
								v204 = v187 & v201
								v205 = v189 * v204
								v209 = int64(base.Ui64(v205)>>(uint(v195)%64)) + v189*v198
								v216 = v204*v190 + v209&v201
								*(*int64)(unsafe.Add(mBase, uint32(v186)+8)) = v187*v190 + v188*v189 + v190*v198 + int64(base.Ui64(v209)>>(uint(v195)%64)) + int64(base.Ui64(v216)>>(uint(v195)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v186))) = v205&v201 | v216<<(uint(v195)%64)
								v228 = v157 - int32(2)
								v229 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
								v230 = v229 + v165
								*(*uint16)(unsafe.Add(mBase, uint32(v228))) = uint16(v230)
								v234 = int64(0)
								v239 = v155 + int32(1)
								if v164 == v234 {
									v240 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v165))
								} else {
									v240 = base.B2i32(v164 != v234)
								}
								if v240 != 0 {
									v155 = v239
									v157 = v228
									v164 = v188
									v165 = v187
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v228
							v242 = v239
							v246 = v155
						} else {
							v242 = int32(0)
							v246 = v2
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v246
					*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v242
					F_numericvar_serialize(m, v16-int32(-64), v16+int32(40))
					mBase = m.M
					v262 = m.ExcPending
					if v262 != 0 {
						return int32(0)
					} else {
						v264 = v16 - int32(-64)
						v266 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
						v267 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v266))) = v267 << (uint(int32(2)) % 32)
						F_pfree(m, v109)
						mBase = m.M
						v272 = m.ExcPending
						if v272 != 0 {
							return int32(0)
						} else {
							m.G0 = v16 + int32(80)
							return v266
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v138 = m.ExcPending
		if v138 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(66863), int32(0))
			mBase = m.M
			v142 = m.ExcPending
			if v142 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(523892), int32(6007), int32(358873))
				mBase = m.M
				v147 = m.ExcPending
				if v147 != 0 {
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
