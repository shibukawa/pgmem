package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ghstore_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v262 int32
	_ = v262
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 == v2 {
		v27 = v2
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
		if v14 == int32(0) {
			v27 = v2
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			if v17 != int32(7) {
				v27 = v2
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				if v20 != int32(17) {
					v27 = v2
				} else {
					v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)))
					v27 = v23 ^ int32(1)
				}
			}
		}
	}
	if v27&int32(1) != 0 {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v31 = F_get_fn_opclass_options(m, v30)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
			v36 = v35
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v39 = int32(0)
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
			v46 = int32(4)
			v47 = v45 & v46
			v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+4)))
			if v48&v46 != 0 {
				if v47 != 0 {
					v262 = int32(0)
				} else {
					v53 = v36 << (uint(int32(3)) % 32)
					if v36 <= int32(0) {
						v262 = v53
					} else {
						v59 = v38 + int32(8)
						v62 = int32(0)
						v64 = v39
						for {
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
							v69 = int32(1)
							v104 = v68&v69 + v62 + int32(base.Ui32(v68)>>(uint(int32(7))%32)) + int32(base.Ui32(v68)>>(uint(v69)%32))&v69 + int32(base.Ui32(v68)>>(uint(int32(2))%32))&v69 + int32(base.Ui32(v68)>>(uint(int32(3))%32))&v69 + int32(base.Ui32(v68)>>(uint(int32(4))%32))&v69 + int32(base.Ui32(v68)>>(uint(int32(5))%32))&v69 + int32(base.Ui32(v68)>>(uint(int32(6))%32))&v69
							v108 = v64 + v69
							if v108 != v36 {
								v59 = v59 + v69
								v62 = v104
								v64 = v108
								continue
							} else {
								break
							}
							break
						}
						v262 = v53 - v104
					}
				}
			} else {
				if v47 != 0 {
					v112 = v36 << (uint(int32(3)) % 32)
					if v36 <= int32(0) {
						v262 = v112
					} else {
						v118 = v37 + int32(8)
						v121 = int32(0)
						v123 = v39
						for {
							v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
							v128 = int32(1)
							v163 = v127&v128 + v121 + int32(base.Ui32(v127)>>(uint(int32(7))%32)) + int32(base.Ui32(v127)>>(uint(v128)%32))&v128 + int32(base.Ui32(v127)>>(uint(int32(2))%32))&v128 + int32(base.Ui32(v127)>>(uint(int32(3))%32))&v128 + int32(base.Ui32(v127)>>(uint(int32(4))%32))&v128 + int32(base.Ui32(v127)>>(uint(int32(5))%32))&v128 + int32(base.Ui32(v127)>>(uint(int32(6))%32))&v128
							v167 = v123 + v128
							if v167 != v36 {
								v118 = v118 + v128
								v121 = v163
								v123 = v167
								continue
							} else {
								break
							}
							break
						}
						v262 = v112 - v163
					}
				} else {
					if v36 <= int32(0) {
						v262 = int32(0)
					} else {
						v173 = int32(8)
						v174 = v38 + v173
						v176 = v37 + v173
						v177 = int32(1)
						v179 = v36 << (uint(int32(3)) % 32)
						if v179 <= v177 {
							v182 = v177
						} else {
							v182 = v179
						}
						v183 = int32(1)
						if v182 == v183 {
							v187 = int32(0)
							v228 = v187
							v229 = v187
						} else {
							v191 = int32(0)
							v194 = v191
							v195 = v191
							v196 = v191
							for {
								v204 = int32(base.Ui32(v195) >> (uint(int32(3)) % 32))
								v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+v204))))
								v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204+v176))))
								v210 = base.I32_extend8_s(v206 ^ v208)
								v212 = v195 & int32(6)
								v213 = int32(1)
								v222 = int32(base.Ui32(v210)>>(uint(v212|v213)%32))&v213 + (int32(base.Ui32(v210)>>(uint(v212)%32))&v213 + v194)
								v223 = int32(2)
								v224 = v195 + v223
								v226 = v196 + v223
								if v226 != v182&int32(2147483640) {
									v194 = v222
									v195 = v224
									v196 = v226
									continue
								} else {
									break
								}
								break
							}
							v228 = v222
							v229 = v224
						}
						if v182&v183 != 0 {
							v238 = int32(base.Ui32(v229) >> (uint(int32(3)) % 32))
							v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+v238))))
							v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238+v176))))
							v252 = int32(base.Ui32(base.I32_extend8_s(v240^v242))>>(uint(v229&int32(7))%32))&int32(1) + v228
						} else {
							v252 = v228
						}
						v262 = v252
					}
				}
			}
			*(*float32)(unsafe.Add(mBase, uint32(v6))) = base.F32_convert_i32_s(v262)
			return v6
		}
	} else {
		v36 = int32(16)
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		v39 = int32(0)
		v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
		v46 = int32(4)
		v47 = v45 & v46
		v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+4)))
		if v48&v46 != 0 {
			if v47 != 0 {
				v262 = int32(0)
			} else {
				v53 = v36 << (uint(int32(3)) % 32)
				if v36 <= int32(0) {
					v262 = v53
				} else {
					v59 = v38 + int32(8)
					v62 = int32(0)
					v64 = v39
					for {
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
						v69 = int32(1)
						v104 = v68&v69 + v62 + int32(base.Ui32(v68)>>(uint(int32(7))%32)) + int32(base.Ui32(v68)>>(uint(v69)%32))&v69 + int32(base.Ui32(v68)>>(uint(int32(2))%32))&v69 + int32(base.Ui32(v68)>>(uint(int32(3))%32))&v69 + int32(base.Ui32(v68)>>(uint(int32(4))%32))&v69 + int32(base.Ui32(v68)>>(uint(int32(5))%32))&v69 + int32(base.Ui32(v68)>>(uint(int32(6))%32))&v69
						v108 = v64 + v69
						if v108 != v36 {
							v59 = v59 + v69
							v62 = v104
							v64 = v108
							continue
						} else {
							break
						}
						break
					}
					v262 = v53 - v104
				}
			}
		} else {
			if v47 != 0 {
				v112 = v36 << (uint(int32(3)) % 32)
				if v36 <= int32(0) {
					v262 = v112
				} else {
					v118 = v37 + int32(8)
					v121 = int32(0)
					v123 = v39
					for {
						v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
						v128 = int32(1)
						v163 = v127&v128 + v121 + int32(base.Ui32(v127)>>(uint(int32(7))%32)) + int32(base.Ui32(v127)>>(uint(v128)%32))&v128 + int32(base.Ui32(v127)>>(uint(int32(2))%32))&v128 + int32(base.Ui32(v127)>>(uint(int32(3))%32))&v128 + int32(base.Ui32(v127)>>(uint(int32(4))%32))&v128 + int32(base.Ui32(v127)>>(uint(int32(5))%32))&v128 + int32(base.Ui32(v127)>>(uint(int32(6))%32))&v128
						v167 = v123 + v128
						if v167 != v36 {
							v118 = v118 + v128
							v121 = v163
							v123 = v167
							continue
						} else {
							break
						}
						break
					}
					v262 = v112 - v163
				}
			} else {
				if v36 <= int32(0) {
					v262 = int32(0)
				} else {
					v173 = int32(8)
					v174 = v38 + v173
					v176 = v37 + v173
					v177 = int32(1)
					v179 = v36 << (uint(int32(3)) % 32)
					if v179 <= v177 {
						v182 = v177
					} else {
						v182 = v179
					}
					v183 = int32(1)
					if v182 == v183 {
						v187 = int32(0)
						v228 = v187
						v229 = v187
					} else {
						v191 = int32(0)
						v194 = v191
						v195 = v191
						v196 = v191
						for {
							v204 = int32(base.Ui32(v195) >> (uint(int32(3)) % 32))
							v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+v204))))
							v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204+v176))))
							v210 = base.I32_extend8_s(v206 ^ v208)
							v212 = v195 & int32(6)
							v213 = int32(1)
							v222 = int32(base.Ui32(v210)>>(uint(v212|v213)%32))&v213 + (int32(base.Ui32(v210)>>(uint(v212)%32))&v213 + v194)
							v223 = int32(2)
							v224 = v195 + v223
							v226 = v196 + v223
							if v226 != v182&int32(2147483640) {
								v194 = v222
								v195 = v224
								v196 = v226
								continue
							} else {
								break
							}
							break
						}
						v228 = v222
						v229 = v224
					}
					if v182&v183 != 0 {
						v238 = int32(base.Ui32(v229) >> (uint(int32(3)) % 32))
						v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+v238))))
						v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238+v176))))
						v252 = int32(base.Ui32(base.I32_extend8_s(v240^v242))>>(uint(v229&int32(7))%32))&int32(1) + v228
					} else {
						v252 = v228
					}
					v262 = v252
				}
			}
		}
		*(*float32)(unsafe.Add(mBase, uint32(v6))) = base.F32_convert_i32_s(v262)
		return v6
	}
}
