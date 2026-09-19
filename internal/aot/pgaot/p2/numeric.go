package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_do_numeric_discard(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int64
	_ = v19
	var v23 int64
	_ = v23
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v89 int32
	_ = v89
	var v95 int64
	_ = v95
	var v101 int64
	_ = v101
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int64
	_ = v127
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v240 int32
	_ = v240
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if base.Ui32(int32(_a_F_do_numeric_discard_0)) <= base.Ui32(v12) {
		if v12 != int32(_a_F_do_numeric_discard_1) {
			if v12 != int32(_a_F_do_numeric_discard_2) {
				v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v27 - int64(1)
			} else {
				v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+96)) = v19 - int64(1)
			}
		} else {
			v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+104)) = v23 - int64(1)
		}
		v240 = int32(1)
		m.G0 = v10 + int32(48)
		return v240
	} else {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v36 = base.I32_extend16_s(v12)
		v38 = base.B2i32(int32(0) <= v36)
		if int32(0) <= v36 {
			v39 = int32(-8)
		} else {
			v39 = int32(-6)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(base.Ui32(int32(base.Ui32(v31)>>(uint(int32(2))%32))+v39) >> (uint(int32(1)) % 32))
		if int32(0) <= v36 {
			v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
			v54 = v44
		} else {
			v54 = v12<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v12&int32(63)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v54
		v56 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v56
		v65 = base.B2i32(v36 < v56)
		if v36 < v56 {
			v66 = int32(base.Ui32(v12)>>(uint(int32(7))%32)) & int32(63)
		} else {
			v66 = v12 & int32(_a_F_do_numeric_discard_3)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v66
		v73 = v12 & int32(_a_F_do_numeric_discard_0)
		if v73 == int32(_a_F_do_numeric_discard_4) {
			v76 = v12 << (uint(int32(1)) % 32) & int32(_a_F_do_numeric_discard_5)
		} else {
			v76 = v73
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v76
		if v36 < v56 {
			v80 = int32(6)
		} else {
			v80 = int32(8)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = l1 + v80
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		if v83 == v66 {
			v86 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
			if v86 < int64(2) {
				v89 = v66
			} else {
				v89 = int32(0)
			}
			if v89 == int32(0) {
				v101 = v86 - int64(1)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v101
				v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				if v104 == int32(1) {
					v107 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v107
					*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v107
					*(*int64)(unsafe.Add(mBase, uint32(v10))) = v107
					v114 = v10 + int32(24)
					F_mul_var(m, v114, v114, v10, v66<<(uint(int32(1))%32))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						v122 = int32(_a_F_do_numeric_discard_6)
						v123 = *(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0]))
						v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v125
						v127 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v127 - int64(1)
						if int64(2) <= v127 {
							if v76 != 0 {
								v135 = int32(0)
							} else {
								v135 = int32(_a_F_do_numeric_discard_5)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v135
							F_accum_sum_add(m, l0+int32(16), v10+int32(24))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
								if v143 != int32(1) {
									*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
									v240 = int32(1)
									m.G0 = v10 + int32(48)
									return v240
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(_a_F_do_numeric_discard_5)
									F_accum_sum_add(m, l0+int32(44), v10)
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
										v240 = int32(1)
										m.G0 = v10 + int32(48)
										return v240
									}
								}
							}
						} else {
							v152 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v152
							v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							if v152 < v154 {
								v159 = int32(0)
								for {
									v166 = v159 << (uint(int32(2)) % 32)
									v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v169 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v166+v167))) = v169
									v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									*(*int32)(unsafe.Add(mBase, uint32(v171+v166))) = v169
									v176 = v159 + int32(1)
									v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v176 < v177 {
										v159 = v176
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
							if v186 != int32(1) {
							} else {
								v189 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v189
								v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								if v191 <= v189 {
								} else {
									v196 = int32(0)
									for {
										v203 = v196 << (uint(int32(2)) % 32)
										v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
										v206 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v203+v204))) = v206
										v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
										*(*int32)(unsafe.Add(mBase, uint32(v208+v203))) = v206
										v213 = v196 + int32(1)
										v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										if v213 < v214 {
											v196 = v213
											continue
										} else {
											break
										}
										break
									}
								}
							}
							*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
							v240 = int32(1)
							m.G0 = v10 + int32(48)
							return v240
						}
					}
				} else {
					v122 = int32(_a_F_do_numeric_discard_6)
					v123 = *(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0]))
					v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v125
					v127 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v127 - int64(1)
					if int64(2) <= v127 {
						if v76 != 0 {
							v135 = int32(0)
						} else {
							v135 = int32(_a_F_do_numeric_discard_5)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v135
						F_accum_sum_add(m, l0+int32(16), v10+int32(24))
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return int32(0)
						} else {
							v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
							if v143 != int32(1) {
								*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
								v240 = int32(1)
								m.G0 = v10 + int32(48)
								return v240
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(_a_F_do_numeric_discard_5)
								F_accum_sum_add(m, l0+int32(44), v10)
								mBase = m.M
								v151 = m.ExcPending
								if v151 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
									v240 = int32(1)
									m.G0 = v10 + int32(48)
									return v240
								}
							}
						}
					} else {
						v152 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v152
						v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v152 < v154 {
							v159 = int32(0)
							for {
								v166 = v159 << (uint(int32(2)) % 32)
								v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								v169 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v166+v167))) = v169
								v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v171+v166))) = v169
								v176 = v159 + int32(1)
								v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v176 < v177 {
									v159 = v176
									continue
								} else {
									break
								}
								break
							}
						} else {
						}
						v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
						if v186 != int32(1) {
						} else {
							v189 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v189
							v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v191 <= v189 {
							} else {
								v196 = int32(0)
								for {
									v203 = v196 << (uint(int32(2)) % 32)
									v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
									v206 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v203+v204))) = v206
									v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
									*(*int32)(unsafe.Add(mBase, uint32(v208+v203))) = v206
									v213 = v196 + int32(1)
									v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									if v213 < v214 {
										v196 = v213
										continue
									} else {
										break
									}
									break
								}
							}
						}
						*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
						v240 = int32(1)
						m.G0 = v10 + int32(48)
						return v240
					}
				}
			} else {
				v95 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				if v95 != int64(1) {
					v240 = int32(0)
					m.G0 = v10 + int32(48)
					return v240
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = int32(0)
					v101 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v101
					v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v104 == int32(1) {
						v107 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v107
						*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v107
						*(*int64)(unsafe.Add(mBase, uint32(v10))) = v107
						v114 = v10 + int32(24)
						F_mul_var(m, v114, v114, v10, v66<<(uint(int32(1))%32))
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							v122 = int32(_a_F_do_numeric_discard_6)
							v123 = *(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0]))
							v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v125
							v127 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v127 - int64(1)
							if int64(2) <= v127 {
								if v76 != 0 {
									v135 = int32(0)
								} else {
									v135 = int32(_a_F_do_numeric_discard_5)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v135
								F_accum_sum_add(m, l0+int32(16), v10+int32(24))
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return int32(0)
								} else {
									v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
									if v143 != int32(1) {
										*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
										v240 = int32(1)
										m.G0 = v10 + int32(48)
										return v240
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(_a_F_do_numeric_discard_5)
										F_accum_sum_add(m, l0+int32(44), v10)
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
											v240 = int32(1)
											m.G0 = v10 + int32(48)
											return v240
										}
									}
								}
							} else {
								v152 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v152
								v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v152 < v154 {
									v159 = int32(0)
									for {
										v166 = v159 << (uint(int32(2)) % 32)
										v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										v169 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v166+v167))) = v169
										v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										*(*int32)(unsafe.Add(mBase, uint32(v171+v166))) = v169
										v176 = v159 + int32(1)
										v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v176 < v177 {
											v159 = v176
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
								if v186 != int32(1) {
								} else {
									v189 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v189
									v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									if v191 <= v189 {
									} else {
										v196 = int32(0)
										for {
											v203 = v196 << (uint(int32(2)) % 32)
											v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
											v206 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v203+v204))) = v206
											v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											*(*int32)(unsafe.Add(mBase, uint32(v208+v203))) = v206
											v213 = v196 + int32(1)
											v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											if v213 < v214 {
												v196 = v213
												continue
											} else {
												break
											}
											break
										}
									}
								}
								*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
								v240 = int32(1)
								m.G0 = v10 + int32(48)
								return v240
							}
						}
					} else {
						v122 = int32(_a_F_do_numeric_discard_6)
						v123 = *(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0]))
						v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v125
						v127 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v127 - int64(1)
						if int64(2) <= v127 {
							if v76 != 0 {
								v135 = int32(0)
							} else {
								v135 = int32(_a_F_do_numeric_discard_5)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v135
							F_accum_sum_add(m, l0+int32(16), v10+int32(24))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
								if v143 != int32(1) {
									*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
									v240 = int32(1)
									m.G0 = v10 + int32(48)
									return v240
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(_a_F_do_numeric_discard_5)
									F_accum_sum_add(m, l0+int32(44), v10)
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
										v240 = int32(1)
										m.G0 = v10 + int32(48)
										return v240
									}
								}
							}
						} else {
							v152 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v152
							v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							if v152 < v154 {
								v159 = int32(0)
								for {
									v166 = v159 << (uint(int32(2)) % 32)
									v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v169 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v166+v167))) = v169
									v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									*(*int32)(unsafe.Add(mBase, uint32(v171+v166))) = v169
									v176 = v159 + int32(1)
									v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v176 < v177 {
										v159 = v176
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
							if v186 != int32(1) {
							} else {
								v189 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v189
								v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								if v191 <= v189 {
								} else {
									v196 = int32(0)
									for {
										v203 = v196 << (uint(int32(2)) % 32)
										v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
										v206 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v203+v204))) = v206
										v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
										*(*int32)(unsafe.Add(mBase, uint32(v208+v203))) = v206
										v213 = v196 + int32(1)
										v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										if v213 < v214 {
											v196 = v213
											continue
										} else {
											break
										}
										break
									}
								}
							}
							*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
							v240 = int32(1)
							m.G0 = v10 + int32(48)
							return v240
						}
					}
				}
			}
		} else {
			v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v104 == int32(1) {
				v107 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v107
				*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v107
				*(*int64)(unsafe.Add(mBase, uint32(v10))) = v107
				v114 = v10 + int32(24)
				F_mul_var(m, v114, v114, v10, v66<<(uint(int32(1))%32))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int32(0)
				} else {
					v122 = int32(_a_F_do_numeric_discard_6)
					v123 = *(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0]))
					v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v125
					v127 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v127 - int64(1)
					if int64(2) <= v127 {
						if v76 != 0 {
							v135 = int32(0)
						} else {
							v135 = int32(_a_F_do_numeric_discard_5)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v135
						F_accum_sum_add(m, l0+int32(16), v10+int32(24))
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return int32(0)
						} else {
							v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
							if v143 != int32(1) {
								*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
								v240 = int32(1)
								m.G0 = v10 + int32(48)
								return v240
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(_a_F_do_numeric_discard_5)
								F_accum_sum_add(m, l0+int32(44), v10)
								mBase = m.M
								v151 = m.ExcPending
								if v151 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
									v240 = int32(1)
									m.G0 = v10 + int32(48)
									return v240
								}
							}
						}
					} else {
						v152 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v152
						v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v152 < v154 {
							v159 = int32(0)
							for {
								v166 = v159 << (uint(int32(2)) % 32)
								v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								v169 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v166+v167))) = v169
								v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v171+v166))) = v169
								v176 = v159 + int32(1)
								v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v176 < v177 {
									v159 = v176
									continue
								} else {
									break
								}
								break
							}
						} else {
						}
						v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
						if v186 != int32(1) {
						} else {
							v189 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v189
							v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v191 <= v189 {
							} else {
								v196 = int32(0)
								for {
									v203 = v196 << (uint(int32(2)) % 32)
									v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
									v206 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v203+v204))) = v206
									v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
									*(*int32)(unsafe.Add(mBase, uint32(v208+v203))) = v206
									v213 = v196 + int32(1)
									v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									if v213 < v214 {
										v196 = v213
										continue
									} else {
										break
									}
									break
								}
							}
						}
						*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
						v240 = int32(1)
						m.G0 = v10 + int32(48)
						return v240
					}
				}
			} else {
				v122 = int32(_a_F_do_numeric_discard_6)
				v123 = *(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0]))
				v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v125
				v127 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v127 - int64(1)
				if int64(2) <= v127 {
					if v76 != 0 {
						v135 = int32(0)
					} else {
						v135 = int32(_a_F_do_numeric_discard_5)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v135
					F_accum_sum_add(m, l0+int32(16), v10+int32(24))
					mBase = m.M
					v142 = m.ExcPending
					if v142 != 0 {
						return int32(0)
					} else {
						v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
						if v143 != int32(1) {
							*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
							v240 = int32(1)
							m.G0 = v10 + int32(48)
							return v240
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(_a_F_do_numeric_discard_5)
							F_accum_sum_add(m, l0+int32(44), v10)
							mBase = m.M
							v151 = m.ExcPending
							if v151 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
								v240 = int32(1)
								m.G0 = v10 + int32(48)
								return v240
							}
						}
					}
				} else {
					v152 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v152
					v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v152 < v154 {
						v159 = int32(0)
						for {
							v166 = v159 << (uint(int32(2)) % 32)
							v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v169 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v166+v167))) = v169
							v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v171+v166))) = v169
							v176 = v159 + int32(1)
							v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							if v176 < v177 {
								v159 = v176
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v186 != int32(1) {
					} else {
						v189 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v189
						v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						if v191 <= v189 {
						} else {
							v196 = int32(0)
							for {
								v203 = v196 << (uint(int32(2)) % 32)
								v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
								v206 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v203+v204))) = v206
								v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
								*(*int32)(unsafe.Add(mBase, uint32(v208+v203))) = v206
								v213 = v196 + int32(1)
								v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								if v213 < v214 {
									v196 = v213
									continue
								} else {
									break
								}
								break
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_discard[0])) = v123
					v240 = int32(1)
					m.G0 = v10 + int32(48)
					return v240
				}
			}
		}
	}
}
func F_numeric_abs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v12 = F_palloc(m, int32(base.Ui32(v9)>>(uint(int32(2))%32)))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v16 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
			if v16 != 0 {
				base.MemoryCopy(m, v12, v5, v16)
			} else {
			}
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+4)))
			if base.Ui32(int32(_a_F_numeric_abs_0)) < base.Ui32(v18) {
				v24 = int32(-8193)
			} else {
				v24 = int32(_a_F_numeric_abs_1)
			}
			if base.I32_extend16_s(v18) < int32(-16384) {
				v28 = int32(-24577)
			} else {
				v28 = v24
			}
			v29 = v18 & v28
			*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v29)
			return v12
		}
	}
}
func F_numeric_avg(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v23 int64
	_ = v23
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v13 != 0 {
		v31 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
		v96 = int32(0)
		m.G0 = v11 + int32(32)
		return v96
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v14 == int32(0) {
			v31 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
			v96 = int32(0)
			m.G0 = v11 + int32(32)
			return v96
		} else {
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v14)+96))
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v14)+88))
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v14)+104))
			if v17+(v18+v19) != int64(0)-v23 {
				if int64(0) < v18 {
					v38 = F_make_result_opt_error(m, int32(_a_F_numeric_avg_0), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v96 = v38
						m.G0 = v11 + int32(32)
						return v96
					}
				} else {
					v42 = int64(0)
					if base.B2i32(v17 <= v42)|base.B2i32(v23 <= v42) == int32(0) {
						v51 = F_make_result_opt_error(m, int32(_a_F_numeric_avg_0), int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v96 = v51
							m.G0 = v11 + int32(32)
							return v96
						}
					} else {
						if int64(0) < v17 {
							v57 = F_make_result_opt_error(m, int32(_a_F_numeric_avg_1), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								v96 = v57
								m.G0 = v11 + int32(32)
								return v96
							}
						} else {
							if int64(0) < v23 {
								v63 = F_make_result_opt_error(m, int32(_a_F_numeric_avg_2), int32(0))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									v96 = v63
									m.G0 = v11 + int32(32)
									return v96
								}
							} else {
								v65 = F_int64_to_numeric(m, v19)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v67 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v67
									*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v67
									*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v67
									v76 = v11 + int32(8)
									F_accum_sum_final(m, v14+int32(16), v76)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										v80 = F_make_result_opt_error(m, v76, int32(0))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
											if v82 != 0 {
												F_pfree(m, v82)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													v87 = F_DirectFunctionCall2Coll(m, int32(1260), int32(0), v80, v65)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return int32(0)
													} else {
														v96 = v87
														m.G0 = v11 + int32(32)
														return v96
													}
												}
											} else {
												v87 = F_DirectFunctionCall2Coll(m, int32(1260), int32(0), v80, v65)
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return int32(0)
												} else {
													v96 = v87
													m.G0 = v11 + int32(32)
													return v96
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v31 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
				v96 = int32(0)
				m.G0 = v11 + int32(32)
				return v96
			}
		}
	}
}
func F_numeric_avg_accum(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13951(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_numeric_avg_combine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int64
	_ = v135
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = v8 + int32(4)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 == v2 {
		v30 = int32(0)
		if v11 == v30 {
			v38 = v30
		} else {
			v33 = v30
			v34 = v2
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v33
			v38 = v34
		}
		v41 = v38
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		switch v16 - int32(429) {
		case 0:
			if v11 == int32(0) {
				v41 = int32(1)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+168))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
				v33 = v23
				v34 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v33
				v38 = v34
				v41 = v38
			}
		case 1:
			if v11 == int32(0) {
				v41 = int32(2)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+368))
				v33 = v28
				v34 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v33
				v38 = v34
				v41 = v38
			}
		default:
			v30 = int32(0)
			if v11 == v30 {
				v38 = v30
			} else {
				v33 = v30
				v34 = v2
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v33
				v38 = v34
			}
			v41 = v38
		}
	}
	if v41 != 0 {
		v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v42 == int32(0) {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v46 = v45
		} else {
			v46 = v2
		}
		v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v47 != 0 {
			v175 = v46
			m.G0 = v8 + int32(32)
			return v175
		} else {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v48 == int32(0) {
				v175 = v46
				m.G0 = v8 + int32(32)
				return v175
			} else {
				if v46 == int32(0) {
					v53 = int32(_a_F_numeric_avg_combine_0)
					v54 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_avg_combine[0]))
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*int32)(unsafe.Add(mBase, _c_F_numeric_avg_combine[0])) = v56
					v59 = F_palloc0(m, int32(112))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v63)
						v66 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_avg_combine[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v66
						v68 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v59)+8)) = v68
						v70 = *(*int64)(unsafe.Add(mBase, uint32(v48)+88))
						*(*int64)(unsafe.Add(mBase, uint32(v59)+88)) = v70
						v72 = *(*int64)(unsafe.Add(mBase, uint32(v48)+96))
						*(*int64)(unsafe.Add(mBase, uint32(v59)+96)) = v72
						v74 = *(*int64)(unsafe.Add(mBase, uint32(v48)+104))
						*(*int64)(unsafe.Add(mBase, uint32(v59)+104)) = v74
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v48)+72))
						*(*int32)(unsafe.Add(mBase, uint32(v59)+72)) = v76
						v78 = *(*int64)(unsafe.Add(mBase, uint32(v48)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v59)+80)) = v78
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
						v83 = F_palloc(m, v80<<(uint(int32(2))%32))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v59)+36)) = v83
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							v89 = F_palloc(m, v86<<(uint(int32(2))%32))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v59)+40)) = v89
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								v94 = v92 << (uint(int32(2)) % 32)
								if v94 != 0 {
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v59)+36))
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+36))
									base.MemoryCopy(m, v95, v96, v94)
								} else {
								}
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								v100 = v98 << (uint(int32(2)) % 32)
								if v100 != 0 {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v59)+40))
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v48)+40))
									base.MemoryCopy(m, v101, v102, v100)
								} else {
								}
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v48)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+28)) = v104
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v106
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v108
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = v110
								v169 = v59
								v170 = v54
								*(*int32)(unsafe.Add(mBase, _c_F_numeric_avg_combine[0])) = v170
								v175 = v169
								m.G0 = v8 + int32(32)
								return v175
							}
						}
					}
				} else {
					v112 = *(*int64)(unsafe.Add(mBase, uint32(v46)+8))
					v113 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+8)) = v112 + v113
					v116 = *(*int64)(unsafe.Add(mBase, uint32(v46)+88))
					v117 = *(*int64)(unsafe.Add(mBase, uint32(v48)+88))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+88)) = v116 + v117
					v120 = *(*int64)(unsafe.Add(mBase, uint32(v46)+96))
					v121 = *(*int64)(unsafe.Add(mBase, uint32(v48)+96))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+96)) = v120 + v121
					v124 = *(*int64)(unsafe.Add(mBase, uint32(v46)+104))
					v125 = *(*int64)(unsafe.Add(mBase, uint32(v48)+104))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+104)) = v124 + v125
					v128 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
					if v128 <= int64(0) {
						v175 = v46
						m.G0 = v8 + int32(32)
						return v175
					} else {
						v131 = *(*int32)(unsafe.Add(mBase, uint32(v48)+72))
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v46)+72))
						if v132 < v131 {
							*(*int32)(unsafe.Add(mBase, uint32(v46)+72)) = v131
							v135 = *(*int64)(unsafe.Add(mBase, uint32(v48)+80))
							*(*int64)(unsafe.Add(mBase, uint32(v46)+80)) = v135
						} else {
							if v131 != v132 {
							} else {
								v138 = *(*int64)(unsafe.Add(mBase, uint32(v46)+80))
								v139 = *(*int64)(unsafe.Add(mBase, uint32(v48)+80))
								*(*int64)(unsafe.Add(mBase, uint32(v46)+80)) = v138 + v139
							}
						}
						v142 = int32(_a_F_numeric_avg_combine_0)
						v143 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_avg_combine[0]))
						v145 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						*(*int32)(unsafe.Add(mBase, _c_F_numeric_avg_combine[0])) = v145
						v147 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v147
						*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v147
						*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v147
						v156 = v8 + int32(8)
						F_accum_sum_final(m, v48+int32(16), v156)
						mBase = m.M
						v158 = m.ExcPending
						if v158 != 0 {
							return int32(0)
						} else {
							F_accum_sum_add(m, v46+int32(16), v156)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								v163 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
								if v163 == int32(0) {
									v169 = v46
									v170 = v143
									*(*int32)(unsafe.Add(mBase, _c_F_numeric_avg_combine[0])) = v170
									v175 = v169
									m.G0 = v8 + int32(32)
									return v175
								} else {
									F_pfree(m, v163)
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return int32(0)
									} else {
										v169 = v46
										v170 = v143
										*(*int32)(unsafe.Add(mBase, _c_F_numeric_avg_combine[0])) = v170
										v175 = v169
										m.G0 = v8 + int32(32)
										return v175
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v185 = m.ExcPending
		if v185 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_avg_combine_1), int32(0))
			mBase = m.M
			v189 = m.ExcPending
			if v189 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_avg_combine_2), int32(_a_F_numeric_avg_combine_3), int32(_a_F_numeric_avg_combine_4))
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
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
func F_numeric_div_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int64
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v18)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v21 = base.I32_extend16_s(v20)
	if base.Ui32(v20) <= base.Ui32(int32(_a_F_numeric_div_opt_error_0)) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	m.G0 = v16 + int32(80)
	return v421
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v191
	v193 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v193
	v202 = base.B2i32(v21 < v193)
	if v21 < v193 {
		goto L76
	} else {
		goto L77
	}
L6:
	;
	v190 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	v191 = v190
	goto L5
L7:
	;
	if v20 != int32(_a_F_numeric_div_opt_error_1) {
		goto L23
	} else {
		goto L24
	}
L8:
	;
	v62 = F_make_result_opt_error(m, int32(_a_F_numeric_div_opt_error_2), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L20
	} else {
		goto L21
	}
L9:
	;
	if v52&int32(_a_F_numeric_div_opt_error_3) != int32(_a_F_numeric_div_opt_error_4) {
		goto L7
	} else {
		goto L19
	}
L10:
	;
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v25 = base.I32_extend16_s(v24)
	if base.Ui32(int32(_a_F_numeric_div_opt_error_0)) < base.Ui32(v24) {
		v52 = v25
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v21 == int32(-16384) {
		goto L8
	} else {
		goto L18
	}
L13:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = base.B2i32(int32(0) <= v21)
	if int32(0) <= v21 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v35 = int32(-8)
	goto L16
L15:
	;
	v35 = int32(-6)
	goto L16
L16:
	;
	v38 = int32(base.Ui32(int32(base.Ui32(v28)>>(uint(int32(2))%32))+v35) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v38
	if int32(0) <= v21 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v191 = v20<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v20&int32(63)
	goto L5
L18:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v52 = v51
	goto L9
L19:
	;
	goto L8
L20:
	;
	return int32(0)
L21:
	;
	v421 = v62
	goto L4
L22:
	;
	v188 = F_make_result_opt_error(m, int32(_a_F_numeric_div_opt_error_5), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L20
	} else {
		goto L75
	}
L23:
	;
	if v20 != int32(_a_F_numeric_div_opt_error_6) {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if base.Ui32(int32(_a_F_numeric_div_opt_error_4)) <= base.Ui32(v52&int32(_a_F_numeric_div_opt_error_3)) {
		goto L51
	} else {
		goto L52
	}
L26:
	;
	if base.Ui32(int32(_a_F_numeric_div_opt_error_4)) <= base.Ui32(v52&int32(_a_F_numeric_div_opt_error_3)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v76 = F_make_result_opt_error(m, int32(_a_F_numeric_div_opt_error_2), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L20
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) <= base.I32_extend16_s(v52) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v421 = v76
	goto L4
L31:
	;
	v126 = F_make_result_opt_error(m, int32(_a_F_numeric_div_opt_error_7), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L20
	} else {
		goto L50
	}
L32:
	;
	v86 = int32(-8)
	goto L34
L33:
	;
	v86 = int32(-6)
	goto L34
L34:
	;
	if base.Ui32(int32(2)) <= base.Ui32(int32(base.Ui32(v78)>>(uint(int32(2))%32))+v86) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v95 = v52 & int32(_a_F_numeric_div_opt_error_4)
	if v95 == int32(_a_F_numeric_div_opt_error_8) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	if l2 != 0 {
		goto L43
	} else {
		goto L44
	}
L38:
	;
	v98 = v52 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_opt_error_9)
	goto L40
L39:
	;
	v98 = v95
	goto L40
L40:
	;
	if v98 == int32(_a_F_numeric_div_opt_error_9) {
		goto L31
	} else {
		goto L41
	}
L41:
	;
	v103 = F_make_result_opt_error(m, int32(_a_F_numeric_div_opt_error_10), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L20
	} else {
		goto L42
	}
L42:
	;
	v421 = v103
	goto L4
L43:
	;
	v105 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
	v421 = int32(0)
	goto L4
L44:
	;
	goto L45
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	F_errcode(m, int32(33816706))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L20
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(_a_F_numeric_div_opt_error_11), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L20
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_numeric_div_opt_error_12), int32(3295), int32(_a_F_numeric_div_opt_error_13))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L20
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v421 = v126
	goto L4
L51:
	;
	v134 = F_make_result_opt_error(m, int32(_a_F_numeric_div_opt_error_2), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L20
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) <= base.I32_extend16_s(v52) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v421 = v134
	goto L4
L55:
	;
	v184 = F_make_result_opt_error(m, int32(_a_F_numeric_div_opt_error_10), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L20
	} else {
		goto L74
	}
L56:
	;
	v144 = int32(-8)
	goto L58
L57:
	;
	v144 = int32(-6)
	goto L58
L58:
	;
	if base.Ui32(int32(2)) <= base.Ui32(int32(base.Ui32(v136)>>(uint(int32(2))%32))+v144) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v153 = v52 & int32(_a_F_numeric_div_opt_error_4)
	if v153 == int32(_a_F_numeric_div_opt_error_8) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	if l2 != 0 {
		goto L67
	} else {
		goto L68
	}
L62:
	;
	v156 = v52 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_opt_error_9)
	goto L64
L63:
	;
	v156 = v153
	goto L64
L64:
	;
	if v156 == int32(_a_F_numeric_div_opt_error_9) {
		goto L55
	} else {
		goto L65
	}
L65:
	;
	v161 = F_make_result_opt_error(m, int32(_a_F_numeric_div_opt_error_7), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L20
	} else {
		goto L66
	}
L66:
	;
	v421 = v161
	goto L4
L67:
	;
	v163 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v163)
	v421 = int32(0)
	goto L4
L68:
	;
	goto L69
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L20
	} else {
		goto L70
	}
L70:
	;
	F_errcode(m, int32(33816706))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L20
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(_a_F_numeric_div_opt_error_11), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_numeric_div_opt_error_12), int32(3318), int32(_a_F_numeric_div_opt_error_13))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L20
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v421 = v184
	goto L4
L75:
	;
	v421 = v188
	goto L4
L76:
	;
	v203 = int32(base.Ui32(v20)>>(uint(int32(7))%32)) & int32(63)
	goto L78
L77:
	;
	v203 = v20 & int32(_a_F_numeric_div_opt_error_14)
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v203
	v210 = v20 & int32(_a_F_numeric_div_opt_error_4)
	if v210 == int32(_a_F_numeric_div_opt_error_8) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v213 = v20 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_opt_error_9)
	goto L81
L80:
	;
	v213 = v210
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v213
	if v21 < v193 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v217 = int32(6)
	goto L84
L83:
	;
	v217 = int32(8)
	goto L84
L84:
	;
	v218 = l0 + v217
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v218
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v226 = base.B2i32(int32(0) <= v25)
	if int32(0) <= v25 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v227 = int32(-8)
	goto L87
L86:
	;
	v227 = int32(-6)
	goto L87
L87:
	;
	v230 = int32(base.Ui32(int32(base.Ui32(v220)>>(uint(int32(2))%32))+v227) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v230
	if int32(0) <= v25 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v232 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	v242 = v232
	goto L90
L89:
	;
	v242 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v242
	v244 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v244
	v247 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v247
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v247
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v247
	v256 = base.B2i32(v25 < v244)
	if v25 < v244 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v257 = int32(6)
	goto L93
L92:
	;
	v257 = int32(8)
	goto L93
L93:
	;
	v258 = l1 + v257
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v258
	if v25 < v244 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v266 = int32(base.Ui32(v24)>>(uint(int32(7))%32)) & int32(63)
	goto L96
L95:
	;
	v266 = v24 & int32(_a_F_numeric_div_opt_error_14)
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v266
	v273 = v24 & int32(_a_F_numeric_div_opt_error_4)
	if v273 == int32(_a_F_numeric_div_opt_error_8) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v276 = v24 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_opt_error_9)
	goto L99
L98:
	;
	v276 = v273
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v276
	if v38 != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	if v230 != 0 {
		goto L109
	} else {
		goto L110
	}
L101:
	;
	v279 = int32(0)
	goto L104
L102:
	;
	goto L103
L103:
	;
	v313 = int32(0)
	v321 = v313
	v328 = v313
	goto L100
L104:
	;
	v296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218+v279<<(uint(int32(1))%32)))))
	if v296 != 0 {
		v321 = v296
		v328 = v279 - v191
		goto L100
	} else {
		goto L106
	}
L105:
	;
	goto L103
L106:
	;
	v298 = v279 + int32(1)
	if v298 != v38 {
		v279 = v298
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	if l2 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L109:
	;
	v332 = v244
	goto L112
L110:
	;
	v354 = v244
	goto L111
L111:
	;
	v365 = int32(0)
	v368 = v354
	goto L108
L112:
	;
	v345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258+v332<<(uint(int32(1))%32)))))
	if v345 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v354 = int32(0)
	goto L111
L114:
	;
	v365 = v345
	v368 = v242 - v332
	goto L108
L115:
	;
	goto L116
L116:
	;
	v348 = v332 + int32(1)
	if v348 != v230 {
		v332 = v348
		goto L112
	} else {
		goto L117
	}
L117:
	;
	goto L113
L118:
	;
	v389 = v16 + int32(8)
	v399 = (v368+v328+base.B2i32(base.I32_extend16_s(v321) <= base.I32_extend16_s(v365)))<<(uint(int32(2))%32) + int32(16)
	if v203 < v399 {
		goto L124
	} else {
		goto L125
	}
L119:
	;
	if v230 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v380 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258))))
	if v380 != 0 {
		goto L118
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v381 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v381)
	v421 = int32(0)
	goto L4
L123:
	;
	goto L122
L124:
	;
	v401 = v399
	goto L126
L125:
	;
	v401 = v203
	goto L126
L126:
	;
	if v266 < v401 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v403 = v401
	goto L129
L128:
	;
	v403 = v266
	goto L129
L129:
	;
	if int32(1000) <= v403 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v406 = int32(1000)
	goto L132
L131:
	;
	v406 = v403
	goto L132
L132:
	;
	v407 = int32(1)
	F_div_var(m, v16+int32(56), v16+int32(32), v389, v406, v407, v407)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L20
	} else {
		goto L133
	}
L133:
	;
	v411 = F_make_result_opt_error(m, v389, l2)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L20
	} else {
		goto L134
	}
L134:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if v413 == int32(0) {
		v421 = v411
		goto L4
	} else {
		goto L135
	}
L135:
	;
	F_pfree(m, v413)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L20
	} else {
		goto L136
	}
L136:
	;
	v421 = v411
	goto L4
}
func F_numeric_float4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+4)))
		if base.Ui32(int32(_a_F_numeric_float4_0)) <= base.Ui32(v8) {
			if v8 == int32(_a_F_numeric_float4_1) {
				v16 = int32(-8388608)
			} else {
				v16 = int32(2143289344)
			}
			if v8 == int32(_a_F_numeric_float4_2) {
				v19 = int32(2139095040)
			} else {
				v19 = v16
			}
			return v19
		} else {
			v22 = int32(0)
			v25 = F_DirectFunctionCall1Coll(m, int32(618), v22, v4)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = F_DirectFunctionCall1Coll(m, int32(1454), v22, v25)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v25)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						return v27
					}
				}
			}
		}
	}
}
func F_numeric_gcd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = F_pg_detoast_datum(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
			if base.Ui32(v21) <= base.Ui32(int32(_a_F_numeric_gcd_0)) {
				v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
				if base.Ui32(v24) < base.Ui32(int32(_a_F_numeric_gcd_1)) {
					v32 = base.I32_extend16_s(v24)
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v38 = base.I32_extend16_s(v21)
					v40 = base.B2i32(int32(0) <= v38)
					if int32(0) <= v38 {
						v41 = int32(-8)
					} else {
						v41 = int32(-6)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = int32(base.Ui32(int32(base.Ui32(v33)>>(uint(int32(2))%32))+v41) >> (uint(int32(1)) % 32))
					if int32(0) <= v38 {
						v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+6)))
						v56 = v46
					} else {
						v56 = v21<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v21&int32(63)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v56
					v58 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v58
					v67 = base.B2i32(v38 < v58)
					if v38 < v58 {
						v68 = int32(base.Ui32(v21)>>(uint(int32(7))%32)) & int32(63)
					} else {
						v68 = v21 & int32(_a_F_numeric_gcd_2)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v68
					v75 = v21 & int32(_a_F_numeric_gcd_1)
					if v75 == int32(_a_F_numeric_gcd_3) {
						v78 = v21 << (uint(int32(1)) % 32) & int32(_a_F_numeric_gcd_4)
					} else {
						v78 = v75
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v78
					if v38 < v58 {
						v82 = int32(6)
					} else {
						v82 = int32(8)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v14 + v82
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					v91 = base.B2i32(int32(0) <= v32)
					if int32(0) <= v32 {
						v92 = int32(-8)
					} else {
						v92 = int32(-6)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(base.Ui32(int32(base.Ui32(v85)>>(uint(int32(2))%32))+v92) >> (uint(int32(1)) % 32))
					if int32(0) <= v32 {
						v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+6)))
						v107 = v97
					} else {
						v107 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v107
					v109 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v109
					*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(0)
					v116 = base.B2i32(v32 < v109)
					if v32 < v109 {
						v117 = int32(6)
					} else {
						v117 = int32(8)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v19 + v117
					if v32 < v109 {
						v126 = int32(base.Ui32(v24)>>(uint(int32(7))%32)) & int32(63)
					} else {
						v126 = v24 & int32(_a_F_numeric_gcd_2)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v126
					v133 = v24 & int32(_a_F_numeric_gcd_1)
					if v133 == int32(_a_F_numeric_gcd_3) {
						v136 = v24 << (uint(int32(1)) % 32) & int32(_a_F_numeric_gcd_4)
					} else {
						v136 = v133
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v136
					v143 = v11 + int32(8)
					F_gcd_var(m, v11+int32(56), v11+int32(32), v143)
					mBase = m.M
					v145 = m.ExcPending
					if v145 != 0 {
						return int32(0)
					} else {
						v147 = F_make_result_opt_error(m, v143, int32(0))
						mBase = m.M
						v148 = m.ExcPending
						if v148 != 0 {
							return int32(0)
						} else {
							v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
							if v149 == int32(0) {
								v154 = v147
								m.G0 = v11 + int32(80)
								return v154
							} else {
								F_pfree(m, v149)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return int32(0)
								} else {
									v154 = v147
									m.G0 = v11 + int32(80)
									return v154
								}
							}
						}
					}
				} else {
					v30 = F_make_result_opt_error(m, int32(_a_F_numeric_gcd_5), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v154 = v30
						m.G0 = v11 + int32(80)
						return v154
					}
				}
			} else {
				v30 = F_make_result_opt_error(m, int32(_a_F_numeric_gcd_5), int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v154 = v30
					m.G0 = v11 + int32(80)
					return v154
				}
			}
		}
	}
}
func F_numeric_inc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v18 = int32(base.Ui32(v16) >> (uint(int32(2)) % 32))
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
		if base.Ui32(int32(_a_F_numeric_inc_0)) <= base.Ui32(v19) {
			v22 = F_palloc(m, v18)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v26 = int32(base.Ui32(v24) >> (uint(int32(2)) % 32))
				if v26 == int32(0) {
					v92 = v22
				} else {
					base.MemoryCopy(m, v22, v12, v26)
					v92 = v22
				}
				m.G0 = v9 + int32(32)
				return v92
			}
		} else {
			v32 = base.I32_extend16_s(v19)
			v34 = base.B2i32(int32(0) <= v32)
			if int32(0) <= v32 {
				v35 = int32(-8)
			} else {
				v35 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(base.Ui32(v18+v35) >> (uint(int32(1)) % 32))
			if int32(0) <= v32 {
				v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
				v50 = v40
			} else {
				v50 = v19<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v19&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v50
			v52 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v52
			v57 = base.B2i32(v32 < v52)
			if v32 < v52 {
				v58 = int32(6)
			} else {
				v58 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v12 + v58
			if v32 < v52 {
				v67 = int32(base.Ui32(v19)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v67 = v19 & int32(_a_F_numeric_inc_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v67
			v74 = v19 & int32(_a_F_numeric_inc_0)
			if v74 == int32(_a_F_numeric_inc_2) {
				v77 = v19 << (uint(int32(1)) % 32) & int32(_a_F_numeric_inc_3)
			} else {
				v77 = v74
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v77
			v80 = v9 + int32(8)
			F_add_var(m, v80, int32(_a_F_numeric_inc_4), v80)
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				v85 = F_make_result_opt_error(m, v80, int32(0))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int32(0)
				} else {
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					if v87 == int32(0) {
						v92 = v85
						m.G0 = v9 + int32(32)
						return v92
					} else {
						F_pfree(m, v87)
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v92 = v85
							m.G0 = v9 + int32(32)
							return v92
						}
					}
				}
			}
		}
	}
}
func F_numeric_is_nan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	return base.B2i32(v2 == int32(_a_F_numeric_is_nan_0))
}
func F_numeric_lcm(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int64
	_ = v152
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v20 = F_pg_detoast_datum(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
			if base.Ui32(v22) <= base.Ui32(int32(_a_F_numeric_lcm_0)) {
				v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)))
				if base.Ui32(v25) < base.Ui32(int32(_a_F_numeric_lcm_1)) {
					v33 = base.I32_extend16_s(v25)
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v39 = base.I32_extend16_s(v22)
					v41 = base.B2i32(int32(0) <= v39)
					if int32(0) <= v39 {
						v42 = int32(-8)
					} else {
						v42 = int32(-6)
					}
					v45 = int32(base.Ui32(int32(base.Ui32(v34)>>(uint(int32(2))%32))+v42) >> (uint(int32(1)) % 32))
					*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v45
					if int32(0) <= v39 {
						v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+6)))
						v57 = v47
					} else {
						v57 = v22<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v22&int32(63)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v57
					v59 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v59
					v68 = base.B2i32(v39 < v59)
					if v39 < v59 {
						v69 = int32(base.Ui32(v22)>>(uint(int32(7))%32)) & int32(63)
					} else {
						v69 = v22 & int32(_a_F_numeric_lcm_2)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v69
					v76 = v22 & int32(_a_F_numeric_lcm_1)
					if v76 == int32(_a_F_numeric_lcm_3) {
						v79 = v22 << (uint(int32(1)) % 32) & int32(_a_F_numeric_lcm_4)
					} else {
						v79 = v76
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v79
					if v39 < v59 {
						v83 = int32(6)
					} else {
						v83 = int32(8)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v15 + v83
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
					v92 = base.B2i32(int32(0) <= v33)
					if int32(0) <= v33 {
						v93 = int32(-8)
					} else {
						v93 = int32(-6)
					}
					v96 = int32(base.Ui32(int32(base.Ui32(v86)>>(uint(int32(2))%32))+v93) >> (uint(int32(1)) % 32))
					*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v96
					if int32(0) <= v33 {
						v98 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+6)))
						v108 = v98
					} else {
						v108 = v25<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v25&int32(63)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v108
					*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
					v112 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v112
					v121 = base.B2i32(v33 < v112)
					if v33 < v112 {
						v122 = int32(base.Ui32(v25)>>(uint(int32(7))%32)) & int32(63)
					} else {
						v122 = v25 & int32(_a_F_numeric_lcm_2)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v122
					v129 = v25 & int32(_a_F_numeric_lcm_1)
					if v129 == int32(_a_F_numeric_lcm_3) {
						v132 = v25 << (uint(int32(1)) % 32) & int32(_a_F_numeric_lcm_4)
					} else {
						v132 = v129
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v132
					if v33 < v112 {
						v136 = int32(6)
					} else {
						v136 = int32(8)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v20 + v136
					if v96 != 0 {
						v140 = v45
					} else {
						v140 = int32(0)
					}
					if v140 == int32(0) {
						v144 = F_palloc(m, int32(2))
						mBase = m.M
						v145 = m.ExcPending
						if v145 != 0 {
							return int32(0)
						} else {
							v146 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(v144))) = uint16(v146)
							v149 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_lcm[0]))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v149
							v152 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_lcm[1]))
							*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v152
							*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v144 + int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v144
							v176 = v144
							if base.Ui32(v122) < base.Ui32(v69) {
								v180 = v69
							} else {
								v180 = v122
							}
							*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v180
							v185 = F_make_result_opt_error(m, v12+int32(8), int32(0))
							mBase = m.M
							v186 = m.ExcPending
							if v186 != 0 {
								return int32(0)
							} else {
								if v176 == int32(0) {
									v191 = v185
									m.G0 = v12 + int32(80)
									return v191
								} else {
									F_pfree(m, v176)
									mBase = m.M
									v190 = m.ExcPending
									if v190 != 0 {
										return int32(0)
									} else {
										v191 = v185
										m.G0 = v12 + int32(80)
										return v191
									}
								}
							}
						}
					} else {
						v159 = v12 + int32(56)
						v161 = v12 + int32(32)
						v163 = v12 + int32(8)
						F_gcd_var(m, v159, v161, v163)
						mBase = m.M
						v165 = m.ExcPending
						if v165 != 0 {
							return int32(0)
						} else {
							v166 = int32(0)
							F_div_var(m, v159, v163, v163, v166, v166, int32(1))
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
								return int32(0)
							} else {
								F_mul_var(m, v161, v163, v163, v122)
								mBase = m.M
								v172 = m.ExcPending
								if v172 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
									v175 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
									v176 = v175
									if base.Ui32(v122) < base.Ui32(v69) {
										v180 = v69
									} else {
										v180 = v122
									}
									*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v180
									v185 = F_make_result_opt_error(m, v12+int32(8), int32(0))
									mBase = m.M
									v186 = m.ExcPending
									if v186 != 0 {
										return int32(0)
									} else {
										if v176 == int32(0) {
											v191 = v185
											m.G0 = v12 + int32(80)
											return v191
										} else {
											F_pfree(m, v176)
											mBase = m.M
											v190 = m.ExcPending
											if v190 != 0 {
												return int32(0)
											} else {
												v191 = v185
												m.G0 = v12 + int32(80)
												return v191
											}
										}
									}
								}
							}
						}
					}
				} else {
					v31 = F_make_result_opt_error(m, int32(_a_F_numeric_lcm_5), int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v191 = v31
						m.G0 = v12 + int32(80)
						return v191
					}
				}
			} else {
				v31 = F_make_result_opt_error(m, int32(_a_F_numeric_lcm_5), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v191 = v31
					m.G0 = v12 + int32(80)
					return v191
				}
			}
		}
	}
}
func F_numeric_mod(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v11 = F_numeric_mod_opt_error(m, v3, v8, int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
func F_numeric_mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v11 = F_numeric_mul_opt_error(m, v3, v8, int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
func F_numeric_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
		if base.Ui32(int32(_a_F_numeric_out_0)) <= base.Ui32(v15) {
			if v15 != int32(_a_F_numeric_out_1) {
				if v15 != int32(_a_F_numeric_out_2) {
					v29 = F_pstrdup(m, int32(_a_F_numeric_out_3))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v90 = v29
						m.G0 = v8 + int32(32)
						return v90
					}
				} else {
					v23 = F_pstrdup(m, int32(_a_F_numeric_out_4))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v90 = v23
						m.G0 = v8 + int32(32)
						return v90
					}
				}
			} else {
				v26 = F_pstrdup(m, int32(_a_F_numeric_out_5))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v90 = v26
					m.G0 = v8 + int32(32)
					return v90
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v36 = base.I32_extend16_s(v15)
			v38 = base.B2i32(int32(0) <= v36)
			if int32(0) <= v36 {
				v39 = int32(-8)
			} else {
				v39 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(base.Ui32(int32(base.Ui32(v31)>>(uint(int32(2))%32))+v39) >> (uint(int32(1)) % 32))
			if int32(0) <= v36 {
				v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+6)))
				v54 = v44
			} else {
				v54 = v15<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v15&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v54
			v56 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v56
			v65 = base.B2i32(v36 < v56)
			if v36 < v56 {
				v66 = int32(base.Ui32(v15)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v66 = v15 & int32(_a_F_numeric_out_6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v66
			v73 = v15 & int32(_a_F_numeric_out_0)
			if v73 == int32(_a_F_numeric_out_7) {
				v76 = v15 << (uint(int32(1)) % 32) & int32(_a_F_numeric_out_8)
			} else {
				v76 = v73
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v76
			if v36 < v56 {
				v80 = int32(6)
			} else {
				v80 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v11 + v80
			v85 = F_get_str_from_var(m, v8+int32(8))
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return int32(0)
			} else {
				v90 = v85
				m.G0 = v8 + int32(32)
				return v90
			}
		}
	}
}
func F_numeric_out_sci(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v95 int64
	_ = v95
	var v110 int32
	_ = v110
	var v112 int64
	_ = v112
	var v122 int64
	_ = v122
	var v126 int64
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 float64
	_ = v135
	var v137 float64
	_ = v137
	var v150 float64
	_ = v150
	var v153 float64
	_ = v153
	var v158 float64
	_ = v158
	var v159 float64
	_ = v159
	var v160 float64
	_ = v160
	var v161 float64
	_ = v161
	var v166 float64
	_ = v166
	var v167 float64
	_ = v167
	var v168 float64
	_ = v168
	var v193 float64
	_ = v193
	var v205 float64
	_ = v205
	var v227 float64
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int64
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if base.Ui32(int32(_a_F_numeric_out_sci_0)) <= base.Ui32(v12) {
		if v12 != int32(_a_F_numeric_out_sci_1) {
			if v12 != int32(_a_F_numeric_out_sci_2) {
				v28 = F_pstrdup(m, int32(_a_F_numeric_out_sci_3))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v296 = v28
					m.G0 = v10 - int32(-64)
					return v296
				}
			} else {
				v20 = F_pstrdup(m, int32(_a_F_numeric_out_sci_4))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v296 = v20
					m.G0 = v10 - int32(-64)
					return v296
				}
			}
		} else {
			v25 = F_pstrdup(m, int32(_a_F_numeric_out_sci_5))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v296 = v25
				m.G0 = v10 - int32(-64)
				return v296
			}
		}
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v35 = base.I32_extend16_s(v12)
		v37 = base.B2i32(int32(0) <= v35)
		if int32(0) <= v35 {
			v38 = int32(-8)
		} else {
			v38 = int32(-6)
		}
		v41 = int32(base.Ui32(int32(base.Ui32(v30)>>(uint(int32(2))%32))+v38) >> (uint(int32(1)) % 32))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v41
		if int32(0) <= v35 {
			v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
			v53 = v43
		} else {
			v53 = v12<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v12&int32(63)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v53
		v55 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v55
		v60 = base.B2i32(v35 < v55)
		if v35 < v55 {
			v61 = int32(6)
		} else {
			v61 = int32(8)
		}
		v62 = l0 + v61
		*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v62
		if v35 < v55 {
			v70 = int32(base.Ui32(v12)>>(uint(int32(7))%32)) & int32(63)
		} else {
			v70 = v12 & int32(_a_F_numeric_out_sci_6)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v70
		v77 = v12 & int32(_a_F_numeric_out_sci_0)
		if v77 == int32(_a_F_numeric_out_sci_7) {
			v80 = v12 << (uint(int32(1)) % 32) & int32(_a_F_numeric_out_sci_8)
		} else {
			v80 = v77
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v80
		if v41 != 0 {
			v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v62))))
			v84 = base.F64_convert_i32_s(v83)
			v95 = base.I64_reinterpret_f64(v84)
			if v95 <= int64(4503599627370495) {
				if base.F64_eq(v84, float64(0)) != 0 {
					v227 = base.F64_div(float64(-1), base.F64_mul(v84, v84))
				} else {
					if int64(0) <= v95 {
						v122 = base.I64_reinterpret_f64(base.F64_mul(v84, float64(1.8014398509481984e+16)))
						v126 = v122
						v128 = int32(-1077)
						v129 = base.I32_wrap_i64(int64(base.Ui64(v122) >> (uint(int64(32)) % 64)))
						v131 = v129 + int32(_a_F_numeric_out_sci_9)
						v135 = base.F64_convert_i32_s(int32(base.Ui32(v131)>>(uint(int32(20))%32)) + v128)
						v137 = base.F64_mul(v135, float64(0.30102999566361177))
						v150 = base.F64_add(base.F64_reinterpret_i64(v126&int64(4294967295)|base.I64_extend_i32_u(v131&int32(_a_F_numeric_out_sci_10)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
						v153 = base.F64_mul(v150, base.F64_mul(v150, float64(0.5)))
						v158 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v150, v153)) & int64(-4294967296))
						v159 = float64(0.4342944818781689)
						v160 = base.F64_mul(v158, v159)
						v161 = base.F64_add(v137, v160)
						v166 = base.F64_div(v150, base.F64_add(v150, float64(2)))
						v167 = base.F64_mul(v166, v166)
						v168 = base.F64_mul(v167, v167)
						v193 = base.F64_add(base.F64_mul(v166, base.F64_add(v153, base.F64_add(base.F64_mul(v168, base.F64_add(base.F64_mul(v168, base.F64_add(base.F64_mul(v168, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v167, base.F64_add(base.F64_mul(v168, base.F64_add(base.F64_mul(v168, base.F64_add(base.F64_mul(v168, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v150, v158), v153))
						v205 = base.F64_add(v161, base.F64_add(base.F64_add(v160, base.F64_sub(v137, v161)), base.F64_add(base.F64_mul(v193, v159), base.F64_add(base.F64_mul(v135, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v193, v158), float64(2.5082946711645275e-11))))))
						v227 = v205
					} else {
						v227 = base.F64_div(base.F64_sub(v84, v84), float64(0))
					}
				}
			} else {
				if base.Ui64(int64(9218868437227405311)) < base.Ui64(v95) {
					v205 = v84
					v227 = v205
				} else {
					v110 = int32(-1023)
					v112 = int64(base.Ui64(v95) >> (uint(int64(32)) % 64))
					if v112 != int64(1072693248) {
						v126 = v95
						v128 = v110
						v129 = base.I32_wrap_i64(v112)
						v131 = v129 + int32(_a_F_numeric_out_sci_9)
						v135 = base.F64_convert_i32_s(int32(base.Ui32(v131)>>(uint(int32(20))%32)) + v128)
						v137 = base.F64_mul(v135, float64(0.30102999566361177))
						v150 = base.F64_add(base.F64_reinterpret_i64(v126&int64(4294967295)|base.I64_extend_i32_u(v131&int32(_a_F_numeric_out_sci_10)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
						v153 = base.F64_mul(v150, base.F64_mul(v150, float64(0.5)))
						v158 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v150, v153)) & int64(-4294967296))
						v159 = float64(0.4342944818781689)
						v160 = base.F64_mul(v158, v159)
						v161 = base.F64_add(v137, v160)
						v166 = base.F64_div(v150, base.F64_add(v150, float64(2)))
						v167 = base.F64_mul(v166, v166)
						v168 = base.F64_mul(v167, v167)
						v193 = base.F64_add(base.F64_mul(v166, base.F64_add(v153, base.F64_add(base.F64_mul(v168, base.F64_add(base.F64_mul(v168, base.F64_add(base.F64_mul(v168, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v167, base.F64_add(base.F64_mul(v168, base.F64_add(base.F64_mul(v168, base.F64_add(base.F64_mul(v168, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v150, v158), v153))
						v205 = base.F64_add(v161, base.F64_add(base.F64_add(v160, base.F64_sub(v137, v161)), base.F64_add(base.F64_mul(v193, v159), base.F64_add(base.F64_mul(v135, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v193, v158), float64(2.5082946711645275e-11))))))
						v227 = v205
					} else {
						if base.I32_wrap_i64(v95) != 0 {
							v126 = v95
							v128 = v110
							v129 = int32(1072693248)
							v131 = v129 + int32(_a_F_numeric_out_sci_9)
							v135 = base.F64_convert_i32_s(int32(base.Ui32(v131)>>(uint(int32(20))%32)) + v128)
							v137 = base.F64_mul(v135, float64(0.30102999566361177))
							v150 = base.F64_add(base.F64_reinterpret_i64(v126&int64(4294967295)|base.I64_extend_i32_u(v131&int32(_a_F_numeric_out_sci_10)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
							v153 = base.F64_mul(v150, base.F64_mul(v150, float64(0.5)))
							v158 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v150, v153)) & int64(-4294967296))
							v159 = float64(0.4342944818781689)
							v160 = base.F64_mul(v158, v159)
							v161 = base.F64_add(v137, v160)
							v166 = base.F64_div(v150, base.F64_add(v150, float64(2)))
							v167 = base.F64_mul(v166, v166)
							v168 = base.F64_mul(v167, v167)
							v193 = base.F64_add(base.F64_mul(v166, base.F64_add(v153, base.F64_add(base.F64_mul(v168, base.F64_add(base.F64_mul(v168, base.F64_add(base.F64_mul(v168, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v167, base.F64_add(base.F64_mul(v168, base.F64_add(base.F64_mul(v168, base.F64_add(base.F64_mul(v168, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v150, v158), v153))
							v205 = base.F64_add(v161, base.F64_add(base.F64_add(v160, base.F64_sub(v137, v161)), base.F64_add(base.F64_mul(v193, v159), base.F64_add(base.F64_mul(v135, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v193, v158), float64(2.5082946711645275e-11))))))
							v227 = v205
						} else {
							v227 = float64(0)
						}
					}
				}
			}
			v232 = base.I32_trunc_sat_f64_s(v227) + v53<<(uint(int32(2))%32)
		} else {
			v232 = int32(0)
		}
		v234 = F_palloc(m, int32(4))
		mBase = m.M
		v235 = m.ExcPending
		if v235 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v234))) = int32(_a_F_numeric_out_sci_11)
			v239 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_out_sci[0]))
			*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v239
			v242 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_out_sci[1]))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v242
			v244 = int32(2)
			v245 = v234 + v244
			*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v245
			*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v234
			v248 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = (v248 - v232) & (v232 >> (uint(int32(31)) % 32))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v232 >> (uint(v244) % 32)
			if v248 < l1 {
				v260 = l1
			} else {
				v260 = v248
			}
			switch v232&int32(3) - int32(1) {
			case 0:
				v268 = int32(10)
				*(*uint16)(unsafe.Add(mBase, uint32(v245))) = uint16(v268)
			case 1:
				v268 = int32(100)
				*(*uint16)(unsafe.Add(mBase, uint32(v245))) = uint16(v268)
			case 2:
				v268 = int32(1000)
				*(*uint16)(unsafe.Add(mBase, uint32(v245))) = uint16(v268)
			default:
			}
			v274 = v8 + int32(-24)
			v275 = int32(1)
			F_div_var(m, v8+int32(-48), v274, v274, v260, v275, v275)
			mBase = m.M
			v278 = m.ExcPending
			if v278 != 0 {
				return int32(0)
			} else {
				v279 = F_get_str_from_var(m, v274)
				mBase = m.M
				v280 = m.ExcPending
				if v280 != 0 {
					return int32(0)
				} else {
					v281 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
					if v281 != 0 {
						F_pfree(m, v281)
						mBase = m.M
						v283 = m.ExcPending
						if v283 != 0 {
							return int32(0)
						} else {
							v284 = F_strlen(m, v279)
							mBase = m.M
							v286 = v284 + int32(13)
							v287 = F_palloc(m, v286)
							mBase = m.M
							v288 = m.ExcPending
							if v288 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v232
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v279
								v292 = F_pg_snprintf(m, v287, v286, int32(_a_F_numeric_out_sci_12), v10)
								mBase = m.M
								v293 = m.ExcPending
								if v293 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v279)
									mBase = m.M
									v295 = m.ExcPending
									if v295 != 0 {
										return int32(0)
									} else {
										v296 = v287
										m.G0 = v10 - int32(-64)
										return v296
									}
								}
							}
						}
					} else {
						v284 = F_strlen(m, v279)
						mBase = m.M
						v286 = v284 + int32(13)
						v287 = F_palloc(m, v286)
						mBase = m.M
						v288 = m.ExcPending
						if v288 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v232
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v279
							v292 = F_pg_snprintf(m, v287, v286, int32(_a_F_numeric_out_sci_12), v10)
							mBase = m.M
							v293 = m.ExcPending
							if v293 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v279)
								mBase = m.M
								v295 = m.ExcPending
								if v295 != 0 {
									return int32(0)
								} else {
									v296 = v287
									m.G0 = v10 - int32(-64)
									return v296
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_numeric_poly_sum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v25 int32
	_ = v25
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v48 int64
	_ = v48
	var v56 int64
	_ = v56
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v102 int64
	_ = v102
	var v105 int64
	_ = v105
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v116 int64
	_ = v116
	var v123 int64
	_ = v123
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v141 int64
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v17 != 0 {
		v25 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v25)
		v170 = int32(0)
		m.G0 = v15 - int32(-64)
		return v170
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v18 == int32(0) {
			v25 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v25)
			v170 = int32(0)
			m.G0 = v15 - int32(-64)
			return v170
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
			if v21 != int64(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = int64(0)
				v30 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
				v31 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
				v33 = F_palloc(m, int32(22))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v33
					v38 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v33))) = uint16(v38)
					*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v33 + int32(2)
					if v31 < int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = int64(16384)
						v48 = int64(0)
						v61 = v48 - v30
						v62 = v48 - (v31 + base.I64_extend_i32_u(base.B2i32(v30 != v48)))
						v65 = v38
						v68 = v33 + int32(22)
						v73 = v61
						v74 = v62
						for {
							v77 = int32(16)
							v78 = v13 + int32(-48)
							v81 = m.G0
							v83 = v81 - v77
							m.G0 = v83
							F___udivmodti4(m, v83, v73, v74, int64(10000), int64(0))
							mBase = m.M
							v87 = *(*int64)(unsafe.Add(mBase, uint32(v83)+8))
							v88 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
							*(*int64)(unsafe.Add(mBase, uint32(v78))) = v88
							*(*int64)(unsafe.Add(mBase, uint32(v78)+8)) = v87
							m.G0 = v83 + v77
							v94 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
							v95 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
							v96 = int64(55536)
							v97 = int64(0)
							v102 = int64(32)
							v105 = int64(base.Ui64(v94) >> (uint(v102) % 64))
							v108 = int64(4294967295)
							v111 = v94 & v108
							v112 = v96 * v111
							v116 = int64(base.Ui64(v112)>>(uint(v102)%64)) + v96*v105
							v123 = v111*v97 + v116&v108
							*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v94*v97 + v95*v96 + v97*v105 + int64(base.Ui64(v116)>>(uint(v102)%64)) + int64(base.Ui64(v123)>>(uint(v102)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v15))) = v112&v108 | v123<<(uint(v102)%64)
							v135 = v68 - int32(2)
							v136 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							v137 = v136 + v73
							*(*uint16)(unsafe.Add(mBase, uint32(v135))) = uint16(v137)
							v141 = int64(0)
							v146 = v65 + int32(1)
							if v74 == v141 {
								v147 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v73))
							} else {
								v147 = base.B2i32(v74 != v141)
							}
							if v147 != 0 {
								v65 = v146
								v68 = v135
								v73 = v94
								v74 = v95
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v135
						v149 = v146
						v156 = v65
					} else {
						v56 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v56
						if v30|v31 == v56 {
							v149 = v38
							v156 = int32(0)
						} else {
							v61 = v30
							v62 = v31
							v65 = v38
							v68 = v33 + int32(22)
							v73 = v61
							v74 = v62
							for {
								v77 = int32(16)
								v78 = v13 + int32(-48)
								v81 = m.G0
								v83 = v81 - v77
								m.G0 = v83
								F___udivmodti4(m, v83, v73, v74, int64(10000), int64(0))
								mBase = m.M
								v87 = *(*int64)(unsafe.Add(mBase, uint32(v83)+8))
								v88 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
								*(*int64)(unsafe.Add(mBase, uint32(v78))) = v88
								*(*int64)(unsafe.Add(mBase, uint32(v78)+8)) = v87
								m.G0 = v83 + v77
								v94 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
								v95 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
								v96 = int64(55536)
								v97 = int64(0)
								v102 = int64(32)
								v105 = int64(base.Ui64(v94) >> (uint(v102) % 64))
								v108 = int64(4294967295)
								v111 = v94 & v108
								v112 = v96 * v111
								v116 = int64(base.Ui64(v112)>>(uint(v102)%64)) + v96*v105
								v123 = v111*v97 + v116&v108
								*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v94*v97 + v95*v96 + v97*v105 + int64(base.Ui64(v116)>>(uint(v102)%64)) + int64(base.Ui64(v123)>>(uint(v102)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v15))) = v112&v108 | v123<<(uint(v102)%64)
								v135 = v68 - int32(2)
								v136 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
								v137 = v136 + v73
								*(*uint16)(unsafe.Add(mBase, uint32(v135))) = uint16(v137)
								v141 = int64(0)
								v146 = v65 + int32(1)
								if v74 == v141 {
									v147 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v73))
								} else {
									v147 = base.B2i32(v74 != v141)
								}
								if v147 != 0 {
									v65 = v146
									v68 = v135
									v73 = v94
									v74 = v95
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v135
							v149 = v146
							v156 = v65
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v156
					*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v149
					v166 = F_make_result_opt_error(m, v13+int32(-24), int32(0))
					mBase = m.M
					v167 = m.ExcPending
					if v167 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v33)
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return int32(0)
						} else {
							v170 = v166
							m.G0 = v15 - int32(-64)
							return v170
						}
					}
				}
			} else {
				v25 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v25)
				v170 = int32(0)
				m.G0 = v15 - int32(-64)
				return v170
			}
		}
	}
}
func F_numeric_poly_var_samp(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(1)
	v4 = Fn13952(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_numeric_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v122 int32
	_ = v122
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pq_getmsgint(m, v18, int32(2))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = v20 & int32(_a_F_numeric_recv_0)
	v30 = F_palloc(m, v25<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v30
	v33 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v30))) = uint16(v33)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v25
	v36 = int32(2)
	v37 = v30 + v36
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v37
	v40 = F_pq_getmsgint(m, v18, v36)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = base.I32_extend16_s(v40)
	v45 = F_pq_getmsgint(m, v18, int32(2))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v48 = v45 & int32(_a_F_numeric_recv_0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v48
	v52 = int32(0)
	if base.B2i32(base.B2i32(v45&int32(_a_F_numeric_recv_1) == v52)|base.B2i32(v48 == int32(_a_F_numeric_recv_2)) == v52)&base.B2i32(v45&int32(_a_F_numeric_recv_3) != int32(_a_F_numeric_recv_4)) == v52 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L53
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L49
	}
L8:
	;
	v67 = F_pq_getmsgint(m, v18, int32(2))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L45
	}
L11:
	;
	v70 = v67 & int32(_a_F_numeric_recv_0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v70
	if v67&int32(_a_F_numeric_recv_4) != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v25 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v75 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	if v45&int32(_a_F_numeric_recv_5) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v88 = F_pq_getmsgint(m, v18, int32(2))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L15
L18:
	;
	if base.Ui32(int32(_a_F_numeric_recv_6)) <= base.Ui32(v88&int32(_a_F_numeric_recv_0)) {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v94 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v37+v75<<(uint(v94)%32)))) = uint16(v88)
	v99 = v75 + v94
	if v99 != v25 {
		v75 = v99
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	m.G0 = v15 + int32(32)
	return v212
L22:
	;
	F_pfree(m, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L44
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v70
	v122 = v40<<(uint(int32(16))%32)>>(uint(int32(14))%32) + v70
	if v122+int32(4) <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	v167 = F_make_result_opt_error(m, v15+int32(8), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L35
	}
L26:
	;
	v156 = v15 + int32(8)
	v158 = F_apply_typmod(m, v156, v17, int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L32
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = int64(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v134 = int32(base.Ui32(v122+int32(7)) >> (uint(int32(2)) % 32))
	if base.Ui32(v25) < base.Ui32(v134) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v134
	v138 = v67 & int32(3)
	if v138 == int32(0) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v144 = int32(2)
	v145 = v37 + v134<<(uint(int32(1))%32) - v144
	v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(v145))))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v138<<(uint(v144)%32))+uint32(_c_F_numeric_recv[0])))
	v150 = base.I32_rem_s(v146, v149)
	v151 = v146 - v150
	*(*uint16)(unsafe.Add(mBase, uint32(v145))) = uint16(v151)
	goto L26
L32:
	;
	v161 = F_make_result_opt_error(m, v156, int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	if v163 != 0 {
		v207 = v161
		v209 = v163
		goto L22
	} else {
		goto L34
	}
L34:
	;
	v212 = v161
	goto L21
L35:
	;
	if v17 < int32(4) {
		v207 = v167
		v209 = v30
		goto L22
	} else {
		goto L36
	}
L36:
	;
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+4)))
	if v171 == int32(_a_F_numeric_recv_4) {
		v207 = v167
		v209 = v30
		goto L22
	} else {
		goto L37
	}
L37:
	;
	v175 = F_errsave_start(m, int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v175 == int32(0) {
		v207 = v167
		v209 = v30
		goto L22
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(_a_F_numeric_recv_7), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(base.Ui32(v17-int32(4)) >> (uint(int32(16)) % 32))
	v191 = int32(21)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = (v17<<(uint(v191)%32) - int32(_a_F_numeric_recv_8)) >> (uint(v191) % 32)
	F_errdetail(m, int32(_a_F_numeric_recv_9), v15)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errsave_finish(m, int32(0), int32(_a_F_numeric_recv_10), int32(_a_F_numeric_recv_11), int32(_a_F_numeric_recv_12))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v207 = v167
	v209 = v30
	goto L22
L44:
	;
	v212 = v207
	goto L21
L45:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(_a_F_numeric_recv_13), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_numeric_recv_10), int32(1108), int32(_a_F_numeric_recv_14))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_numeric_recv_15), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_numeric_recv_10), int32(1114), int32(_a_F_numeric_recv_14))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(_a_F_numeric_recv_16), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_numeric_recv_10), int32(1123), int32(_a_F_numeric_recv_14))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_numeric_scale(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+4)))
		v9 = int32(_a_F_numeric_scale_0)
		if v8&v9 == v9 {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
		} else {
			if base.I32_extend16_s(v8) < int32(0) {
				v26 = int32(base.Ui32(v8)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v26 = v8 & int32(_a_F_numeric_scale_1)
			}
			return v26
		}
	}
}
func F_numeric_serialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v135 int64
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int64
	_ = v142
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v149 int64
	_ = v149
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v178 int32
	_ = v178
	var v181 int64
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int64
	_ = v188
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v195 int64
	_ = v195
	var v197 int64
	_ = v197
	var v199 int64
	_ = v199
	var v201 int64
	_ = v201
	var v224 int32
	_ = v224
	var v227 int64
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int64
	_ = v234
	var v236 int64
	_ = v236
	var v238 int64
	_ = v238
	var v241 int64
	_ = v241
	var v243 int64
	_ = v243
	var v245 int64
	_ = v245
	var v247 int64
	_ = v247
	var v270 int32
	_ = v270
	var v273 int64
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int64
	_ = v280
	var v282 int64
	_ = v282
	var v284 int64
	_ = v284
	var v287 int64
	_ = v287
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v293 int64
	_ = v293
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 == int32(0) {
		v41 = int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		switch v16 - int32(429) {
		case 0:
			v41 = int32(1)
		case 1:
			v41 = int32(2)
		default:
			v41 = int32(0)
		}
	}
	if v41 != 0 {
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v43 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v43
		*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v43
		*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v43
		v50 = v9 + int32(32)
		F_pq_begintypsend(m, v50)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int32(0)
		} else {
			v55 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
			F_enlargeStringInfo(m, v50, int32(8))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
				v62 = int64(56)
				v64 = int64(65280)
				v66 = int64(40)
				v69 = int64(16711680)
				v71 = int64(24)
				v73 = int64(4278190080)
				v75 = int64(8)
				*(*int64)(unsafe.Add(mBase, uint32(v59+v60))) = v55<<(uint(v62)%64) | v55&v64<<(uint(v66)%64) | (v55&v69<<(uint(v71)%64) | v55&v73<<(uint(v75)%64)) | (int64(base.Ui64(v55)>>(uint(v75)%64))&v73 | int64(base.Ui64(v55)>>(uint(v71)%64))&v69 | (int64(base.Ui64(v55)>>(uint(v66)%64))&v64 | int64(base.Ui64(v55)>>(uint(v62)%64))))
				v98 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v59 + v98
				v104 = v9 + v98
				F_accum_sum_final(m, v42+int32(16), v104)
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					F_numericvar_serialize(m, v50, v104)
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						F_accum_sum_final(m, v42+int32(44), v104)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							F_numericvar_serialize(m, v50, v104)
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int32(0)
							} else {
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v42)+72))
								F_enlargeStringInfo(m, v50, int32(4))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
									v120 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
									v124 = int32(16711935)
									v128 = int32(8)
									*(*int32)(unsafe.Add(mBase, uint32(v119+v120))) = base.I32_rotr(v115, int32(24))&v124 | base.I32_rotr(v115&v124, v128)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v119 + int32(4)
									v135 = *(*int64)(unsafe.Add(mBase, uint32(v42)+80))
									F_enlargeStringInfo(m, v50, v128)
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return int32(0)
									} else {
										v139 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
										v140 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
										v142 = int64(56)
										v144 = int64(65280)
										v146 = int64(40)
										v149 = int64(16711680)
										v151 = int64(24)
										v153 = int64(4278190080)
										v155 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v139+v140))) = v135<<(uint(v142)%64) | v135&v144<<(uint(v146)%64) | (v135&v149<<(uint(v151)%64) | v135&v153<<(uint(v155)%64)) | (int64(base.Ui64(v135)>>(uint(v155)%64))&v153 | int64(base.Ui64(v135)>>(uint(v151)%64))&v149 | (int64(base.Ui64(v135)>>(uint(v146)%64))&v144 | int64(base.Ui64(v135)>>(uint(v142)%64))))
										v178 = int32(8)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v139 + v178
										v181 = *(*int64)(unsafe.Add(mBase, uint32(v42)+88))
										F_enlargeStringInfo(m, v50, v178)
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
											return int32(0)
										} else {
											v185 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
											v186 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
											v188 = int64(56)
											v190 = int64(65280)
											v192 = int64(40)
											v195 = int64(16711680)
											v197 = int64(24)
											v199 = int64(4278190080)
											v201 = int64(8)
											*(*int64)(unsafe.Add(mBase, uint32(v185+v186))) = v181<<(uint(v188)%64) | v181&v190<<(uint(v192)%64) | (v181&v195<<(uint(v197)%64) | v181&v199<<(uint(v201)%64)) | (int64(base.Ui64(v181)>>(uint(v201)%64))&v199 | int64(base.Ui64(v181)>>(uint(v197)%64))&v195 | (int64(base.Ui64(v181)>>(uint(v192)%64))&v190 | int64(base.Ui64(v181)>>(uint(v188)%64))))
											v224 = int32(8)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v185 + v224
											v227 = *(*int64)(unsafe.Add(mBase, uint32(v42)+96))
											F_enlargeStringInfo(m, v50, v224)
											mBase = m.M
											v230 = m.ExcPending
											if v230 != 0 {
												return int32(0)
											} else {
												v231 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
												v232 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
												v234 = int64(56)
												v236 = int64(65280)
												v238 = int64(40)
												v241 = int64(16711680)
												v243 = int64(24)
												v245 = int64(4278190080)
												v247 = int64(8)
												*(*int64)(unsafe.Add(mBase, uint32(v231+v232))) = v227<<(uint(v234)%64) | v227&v236<<(uint(v238)%64) | (v227&v241<<(uint(v243)%64) | v227&v245<<(uint(v247)%64)) | (int64(base.Ui64(v227)>>(uint(v247)%64))&v245 | int64(base.Ui64(v227)>>(uint(v243)%64))&v241 | (int64(base.Ui64(v227)>>(uint(v238)%64))&v236 | int64(base.Ui64(v227)>>(uint(v234)%64))))
												v270 = int32(8)
												*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v231 + v270
												v273 = *(*int64)(unsafe.Add(mBase, uint32(v42)+104))
												F_enlargeStringInfo(m, v50, v270)
												mBase = m.M
												v276 = m.ExcPending
												if v276 != 0 {
													return int32(0)
												} else {
													v277 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
													v278 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
													v280 = int64(56)
													v282 = int64(65280)
													v284 = int64(40)
													v287 = int64(16711680)
													v289 = int64(24)
													v291 = int64(4278190080)
													v293 = int64(8)
													*(*int64)(unsafe.Add(mBase, uint32(v277+v278))) = v273<<(uint(v280)%64) | v273&v282<<(uint(v284)%64) | (v273&v287<<(uint(v289)%64) | v273&v291<<(uint(v293)%64)) | (int64(base.Ui64(v273)>>(uint(v293)%64))&v291 | int64(base.Ui64(v273)>>(uint(v289)%64))&v287 | (int64(base.Ui64(v273)>>(uint(v284)%64))&v282 | int64(base.Ui64(v273)>>(uint(v280)%64))))
													*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v277 + int32(8)
													v320 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
													v321 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v320))) = v321 << (uint(int32(2)) % 32)
													v325 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
													if v325 != 0 {
														F_pfree(m, v325)
														mBase = m.M
														v327 = m.ExcPending
														if v327 != 0 {
															return int32(0)
														} else {
															m.G0 = v9 + int32(48)
															return v320
														}
													} else {
														m.G0 = v9 + int32(48)
														return v320
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v335 = m.ExcPending
		if v335 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_serialize_0), int32(0))
			mBase = m.M
			v339 = m.ExcPending
			if v339 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_serialize_1), int32(_a_F_numeric_serialize_2), int32(_a_F_numeric_serialize_3))
				mBase = m.M
				v344 = m.ExcPending
				if v344 != 0 {
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
func F_numeric_smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = F_pg_detoast_datum(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+4)))
			v24 = int32(_a_F_numeric_smaller_0)
			v25 = v23 & v24
			if v25 == v24 {
				if v23 != int32(_a_F_numeric_smaller_1) {
					if v23 != int32(_a_F_numeric_smaller_0) {
						if v22 != int32(_a_F_numeric_smaller_2) {
							v44 = int32(-1)
						} else {
							v44 = int32(0)
						}
						v161 = v44
					} else {
						v161 = base.B2i32(v22 != int32(_a_F_numeric_smaller_0))
					}
				} else {
					if v22 == int32(_a_F_numeric_smaller_0) {
						v39 = int32(-1)
					} else {
						v39 = base.B2i32(v22 != int32(_a_F_numeric_smaller_1))
					}
					v161 = v39
				}
			} else {
				if base.Ui32(int32(_a_F_numeric_smaller_0)) <= base.Ui32(v22) {
					if v22 == int32(_a_F_numeric_smaller_2) {
						v51 = int32(1)
					} else {
						v51 = int32(-1)
					}
					v161 = v51
				} else {
					v53 = v4 + int32(6)
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
					v61 = base.B2i32(int32(0) <= base.I32_extend16_s(v23))
					if int32(0) <= base.I32_extend16_s(v23) {
						v62 = int32(-8)
					} else {
						v62 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v23) {
						v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
						v74 = v64
					} else {
						v74 = v23<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v23&int32(63)
					}
					v76 = int32(base.Ui32(int32(base.Ui32(v54)>>(uint(int32(2))%32))+v62) >> (uint(int32(1)) % 32))
					v78 = v9 + int32(6)
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v86 = base.B2i32(int32(0) <= base.I32_extend16_s(v22))
					if int32(0) <= base.I32_extend16_s(v22) {
						v87 = int32(-8)
					} else {
						v87 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v22) {
						v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v78))))
						v99 = v89
					} else {
						v99 = v22<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v22&int32(63)
					}
					v100 = int32(1)
					v101 = int32(base.Ui32(int32(base.Ui32(v79)>>(uint(int32(2))%32))+v87) >> (uint(v100) % 32))
					v107 = v22 & int32(_a_F_numeric_smaller_0)
					if v107 == int32(_a_F_numeric_smaller_3) {
						v110 = v22 << (uint(v100) % 32) & int32(_a_F_numeric_smaller_4)
					} else {
						v110 = v107
					}
					if v76 == int32(0) {
						if v101 == int32(0) {
							v161 = int32(0)
						} else {
							if v110 == int32(_a_F_numeric_smaller_4) {
								v120 = int32(1)
							} else {
								v120 = int32(-1)
							}
							v161 = v120
						}
					} else {
						if v25 == int32(_a_F_numeric_smaller_3) {
							v127 = v23 << (uint(int32(1)) % 32) & int32(_a_F_numeric_smaller_4)
						} else {
							v127 = v25
						}
						if v101 == int32(0) {
							if v127 != 0 {
								v132 = int32(-1)
							} else {
								v132 = int32(1)
							}
							v161 = v132
						} else {
							if int32(0) <= base.I32_extend16_s(v23) {
								v135 = v4 + int32(8)
							} else {
								v135 = v53
							}
							if int32(0) <= base.I32_extend16_s(v22) {
								v138 = v9 + int32(8)
							} else {
								v138 = v78
							}
							if v127 == int32(0) {
								if v110 == int32(_a_F_numeric_smaller_4) {
									v161 = int32(1)
								} else {
									v144 = F_cmp_abs_common(m, v135, v76, v74, v138, v101, v99)
									mBase = m.M
									v161 = v144
								}
							} else {
								if v110 == int32(0) {
									v161 = int32(-1)
								} else {
									v148 = F_cmp_abs_common(m, v138, v101, v99, v135, v76, v74)
									mBase = m.M
									v161 = v148
								}
							}
						}
					}
				}
			}
			if v161 < int32(0) {
				v164 = v4
			} else {
				v164 = v9
			}
			return v164
		}
	}
}
func F_numeric_stddev_samp(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13953(m, l0, int32(1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
