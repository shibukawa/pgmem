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
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v242 int32
	_ = v242
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if base.Ui32(int32(49152)) <= base.Ui32(v12) {
		if v12 != int32(61440) {
			if v12 != int32(53248) {
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
		v242 = int32(1)
		m.G0 = v10 + int32(48)
		return v242
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
			v66 = v12 & int32(16383)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v66
		v73 = v12 & int32(49152)
		if v73 == int32(32768) {
			v76 = v12 << (uint(int32(1)) % 32) & int32(16384)
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
					v122 = m.ExcPending
					if v122 != 0 {
						return int32(0)
					} else {
						v123 = int32(4548768)
						v124 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v126
						v128 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v128 - int64(1)
						if int64(2) <= v128 {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = base.B2i32(v76 == int32(0)) << (uint(int32(14)) % 32)
							F_accum_sum_add(m, l0+int32(16), v10+int32(24))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return int32(0)
							} else {
								v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
								if v145 != int32(1) {
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
									v242 = int32(1)
									m.G0 = v10 + int32(48)
									return v242
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(16384)
									F_accum_sum_add(m, l0+int32(44), v10)
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
										v242 = int32(1)
										m.G0 = v10 + int32(48)
										return v242
									}
								}
							}
						} else {
							v154 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v154
							v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							if v154 < v156 {
								v162 = int32(0)
								for {
									v168 = v162 << (uint(int32(2)) % 32)
									v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v171 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v168+v169))) = v171
									v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									*(*int32)(unsafe.Add(mBase, uint32(v173+v168))) = v171
									v178 = v162 + int32(1)
									v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v178 < v179 {
										v162 = v178
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
							if v188 != int32(1) {
							} else {
								v191 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v191
								v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								if v193 <= v191 {
								} else {
									v199 = int32(0)
									for {
										v205 = v199 << (uint(int32(2)) % 32)
										v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
										v208 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v205+v206))) = v208
										v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
										*(*int32)(unsafe.Add(mBase, uint32(v210+v205))) = v208
										v215 = v199 + int32(1)
										v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										if v215 < v216 {
											v199 = v215
											continue
										} else {
											break
										}
										break
									}
								}
							}
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
							v242 = int32(1)
							m.G0 = v10 + int32(48)
							return v242
						}
					}
				} else {
					v123 = int32(4548768)
					v124 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v126
					v128 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v128 - int64(1)
					if int64(2) <= v128 {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = base.B2i32(v76 == int32(0)) << (uint(int32(14)) % 32)
						F_accum_sum_add(m, l0+int32(16), v10+int32(24))
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int32(0)
						} else {
							v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
							if v145 != int32(1) {
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
								v242 = int32(1)
								m.G0 = v10 + int32(48)
								return v242
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(16384)
								F_accum_sum_add(m, l0+int32(44), v10)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
									v242 = int32(1)
									m.G0 = v10 + int32(48)
									return v242
								}
							}
						}
					} else {
						v154 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v154
						v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v154 < v156 {
							v162 = int32(0)
							for {
								v168 = v162 << (uint(int32(2)) % 32)
								v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								v171 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v168+v169))) = v171
								v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v173+v168))) = v171
								v178 = v162 + int32(1)
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v178 < v179 {
									v162 = v178
									continue
								} else {
									break
								}
								break
							}
						} else {
						}
						v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
						if v188 != int32(1) {
						} else {
							v191 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v191
							v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v193 <= v191 {
							} else {
								v199 = int32(0)
								for {
									v205 = v199 << (uint(int32(2)) % 32)
									v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
									v208 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v205+v206))) = v208
									v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
									*(*int32)(unsafe.Add(mBase, uint32(v210+v205))) = v208
									v215 = v199 + int32(1)
									v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									if v215 < v216 {
										v199 = v215
										continue
									} else {
										break
									}
									break
								}
							}
						}
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
						v242 = int32(1)
						m.G0 = v10 + int32(48)
						return v242
					}
				}
			} else {
				v95 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				if v95 != int64(1) {
					v242 = int32(0)
					m.G0 = v10 + int32(48)
					return v242
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
						v122 = m.ExcPending
						if v122 != 0 {
							return int32(0)
						} else {
							v123 = int32(4548768)
							v124 = *(*int32)(unsafe.Add(mBase, _consts[0]))
							v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v126
							v128 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v128 - int64(1)
							if int64(2) <= v128 {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = base.B2i32(v76 == int32(0)) << (uint(int32(14)) % 32)
								F_accum_sum_add(m, l0+int32(16), v10+int32(24))
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return int32(0)
								} else {
									v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
									if v145 != int32(1) {
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
										v242 = int32(1)
										m.G0 = v10 + int32(48)
										return v242
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(16384)
										F_accum_sum_add(m, l0+int32(44), v10)
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
											v242 = int32(1)
											m.G0 = v10 + int32(48)
											return v242
										}
									}
								}
							} else {
								v154 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v154
								v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v154 < v156 {
									v162 = int32(0)
									for {
										v168 = v162 << (uint(int32(2)) % 32)
										v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										v171 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v168+v169))) = v171
										v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										*(*int32)(unsafe.Add(mBase, uint32(v173+v168))) = v171
										v178 = v162 + int32(1)
										v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v178 < v179 {
											v162 = v178
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
								if v188 != int32(1) {
								} else {
									v191 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v191
									v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									if v193 <= v191 {
									} else {
										v199 = int32(0)
										for {
											v205 = v199 << (uint(int32(2)) % 32)
											v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
											v208 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v205+v206))) = v208
											v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											*(*int32)(unsafe.Add(mBase, uint32(v210+v205))) = v208
											v215 = v199 + int32(1)
											v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											if v215 < v216 {
												v199 = v215
												continue
											} else {
												break
											}
											break
										}
									}
								}
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
								v242 = int32(1)
								m.G0 = v10 + int32(48)
								return v242
							}
						}
					} else {
						v123 = int32(4548768)
						v124 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v126
						v128 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v128 - int64(1)
						if int64(2) <= v128 {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = base.B2i32(v76 == int32(0)) << (uint(int32(14)) % 32)
							F_accum_sum_add(m, l0+int32(16), v10+int32(24))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return int32(0)
							} else {
								v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
								if v145 != int32(1) {
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
									v242 = int32(1)
									m.G0 = v10 + int32(48)
									return v242
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(16384)
									F_accum_sum_add(m, l0+int32(44), v10)
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
										v242 = int32(1)
										m.G0 = v10 + int32(48)
										return v242
									}
								}
							}
						} else {
							v154 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v154
							v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							if v154 < v156 {
								v162 = int32(0)
								for {
									v168 = v162 << (uint(int32(2)) % 32)
									v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v171 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v168+v169))) = v171
									v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									*(*int32)(unsafe.Add(mBase, uint32(v173+v168))) = v171
									v178 = v162 + int32(1)
									v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v178 < v179 {
										v162 = v178
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
							if v188 != int32(1) {
							} else {
								v191 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v191
								v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								if v193 <= v191 {
								} else {
									v199 = int32(0)
									for {
										v205 = v199 << (uint(int32(2)) % 32)
										v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
										v208 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v205+v206))) = v208
										v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
										*(*int32)(unsafe.Add(mBase, uint32(v210+v205))) = v208
										v215 = v199 + int32(1)
										v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										if v215 < v216 {
											v199 = v215
											continue
										} else {
											break
										}
										break
									}
								}
							}
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
							v242 = int32(1)
							m.G0 = v10 + int32(48)
							return v242
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
				v122 = m.ExcPending
				if v122 != 0 {
					return int32(0)
				} else {
					v123 = int32(4548768)
					v124 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v126
					v128 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v128 - int64(1)
					if int64(2) <= v128 {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = base.B2i32(v76 == int32(0)) << (uint(int32(14)) % 32)
						F_accum_sum_add(m, l0+int32(16), v10+int32(24))
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int32(0)
						} else {
							v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
							if v145 != int32(1) {
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
								v242 = int32(1)
								m.G0 = v10 + int32(48)
								return v242
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(16384)
								F_accum_sum_add(m, l0+int32(44), v10)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
									v242 = int32(1)
									m.G0 = v10 + int32(48)
									return v242
								}
							}
						}
					} else {
						v154 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v154
						v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v154 < v156 {
							v162 = int32(0)
							for {
								v168 = v162 << (uint(int32(2)) % 32)
								v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								v171 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v168+v169))) = v171
								v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v173+v168))) = v171
								v178 = v162 + int32(1)
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v178 < v179 {
									v162 = v178
									continue
								} else {
									break
								}
								break
							}
						} else {
						}
						v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
						if v188 != int32(1) {
						} else {
							v191 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v191
							v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v193 <= v191 {
							} else {
								v199 = int32(0)
								for {
									v205 = v199 << (uint(int32(2)) % 32)
									v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
									v208 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v205+v206))) = v208
									v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
									*(*int32)(unsafe.Add(mBase, uint32(v210+v205))) = v208
									v215 = v199 + int32(1)
									v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									if v215 < v216 {
										v199 = v215
										continue
									} else {
										break
									}
									break
								}
							}
						}
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
						v242 = int32(1)
						m.G0 = v10 + int32(48)
						return v242
					}
				}
			} else {
				v123 = int32(4548768)
				v124 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v126
				v128 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v128 - int64(1)
				if int64(2) <= v128 {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = base.B2i32(v76 == int32(0)) << (uint(int32(14)) % 32)
					F_accum_sum_add(m, l0+int32(16), v10+int32(24))
					mBase = m.M
					v144 = m.ExcPending
					if v144 != 0 {
						return int32(0)
					} else {
						v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
						if v145 != int32(1) {
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
							v242 = int32(1)
							m.G0 = v10 + int32(48)
							return v242
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(16384)
							F_accum_sum_add(m, l0+int32(44), v10)
							mBase = m.M
							v153 = m.ExcPending
							if v153 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
								v242 = int32(1)
								m.G0 = v10 + int32(48)
								return v242
							}
						}
					}
				} else {
					v154 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v154
					v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v154 < v156 {
						v162 = int32(0)
						for {
							v168 = v162 << (uint(int32(2)) % 32)
							v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v171 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v168+v169))) = v171
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v173+v168))) = v171
							v178 = v162 + int32(1)
							v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							if v178 < v179 {
								v162 = v178
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v188 != int32(1) {
					} else {
						v191 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v191
						v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						if v193 <= v191 {
						} else {
							v199 = int32(0)
							for {
								v205 = v199 << (uint(int32(2)) % 32)
								v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
								v208 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v205+v206))) = v208
								v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
								*(*int32)(unsafe.Add(mBase, uint32(v210+v205))) = v208
								v215 = v199 + int32(1)
								v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								if v215 < v216 {
									v199 = v215
									continue
								} else {
									break
								}
								break
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
					v242 = int32(1)
					m.G0 = v10 + int32(48)
					return v242
				}
			}
		}
	}
}
func F_numeric_abs(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v11 = F_palloc(m, int32(base.Ui32(v8)>>(uint(int32(2))%32)))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			v15 = int32(base.Ui32(v13) >> (uint(int32(2)) % 32))
			if v15 != 0 {
				v16 = F__emscripten_memcpy_bulkmem(m, v11, v4, v15)
				mBase = m.M
				v17 = v16
			} else {
				v17 = v11
			}
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+4)))
			if base.Ui32(int32(49151)) < base.Ui32(v18) {
				v24 = int32(-8193)
			} else {
				v24 = int32(16383)
			}
			if base.I32_extend16_s(v18) < int32(-16384) {
				v28 = int32(-24577)
			} else {
				v28 = v24
			}
			v29 = v18 & v28
			*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)) = uint16(v29)
			return v17
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v13 != 0 {
		v31 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
		v97 = int32(0)
		m.G0 = v11 + int32(32)
		return v97
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v14 == int32(0) {
			v31 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
			v97 = int32(0)
			m.G0 = v11 + int32(32)
			return v97
		} else {
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v14)+96))
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v14)+88))
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v14)+104))
			if v17+(v18+v19) != int64(0)-v23 {
				if int64(0) < v18 {
					v38 = F_make_result_opt_error(m, int32(1766380), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v97 = v38
						m.G0 = v11 + int32(32)
						return v97
					}
				} else {
					if v17 <= int64(0) {
						if int64(0) < v17 {
							v54 = F_make_result_opt_error(m, int32(1766404), int32(0))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v97 = v54
								m.G0 = v11 + int32(32)
								return v97
							}
						} else {
							if int64(0) < v23 {
								v60 = F_make_result_opt_error(m, int32(1766428), int32(0))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v97 = v60
									m.G0 = v11 + int32(32)
									return v97
								}
							} else {
								v62 = F_int64_to_numeric(m, v19)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v65 = v11 + int32(24)
									v66 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v65))) = v66
									*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v66
									*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v66
									F_accum_sum_final(m, v14+int32(16), v11+int32(8))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										v81 = F_make_result_opt_error(m, v11+int32(8), int32(0))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
											if v83 != 0 {
												F_pfree(m, v83)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int32(0)
												} else {
													v88 = F_DirectFunctionCall2Coll(m, int32(1276), int32(0), v81, v62)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int32(0)
													} else {
														v97 = v88
														m.G0 = v11 + int32(32)
														return v97
													}
												}
											} else {
												v88 = F_DirectFunctionCall2Coll(m, int32(1276), int32(0), v81, v62)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													v97 = v88
													m.G0 = v11 + int32(32)
													return v97
												}
											}
										}
									}
								}
							}
						}
					} else {
						if v23 <= int64(0) {
							if int64(0) < v17 {
								v54 = F_make_result_opt_error(m, int32(1766404), int32(0))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v97 = v54
									m.G0 = v11 + int32(32)
									return v97
								}
							} else {
								if int64(0) < v23 {
									v60 = F_make_result_opt_error(m, int32(1766428), int32(0))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v97 = v60
										m.G0 = v11 + int32(32)
										return v97
									}
								} else {
									v62 = F_int64_to_numeric(m, v19)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v65 = v11 + int32(24)
										v66 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v65))) = v66
										*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v66
										*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v66
										F_accum_sum_final(m, v14+int32(16), v11+int32(8))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											v81 = F_make_result_opt_error(m, v11+int32(8), int32(0))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
												if v83 != 0 {
													F_pfree(m, v83)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return int32(0)
													} else {
														v88 = F_DirectFunctionCall2Coll(m, int32(1276), int32(0), v81, v62)
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return int32(0)
														} else {
															v97 = v88
															m.G0 = v11 + int32(32)
															return v97
														}
													}
												} else {
													v88 = F_DirectFunctionCall2Coll(m, int32(1276), int32(0), v81, v62)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int32(0)
													} else {
														v97 = v88
														m.G0 = v11 + int32(32)
														return v97
													}
												}
											}
										}
									}
								}
							}
						} else {
							v48 = F_make_result_opt_error(m, int32(1766380), int32(0))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								v97 = v48
								m.G0 = v11 + int32(32)
								return v97
							}
						}
					}
				}
			} else {
				v31 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
				v97 = int32(0)
				m.G0 = v11 + int32(32)
				return v97
			}
		}
	}
}
func F_numeric_avg_accum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v9 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v12 != 0 {
			v64 = v12
			v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v66 == int32(0) {
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v70 = F_pg_detoast_datum(m, v69)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_do_numeric_accum(m, v64, v70)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v64
					}
				}
			} else {
				m.G0 = v7 + int32(16)
				return v64
			}
		} else {
			v15 = v7 + int32(12)
			v16 = int32(0)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v17 == v16 {
				v34 = int32(0)
				if v15 == v34 {
					v42 = v34
				} else {
					v37 = v34
					v38 = v16
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
					v42 = v38
				}
				v45 = v42
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				switch v20 - int32(429) {
				case 0:
					if v15 == int32(0) {
						v45 = int32(1)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
						v37 = v27
						v38 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
						v42 = v38
						v45 = v42
					}
				case 1:
					if v15 == int32(0) {
						v45 = int32(2)
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+368))
						v37 = v32
						v38 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
						v42 = v38
						v45 = v42
					}
				default:
					v34 = int32(0)
					if v15 == v34 {
						v42 = v34
					} else {
						v37 = v34
						v38 = v16
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
						v42 = v38
					}
					v45 = v42
				}
			}
			if v45 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(66285), int32(0))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(520410), int32(4943), int32(368516))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v48 = int32(4548768)
				v49 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v51
				v54 = F_palloc0(m, int32(112))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					v58 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v58)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v60
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v49
					v64 = v54
					v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v66 == int32(0) {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v70 = F_pg_detoast_datum(m, v69)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							F_do_numeric_accum(m, v64, v70)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(16)
								return v64
							}
						}
					} else {
						m.G0 = v7 + int32(16)
						return v64
					}
				}
			}
		}
	} else {
		v15 = v7 + int32(12)
		v16 = int32(0)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v17 == v16 {
			v34 = int32(0)
			if v15 == v34 {
				v42 = v34
			} else {
				v37 = v34
				v38 = v16
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
				v42 = v38
			}
			v45 = v42
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			switch v20 - int32(429) {
			case 0:
				if v15 == int32(0) {
					v45 = int32(1)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
					v37 = v27
					v38 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
					v42 = v38
					v45 = v42
				}
			case 1:
				if v15 == int32(0) {
					v45 = int32(2)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+368))
					v37 = v32
					v38 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
					v42 = v38
					v45 = v42
				}
			default:
				v34 = int32(0)
				if v15 == v34 {
					v42 = v34
				} else {
					v37 = v34
					v38 = v16
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
					v42 = v38
				}
				v45 = v42
			}
		}
		if v45 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(66285), int32(0))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(520410), int32(4943), int32(368516))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v48 = int32(4548768)
			v49 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v51
			v54 = F_palloc0(m, int32(112))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v58 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v58)
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v60
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v49
				v64 = v54
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v66 == int32(0) {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v70 = F_pg_detoast_datum(m, v69)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						F_do_numeric_accum(m, v64, v70)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(16)
							return v64
						}
					}
				} else {
					m.G0 = v7 + int32(16)
					return v64
				}
			}
		}
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int64
	_ = v137
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int64
	_ = v151
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
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
			v181 = v46
			m.G0 = v8 + int32(32)
			return v181
		} else {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v48 == int32(0) {
				v181 = v46
				m.G0 = v8 + int32(32)
				return v181
			} else {
				if v46 == int32(0) {
					v53 = int32(4548768)
					v54 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v56
					v59 = F_palloc0(m, int32(112))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v63)
						v66 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v59)+36))
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+36))
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								v96 = v94 << (uint(int32(2)) % 32)
								if v96 != 0 {
									v97 = F__emscripten_memcpy_bulkmem(m, v92, v93, v96)
									mBase = m.M
								} else {
								}
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v59)+40))
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v48)+40))
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								v103 = v101 << (uint(int32(2)) % 32)
								if v103 != 0 {
									v104 = F__emscripten_memcpy_bulkmem(m, v99, v100, v103)
									mBase = m.M
								} else {
								}
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v48)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+28)) = v106
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v108
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v110
								v112 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = v112
								v175 = v59
								v176 = v54
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v176
								v181 = v175
								m.G0 = v8 + int32(32)
								return v181
							}
						}
					}
				} else {
					v114 = *(*int64)(unsafe.Add(mBase, uint32(v46)+8))
					v115 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+8)) = v114 + v115
					v118 = *(*int64)(unsafe.Add(mBase, uint32(v46)+88))
					v119 = *(*int64)(unsafe.Add(mBase, uint32(v48)+88))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+88)) = v118 + v119
					v122 = *(*int64)(unsafe.Add(mBase, uint32(v46)+96))
					v123 = *(*int64)(unsafe.Add(mBase, uint32(v48)+96))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+96)) = v122 + v123
					v126 = *(*int64)(unsafe.Add(mBase, uint32(v46)+104))
					v127 = *(*int64)(unsafe.Add(mBase, uint32(v48)+104))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+104)) = v126 + v127
					v130 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
					if v130 <= int64(0) {
						v181 = v46
						m.G0 = v8 + int32(32)
						return v181
					} else {
						v133 = *(*int32)(unsafe.Add(mBase, uint32(v48)+72))
						v134 = *(*int32)(unsafe.Add(mBase, uint32(v46)+72))
						if v134 < v133 {
							*(*int32)(unsafe.Add(mBase, uint32(v46)+72)) = v133
							v137 = *(*int64)(unsafe.Add(mBase, uint32(v48)+80))
							*(*int64)(unsafe.Add(mBase, uint32(v46)+80)) = v137
						} else {
							if v133 != v134 {
							} else {
								v140 = *(*int64)(unsafe.Add(mBase, uint32(v46)+80))
								v141 = *(*int64)(unsafe.Add(mBase, uint32(v48)+80))
								*(*int64)(unsafe.Add(mBase, uint32(v46)+80)) = v140 + v141
							}
						}
						v144 = int32(4548768)
						v145 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v147 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v147
						v150 = v8 + int32(24)
						v151 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v150))) = v151
						*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v151
						*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v151
						F_accum_sum_final(m, v48+int32(16), v8+int32(8))
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return int32(0)
						} else {
							F_accum_sum_add(m, v46+int32(16), v8+int32(8))
							mBase = m.M
							v168 = m.ExcPending
							if v168 != 0 {
								return int32(0)
							} else {
								v169 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
								if v169 == int32(0) {
									v175 = v46
									v176 = v145
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v176
									v181 = v175
									m.G0 = v8 + int32(32)
									return v181
								} else {
									F_pfree(m, v169)
									mBase = m.M
									v173 = m.ExcPending
									if v173 != 0 {
										return int32(0)
									} else {
										v175 = v46
										v176 = v145
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v176
										v181 = v175
										m.G0 = v8 + int32(32)
										return v181
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
		v191 = m.ExcPending
		if v191 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(66285), int32(0))
			mBase = m.M
			v195 = m.ExcPending
			if v195 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(520410), int32(5259), int32(388902))
				mBase = m.M
				v200 = m.ExcPending
				if v200 != 0 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v246 int64
	_ = v246
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v21)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v24 = base.I32_extend16_s(v23)
	if base.Ui32(v23) <= base.Ui32(int32(49151)) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	m.G0 = v19 + int32(80)
	return v450
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v194
	v196 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v196
	v205 = base.B2i32(v24 < v196)
	if v24 < v196 {
		goto L76
	} else {
		goto L77
	}
L6:
	;
	v193 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	v194 = v193
	goto L5
L7:
	;
	if v23 != int32(61440) {
		goto L23
	} else {
		goto L24
	}
L8:
	;
	v65 = F_make_result_opt_error(m, int32(1766380), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L20
	} else {
		goto L21
	}
L9:
	;
	if v55&int32(65535) != int32(49152) {
		goto L7
	} else {
		goto L19
	}
L10:
	;
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v28 = base.I32_extend16_s(v27)
	if base.Ui32(int32(49151)) < base.Ui32(v27) {
		v55 = v28
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v24 == int32(-16384) {
		goto L8
	} else {
		goto L18
	}
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v37 = base.B2i32(int32(0) <= v24)
	if int32(0) <= v24 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v38 = int32(-8)
	goto L16
L15:
	;
	v38 = int32(-6)
	goto L16
L16:
	;
	v39 = int32(base.Ui32(v31)>>(uint(int32(2))%32)) + v38
	v41 = int32(base.Ui32(v39) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v41
	if int32(0) <= v24 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v194 = v23<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v23&int32(63)
	goto L5
L18:
	;
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v55 = v54
	goto L9
L19:
	;
	goto L8
L20:
	;
	return int32(0)
L21:
	;
	v450 = v65
	goto L4
L22:
	;
	v191 = F_make_result_opt_error(m, int32(1766452), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L20
	} else {
		goto L75
	}
L23:
	;
	if v23 != int32(53248) {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if base.Ui32(int32(49152)) <= base.Ui32(v55&int32(65535)) {
		goto L51
	} else {
		goto L52
	}
L26:
	;
	if base.Ui32(int32(49152)) <= base.Ui32(v55&int32(65535)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v79 = F_make_result_opt_error(m, int32(1766380), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L20
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) <= base.I32_extend16_s(v55) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v450 = v79
	goto L4
L31:
	;
	v129 = F_make_result_opt_error(m, int32(1766428), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L20
	} else {
		goto L50
	}
L32:
	;
	v89 = int32(-8)
	goto L34
L33:
	;
	v89 = int32(-6)
	goto L34
L34:
	;
	if base.Ui32(int32(2)) <= base.Ui32(int32(base.Ui32(v81)>>(uint(int32(2))%32))+v89) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v98 = v55 & int32(49152)
	if v98 == int32(32768) {
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
	v101 = v55 << (uint(int32(1)) % 32) & int32(16384)
	goto L40
L39:
	;
	v101 = v98
	goto L40
L40:
	;
	if v101 == int32(16384) {
		goto L31
	} else {
		goto L41
	}
L41:
	;
	v106 = F_make_result_opt_error(m, int32(1766404), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L20
	} else {
		goto L42
	}
L42:
	;
	v450 = v106
	goto L4
L43:
	;
	v108 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v108)
	v450 = int32(0)
	goto L4
L44:
	;
	goto L45
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	F_errcode(m, int32(33816706))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L20
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(249576), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L20
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(520410), int32(3295), int32(220898))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
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
	v450 = v129
	goto L4
L51:
	;
	v137 = F_make_result_opt_error(m, int32(1766380), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L20
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) <= base.I32_extend16_s(v55) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v450 = v137
	goto L4
L55:
	;
	v187 = F_make_result_opt_error(m, int32(1766404), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L20
	} else {
		goto L74
	}
L56:
	;
	v147 = int32(-8)
	goto L58
L57:
	;
	v147 = int32(-6)
	goto L58
L58:
	;
	if base.Ui32(int32(2)) <= base.Ui32(int32(base.Ui32(v139)>>(uint(int32(2))%32))+v147) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v156 = v55 & int32(49152)
	if v156 == int32(32768) {
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
	v159 = v55 << (uint(int32(1)) % 32) & int32(16384)
	goto L64
L63:
	;
	v159 = v156
	goto L64
L64:
	;
	if v159 == int32(16384) {
		goto L55
	} else {
		goto L65
	}
L65:
	;
	v164 = F_make_result_opt_error(m, int32(1766428), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L20
	} else {
		goto L66
	}
L66:
	;
	v450 = v164
	goto L4
L67:
	;
	v166 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v166)
	v450 = int32(0)
	goto L4
L68:
	;
	goto L69
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L20
	} else {
		goto L70
	}
L70:
	;
	F_errcode(m, int32(33816706))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L20
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(249576), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(520410), int32(3318), int32(220898))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
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
	v450 = v187
	goto L4
L75:
	;
	v450 = v191
	goto L4
L76:
	;
	v206 = int32(base.Ui32(v23)>>(uint(int32(7))%32)) & int32(63)
	goto L78
L77:
	;
	v206 = v23 & int32(16383)
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v206
	v213 = v23 & int32(49152)
	if v213 == int32(32768) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v216 = v23 << (uint(int32(1)) % 32) & int32(16384)
	goto L81
L80:
	;
	v216 = v213
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v216
	if v24 < v196 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v220 = int32(6)
	goto L84
L83:
	;
	v220 = int32(8)
	goto L84
L84:
	;
	v221 = l0 + v220
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v221
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v229 = base.B2i32(int32(0) <= v28)
	if int32(0) <= v28 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v230 = int32(-8)
	goto L87
L86:
	;
	v230 = int32(-6)
	goto L87
L87:
	;
	v231 = int32(base.Ui32(v223)>>(uint(int32(2))%32)) + v230
	v233 = int32(base.Ui32(v231) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v233
	if int32(0) <= v28 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v235 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	v245 = v235
	goto L90
L89:
	;
	v245 = v27<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v27&int32(63)
	goto L90
L90:
	;
	v246 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v246
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v245
	v251 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v251
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v246
	v259 = base.B2i32(v28 < v251)
	if v28 < v251 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v260 = int32(6)
	goto L93
L92:
	;
	v260 = int32(8)
	goto L93
L93:
	;
	v261 = l1 + v260
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v261
	if v28 < v251 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v269 = int32(base.Ui32(v27)>>(uint(int32(7))%32)) & int32(63)
	goto L96
L95:
	;
	v269 = v27 & int32(16383)
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v269
	v276 = v27 & int32(49152)
	if v276 == int32(32768) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v279 = v27 << (uint(int32(1)) % 32) & int32(16384)
	goto L99
L98:
	;
	v279 = v276
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v279
	if base.Ui32(int32(2)) <= base.Ui32(v39) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v231) {
		goto L109
	} else {
		goto L110
	}
L101:
	;
	v284 = int32(0)
	goto L104
L102:
	;
	goto L103
L103:
	;
	v324 = int32(0)
	v327 = v324
	v342 = v324
	goto L100
L104:
	;
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221+v284<<(uint(int32(1))%32)))))
	if v304 != 0 {
		v327 = v304
		v342 = v284 - v194
		goto L100
	} else {
		goto L106
	}
L105:
	;
	goto L103
L106:
	;
	v306 = v284 + int32(1)
	if v306 != v41 {
		v284 = v306
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
	v348 = v251
	goto L112
L110:
	;
	v373 = v251
	goto L111
L111:
	;
	v387 = int32(0)
	v390 = v373
	goto L108
L112:
	;
	v364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261+v348<<(uint(int32(1))%32)))))
	if v364 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v373 = int32(0)
	goto L111
L114:
	;
	v387 = v364
	v390 = v245 - v348
	goto L108
L115:
	;
	goto L116
L116:
	;
	v367 = v348 + int32(1)
	if v367 != v233 {
		v348 = v367
		goto L112
	} else {
		goto L117
	}
L117:
	;
	goto L113
L118:
	;
	v426 = (v390+v342+base.B2i32(base.I32_extend16_s(v327) <= base.I32_extend16_s(v387)))<<(uint(int32(2))%32) + int32(16)
	if v206 < v426 {
		goto L124
	} else {
		goto L125
	}
L119:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v231) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v407 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261))))
	if v407 != 0 {
		goto L118
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v408 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v408)
	v450 = int32(0)
	goto L4
L123:
	;
	goto L122
L124:
	;
	v428 = v426
	goto L126
L125:
	;
	v428 = v206
	goto L126
L126:
	;
	if v269 < v428 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v430 = v428
	goto L129
L128:
	;
	v430 = v269
	goto L129
L129:
	;
	if int32(1000) <= v430 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v433 = int32(1000)
	goto L132
L131:
	;
	v433 = v430
	goto L132
L132:
	;
	v434 = int32(1)
	F_div_var(m, v19+int32(56), v19+int32(32), v19+int32(8), v433, v434, v434)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L20
	} else {
		goto L133
	}
L133:
	;
	v440 = F_make_result_opt_error(m, v19+int32(8), l2)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L20
	} else {
		goto L134
	}
L134:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	if v442 == int32(0) {
		v450 = v440
		goto L4
	} else {
		goto L135
	}
L135:
	;
	F_pfree(m, v442)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L20
	} else {
		goto L136
	}
L136:
	;
	v450 = v440
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
		if base.Ui32(int32(49152)) <= base.Ui32(v8) {
			if v8 == int32(61440) {
				v16 = int32(-8388608)
			} else {
				v16 = int32(2143289344)
			}
			if v8 == int32(53248) {
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
				v27 = F_DirectFunctionCall1Coll(m, int32(1470), v22, v25)
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
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
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
			if base.Ui32(v21) <= base.Ui32(int32(49151)) {
				v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
				if base.Ui32(v24) < base.Ui32(int32(49152)) {
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
						v68 = v21 & int32(16383)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v68
					v75 = v21 & int32(49152)
					if v75 == int32(32768) {
						v78 = v21 << (uint(int32(1)) % 32) & int32(16384)
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
					v108 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v108
					v111 = v11 + int32(24)
					*(*int64)(unsafe.Add(mBase, uint32(v111))) = v108
					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v107
					v115 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v115
					*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v108
					v122 = base.B2i32(v32 < v115)
					if v32 < v115 {
						v123 = int32(6)
					} else {
						v123 = int32(8)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v19 + v123
					if v32 < v115 {
						v132 = int32(base.Ui32(v24)>>(uint(int32(7))%32)) & int32(63)
					} else {
						v132 = v24 & int32(16383)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v132
					v139 = v24 & int32(49152)
					if v139 == int32(32768) {
						v142 = v24 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v142 = v139
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v142
					F_gcd_var(m, v11+int32(56), v11+int32(32), v11+int32(8))
					mBase = m.M
					v151 = m.ExcPending
					if v151 != 0 {
						return int32(0)
					} else {
						v155 = F_make_result_opt_error(m, v11+int32(8), int32(0))
						mBase = m.M
						v156 = m.ExcPending
						if v156 != 0 {
							return int32(0)
						} else {
							v157 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
							if v157 == int32(0) {
								v162 = v155
								m.G0 = v11 + int32(80)
								return v162
							} else {
								F_pfree(m, v157)
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return int32(0)
								} else {
									v162 = v155
									m.G0 = v11 + int32(80)
									return v162
								}
							}
						}
					}
				} else {
					v30 = F_make_result_opt_error(m, int32(1766380), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v162 = v30
						m.G0 = v11 + int32(80)
						return v162
					}
				}
			} else {
				v30 = F_make_result_opt_error(m, int32(1766380), int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v162 = v30
					m.G0 = v11 + int32(80)
					return v162
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
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
		if base.Ui32(int32(49152)) <= base.Ui32(v19) {
			v22 = F_palloc(m, v18)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v26 = int32(base.Ui32(v24) >> (uint(int32(2)) % 32))
				if v26 != 0 {
					v27 = F__emscripten_memcpy_bulkmem(m, v22, v12, v26)
					mBase = m.M
				} else {
				}
				v95 = v22
				m.G0 = v9 + int32(32)
				return v95
			}
		} else {
			v31 = base.I32_extend16_s(v19)
			v33 = base.B2i32(int32(0) <= v31)
			if int32(0) <= v31 {
				v34 = int32(-8)
			} else {
				v34 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(base.Ui32(v18+v34) >> (uint(int32(1)) % 32))
			if int32(0) <= v31 {
				v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
				v49 = v39
			} else {
				v49 = v19<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v19&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v49
			v51 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v51
			v56 = base.B2i32(v31 < v51)
			if v31 < v51 {
				v57 = int32(6)
			} else {
				v57 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v12 + v57
			if v31 < v51 {
				v66 = int32(base.Ui32(v19)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v66 = v19 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v66
			v73 = v19 & int32(49152)
			if v73 == int32(32768) {
				v76 = v19 << (uint(int32(1)) % 32) & int32(16384)
			} else {
				v76 = v73
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v76
			v79 = v9 + int32(8)
			F_add_var(m, v79, int32(1766504), v79)
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int32(0)
			} else {
				v88 = F_make_result_opt_error(m, v9+int32(8), int32(0))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					if v90 == int32(0) {
						v95 = v88
						m.G0 = v9 + int32(32)
						return v95
					} else {
						F_pfree(m, v90)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							v95 = v88
							m.G0 = v9 + int32(32)
							return v95
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
	return base.B2i32(v2 == int32(49152))
}
func F_numeric_lcm(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int64
	_ = v162
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v21 = F_pg_detoast_datum(m, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
			if base.Ui32(v23) <= base.Ui32(int32(49151)) {
				v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
				if base.Ui32(v26) < base.Ui32(int32(49152)) {
					v34 = base.I32_extend16_s(v26)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v40 = base.I32_extend16_s(v23)
					v42 = base.B2i32(int32(0) <= v40)
					if int32(0) <= v40 {
						v43 = int32(-8)
					} else {
						v43 = int32(-6)
					}
					v44 = int32(base.Ui32(v35)>>(uint(int32(2))%32)) + v43
					*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = int32(base.Ui32(v44) >> (uint(int32(1)) % 32))
					if int32(0) <= v40 {
						v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+6)))
						v58 = v48
					} else {
						v58 = v23<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v23&int32(63)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v58
					v60 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v60
					v69 = base.B2i32(v40 < v60)
					if v40 < v60 {
						v70 = int32(base.Ui32(v23)>>(uint(int32(7))%32)) & int32(63)
					} else {
						v70 = v23 & int32(16383)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v70
					v77 = v23 & int32(49152)
					if v77 == int32(32768) {
						v80 = v23 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v80 = v77
					}
					*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v80
					if v40 < v60 {
						v84 = int32(6)
					} else {
						v84 = int32(8)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v16 + v84
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
					v93 = base.B2i32(int32(0) <= v34)
					if int32(0) <= v34 {
						v94 = int32(-8)
					} else {
						v94 = int32(-6)
					}
					v95 = int32(base.Ui32(v87)>>(uint(int32(2))%32)) + v94
					*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(base.Ui32(v95) >> (uint(int32(1)) % 32))
					if int32(0) <= v34 {
						v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+6)))
						v109 = v99
					} else {
						v109 = v26<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v26&int32(63)
					}
					v111 = v13 + int32(16)
					v112 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v111))) = v112
					*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v112
					*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v109
					*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v112
					v119 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v119
					v128 = base.B2i32(v34 < v119)
					if v34 < v119 {
						v129 = int32(base.Ui32(v26)>>(uint(int32(7))%32)) & int32(63)
					} else {
						v129 = v26 & int32(16383)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v129
					v136 = v26 & int32(49152)
					if v136 == int32(32768) {
						v139 = v26 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v139 = v136
					}
					*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v139
					if v34 < v119 {
						v143 = int32(6)
					} else {
						v143 = int32(8)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v21 + v143
					if base.B2i32(base.Ui32(int32(1)) < base.Ui32(v95))&base.B2i32(base.Ui32(int32(2)) <= base.Ui32(v44)) == int32(0) {
						v154 = F_palloc(m, int32(2))
						mBase = m.M
						v155 = m.ExcPending
						if v155 != 0 {
							return int32(0)
						} else {
							v156 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(v154))) = uint16(v156)
							v159 = *(*int32)(unsafe.Add(mBase, _consts[1126]))
							*(*int32)(unsafe.Add(mBase, uint32(v111))) = v159
							v162 = *(*int64)(unsafe.Add(mBase, _consts[1127]))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v162
							*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v154 + int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v154
							v198 = v154
							if base.Ui32(v129) < base.Ui32(v70) {
								v200 = v70
							} else {
								v200 = v129
							}
							*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v200
							v205 = F_make_result_opt_error(m, v13+int32(8), int32(0))
							mBase = m.M
							v206 = m.ExcPending
							if v206 != 0 {
								return int32(0)
							} else {
								if v198 == int32(0) {
									v211 = v205
									m.G0 = v13 + int32(80)
									return v211
								} else {
									F_pfree(m, v198)
									mBase = m.M
									v210 = m.ExcPending
									if v210 != 0 {
										return int32(0)
									} else {
										v211 = v205
										m.G0 = v13 + int32(80)
										return v211
									}
								}
							}
						}
					} else {
						F_gcd_var(m, v13+int32(56), v13+int32(32), v13+int32(8))
						mBase = m.M
						v175 = m.ExcPending
						if v175 != 0 {
							return int32(0)
						} else {
							v179 = v13 + int32(8)
							v182 = int32(0)
							F_div_var(m, v13+int32(56), v179, v179, v182, v182, int32(1))
							mBase = m.M
							v186 = m.ExcPending
							if v186 != 0 {
								return int32(0)
							} else {
								v190 = v13 + int32(8)
								F_mul_var(m, v13+int32(32), v190, v190, v129)
								mBase = m.M
								v194 = m.ExcPending
								if v194 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(0)
									v197 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
									v198 = v197
									if base.Ui32(v129) < base.Ui32(v70) {
										v200 = v70
									} else {
										v200 = v129
									}
									*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v200
									v205 = F_make_result_opt_error(m, v13+int32(8), int32(0))
									mBase = m.M
									v206 = m.ExcPending
									if v206 != 0 {
										return int32(0)
									} else {
										if v198 == int32(0) {
											v211 = v205
											m.G0 = v13 + int32(80)
											return v211
										} else {
											F_pfree(m, v198)
											mBase = m.M
											v210 = m.ExcPending
											if v210 != 0 {
												return int32(0)
											} else {
												v211 = v205
												m.G0 = v13 + int32(80)
												return v211
											}
										}
									}
								}
							}
						}
					}
				} else {
					v32 = F_make_result_opt_error(m, int32(1766380), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v211 = v32
						m.G0 = v13 + int32(80)
						return v211
					}
				}
			} else {
				v32 = F_make_result_opt_error(m, int32(1766380), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v211 = v32
					m.G0 = v13 + int32(80)
					return v211
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
		if base.Ui32(int32(49152)) <= base.Ui32(v15) {
			if v15 != int32(61440) {
				if v15 != int32(53248) {
					v29 = F_pstrdup(m, int32(548689))
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
					v23 = F_pstrdup(m, int32(11499))
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
				v26 = F_pstrdup(m, int32(11488))
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
				v66 = v15 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v66
			v73 = v15 & int32(49152)
			if v73 == int32(32768) {
				v76 = v15 << (uint(int32(1)) % 32) & int32(16384)
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
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
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 float64
	_ = v89
	var v100 int64
	_ = v100
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v127 int64
	_ = v127
	var v131 int64
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 float64
	_ = v140
	var v142 float64
	_ = v142
	var v155 float64
	_ = v155
	var v158 float64
	_ = v158
	var v163 float64
	_ = v163
	var v164 float64
	_ = v164
	var v165 float64
	_ = v165
	var v166 float64
	_ = v166
	var v171 float64
	_ = v171
	var v172 float64
	_ = v172
	var v173 float64
	_ = v173
	var v198 float64
	_ = v198
	var v210 float64
	_ = v210
	var v232 float64
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int64
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if base.Ui32(int32(49152)) <= base.Ui32(v13) {
		if v13 != int32(61440) {
			if v13 != int32(53248) {
				v29 = F_pstrdup(m, int32(548689))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v312 = v29
					m.G0 = v11 - int32(-64)
					return v312
				}
			} else {
				v21 = F_pstrdup(m, int32(11499))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v312 = v21
					m.G0 = v11 - int32(-64)
					return v312
				}
			}
		} else {
			v26 = F_pstrdup(m, int32(11488))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v312 = v26
				m.G0 = v11 - int32(-64)
				return v312
			}
		}
	} else {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v36 = base.I32_extend16_s(v13)
		v38 = base.B2i32(int32(0) <= v36)
		if int32(0) <= v36 {
			v39 = int32(-8)
		} else {
			v39 = int32(-6)
		}
		v40 = int32(base.Ui32(v31)>>(uint(int32(2))%32)) + v39
		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(base.Ui32(v40) >> (uint(int32(1)) % 32))
		if int32(0) <= v36 {
			v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
			v54 = v44
		} else {
			v54 = v13<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v13&int32(63)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v54
		v56 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v56
		v61 = base.B2i32(v36 < v56)
		if v36 < v56 {
			v62 = int32(6)
		} else {
			v62 = int32(8)
		}
		v63 = l0 + v62
		*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v63
		if v36 < v56 {
			v71 = int32(base.Ui32(v13)>>(uint(int32(7))%32)) & int32(63)
		} else {
			v71 = v13 & int32(16383)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v71
		v78 = v13 & int32(49152)
		if v78 == int32(32768) {
			v81 = v13 << (uint(int32(1)) % 32) & int32(16384)
		} else {
			v81 = v78
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v81
		if base.Ui32(int32(2)) <= base.Ui32(v40) {
			v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(v63))))
			v89 = base.F64_convert_i32_s(v88)
			v100 = base.I64_reinterpret_f64(v89)
			if v100 <= int64(4503599627370495) {
				if base.F64_eq(v89, float64(0)) != 0 {
					v232 = base.F64_div(float64(-1), base.F64_mul(v89, v89))
				} else {
					if int64(0) <= v100 {
						v127 = base.I64_reinterpret_f64(base.F64_mul(v89, float64(1.8014398509481984e+16)))
						v131 = v127
						v133 = int32(-1077)
						v134 = base.I32_wrap_i64(int64(base.Ui64(v127) >> (uint(int64(32)) % 64)))
						v136 = v134 + int32(614242)
						v140 = base.F64_convert_i32_s(int32(base.Ui32(v136)>>(uint(int32(20))%32)) + v133)
						v142 = base.F64_mul(v140, float64(0.30102999566361177))
						v155 = base.F64_add(base.F64_reinterpret_i64(v131&int64(4294967295)|base.I64_extend_i32_u(v136&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
						v158 = base.F64_mul(v155, base.F64_mul(v155, float64(0.5)))
						v163 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v155, v158)) & int64(-4294967296))
						v164 = float64(0.4342944818781689)
						v165 = base.F64_mul(v163, v164)
						v166 = base.F64_add(v142, v165)
						v171 = base.F64_div(v155, base.F64_add(v155, float64(2)))
						v172 = base.F64_mul(v171, v171)
						v173 = base.F64_mul(v172, v172)
						v198 = base.F64_add(base.F64_mul(v171, base.F64_add(v158, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v172, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v155, v163), v158))
						v210 = base.F64_add(v166, base.F64_add(base.F64_add(v165, base.F64_sub(v142, v166)), base.F64_add(base.F64_mul(v198, v164), base.F64_add(base.F64_mul(v140, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v198, v163), float64(2.5082946711645275e-11))))))
						v232 = v210
					} else {
						v232 = base.F64_div(base.F64_sub(v89, v89), float64(0))
					}
				}
			} else {
				if base.Ui64(int64(9218868437227405311)) < base.Ui64(v100) {
					v210 = v89
					v232 = v210
				} else {
					v115 = int32(-1023)
					v117 = int64(base.Ui64(v100) >> (uint(int64(32)) % 64))
					if v117 != int64(1072693248) {
						v131 = v100
						v133 = v115
						v134 = base.I32_wrap_i64(v117)
						v136 = v134 + int32(614242)
						v140 = base.F64_convert_i32_s(int32(base.Ui32(v136)>>(uint(int32(20))%32)) + v133)
						v142 = base.F64_mul(v140, float64(0.30102999566361177))
						v155 = base.F64_add(base.F64_reinterpret_i64(v131&int64(4294967295)|base.I64_extend_i32_u(v136&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
						v158 = base.F64_mul(v155, base.F64_mul(v155, float64(0.5)))
						v163 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v155, v158)) & int64(-4294967296))
						v164 = float64(0.4342944818781689)
						v165 = base.F64_mul(v163, v164)
						v166 = base.F64_add(v142, v165)
						v171 = base.F64_div(v155, base.F64_add(v155, float64(2)))
						v172 = base.F64_mul(v171, v171)
						v173 = base.F64_mul(v172, v172)
						v198 = base.F64_add(base.F64_mul(v171, base.F64_add(v158, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v172, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v155, v163), v158))
						v210 = base.F64_add(v166, base.F64_add(base.F64_add(v165, base.F64_sub(v142, v166)), base.F64_add(base.F64_mul(v198, v164), base.F64_add(base.F64_mul(v140, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v198, v163), float64(2.5082946711645275e-11))))))
						v232 = v210
					} else {
						if base.I32_wrap_i64(v100) != 0 {
							v131 = v100
							v133 = v115
							v134 = int32(1072693248)
							v136 = v134 + int32(614242)
							v140 = base.F64_convert_i32_s(int32(base.Ui32(v136)>>(uint(int32(20))%32)) + v133)
							v142 = base.F64_mul(v140, float64(0.30102999566361177))
							v155 = base.F64_add(base.F64_reinterpret_i64(v131&int64(4294967295)|base.I64_extend_i32_u(v136&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
							v158 = base.F64_mul(v155, base.F64_mul(v155, float64(0.5)))
							v163 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v155, v158)) & int64(-4294967296))
							v164 = float64(0.4342944818781689)
							v165 = base.F64_mul(v163, v164)
							v166 = base.F64_add(v142, v165)
							v171 = base.F64_div(v155, base.F64_add(v155, float64(2)))
							v172 = base.F64_mul(v171, v171)
							v173 = base.F64_mul(v172, v172)
							v198 = base.F64_add(base.F64_mul(v171, base.F64_add(v158, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v172, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v155, v163), v158))
							v210 = base.F64_add(v166, base.F64_add(base.F64_add(v165, base.F64_sub(v142, v166)), base.F64_add(base.F64_mul(v198, v164), base.F64_add(base.F64_mul(v140, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v198, v163), float64(2.5082946711645275e-11))))))
							v232 = v210
						} else {
							v232 = float64(0)
						}
					}
				}
			}
			if base.F64_lt(base.F64_abs(v232), float64(2.147483648e+09)) != 0 {
				v236 = base.I32_trunc_f64_s(v232)
				v238 = v236
			} else {
				v238 = int32(-2147483648)
			}
			v240 = v238 + v54<<(uint(int32(2))%32)
		} else {
			v240 = int32(0)
		}
		v243 = F_palloc(m, int32(4))
		mBase = m.M
		v244 = m.ExcPending
		if v244 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v243))) = int32(65536)
			v248 = *(*int32)(unsafe.Add(mBase, _consts[1124]))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v248
			v251 = *(*int64)(unsafe.Add(mBase, _consts[1125]))
			*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v251
			v253 = int32(2)
			v254 = v243 + v253
			*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v254
			*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v243
			v257 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = (v257 - v240) & (v240 >> (uint(int32(31)) % 32))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v240 >> (uint(v253) % 32)
			if v257 < l1 {
				v269 = l1
			} else {
				v269 = v257
			}
			v272 = v240 & int32(3)
			switch v272 {
			case 0:
			case 1:
				v280 = int32(10)
				*(*uint16)(unsafe.Add(mBase, uint32(v254))) = uint16(v280)
			default:
				if base.Ui32(v272-int32(3)) < base.Ui32(int32(-2)) {
					v279 = int32(1000)
				} else {
					v279 = int32(100)
				}
				v280 = v279
				*(*uint16)(unsafe.Add(mBase, uint32(v254))) = uint16(v280)
			}
			v286 = v9 + int32(-24)
			v289 = int32(1)
			F_div_var(m, v9+int32(-48), v286, v286, v269, v289, v289)
			mBase = m.M
			v292 = m.ExcPending
			if v292 != 0 {
				return int32(0)
			} else {
				v295 = F_get_str_from_var(m, v9+int32(-24))
				mBase = m.M
				v296 = m.ExcPending
				if v296 != 0 {
					return int32(0)
				} else {
					v297 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
					if v297 != 0 {
						F_pfree(m, v297)
						mBase = m.M
						v299 = m.ExcPending
						if v299 != 0 {
							return int32(0)
						} else {
							v300 = F_strlen(m, v295)
							mBase = m.M
							v302 = v300 + int32(13)
							v303 = F_palloc(m, v302)
							mBase = m.M
							v304 = m.ExcPending
							if v304 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v240
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v295
								v308 = F_pg_snprintf(m, v303, v302, int32(483293), v11)
								mBase = m.M
								v309 = m.ExcPending
								if v309 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v295)
									mBase = m.M
									v311 = m.ExcPending
									if v311 != 0 {
										return int32(0)
									} else {
										v312 = v303
										m.G0 = v11 - int32(-64)
										return v312
									}
								}
							}
						}
					} else {
						v300 = F_strlen(m, v295)
						mBase = m.M
						v302 = v300 + int32(13)
						v303 = F_palloc(m, v302)
						mBase = m.M
						v304 = m.ExcPending
						if v304 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v240
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v295
							v308 = F_pg_snprintf(m, v303, v302, int32(483293), v11)
							mBase = m.M
							v309 = m.ExcPending
							if v309 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v295)
								mBase = m.M
								v311 = m.ExcPending
								if v311 != 0 {
									return int32(0)
								} else {
									v312 = v303
									m.G0 = v11 - int32(-64)
									return v312
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v49 int64
	_ = v49
	var v57 int64
	_ = v57
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v122 int64
	_ = v122
	var v129 int64
	_ = v129
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v147 int64
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v18 != 0 {
		v26 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
		v177 = int32(0)
		m.G0 = v16 - int32(-64)
		return v177
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v19 == int32(0) {
			v26 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
			v177 = int32(0)
			m.G0 = v16 - int32(-64)
			return v177
		} else {
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
			if v22 != int64(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = int64(0)
				v31 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
				v32 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
				v34 = F_palloc(m, int32(22))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v34
					v39 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v34))) = uint16(v39)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v34 + int32(2)
					if v31 < int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = int64(16384)
						v49 = int64(0)
						v62 = v49 - (v31 + base.I64_extend_i32_u(base.B2i32(v32 != v49)))
						v63 = v49 - v32
						v68 = v39
						v70 = v34 + int32(22)
						v77 = v62
						v78 = v63
						for {
							v82 = v14 + int32(-40)
							v85 = m.G0
							v86 = int32(16)
							v87 = v85 - v86
							m.G0 = v87
							F___udivmodti4(m, v87, v78, v77, int64(10000), int64(0))
							mBase = m.M
							v91 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
							v92 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v82)+8)) = v92
							*(*int64)(unsafe.Add(mBase, uint32(v82))) = v91
							m.G0 = v87 + v86
							v99 = v14 + int32(-56)
							v100 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
							v101 = *(*int64)(unsafe.Add(mBase, uint32(v14+int32(-32))))
							v102 = int64(55536)
							v103 = int64(0)
							v108 = int64(32)
							v111 = int64(base.Ui64(v100) >> (uint(v108) % 64))
							v114 = int64(4294967295)
							v117 = v100 & v114
							v118 = v102 * v117
							v122 = int64(base.Ui64(v118)>>(uint(v108)%64)) + v102*v111
							v129 = v117*v103 + v122&v114
							*(*int64)(unsafe.Add(mBase, uint32(v99)+8)) = v100*v103 + v101*v102 + v103*v111 + int64(base.Ui64(v122)>>(uint(v108)%64)) + int64(base.Ui64(v129)>>(uint(v108)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v99))) = v118&v114 | v129<<(uint(v108)%64)
							v141 = v70 - int32(2)
							v142 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
							v143 = v142 + v78
							*(*uint16)(unsafe.Add(mBase, uint32(v141))) = uint16(v143)
							v147 = int64(0)
							v152 = v68 + int32(1)
							if v77 == v147 {
								v153 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v78))
							} else {
								v153 = base.B2i32(v77 != v147)
							}
							if v153 != 0 {
								v68 = v152
								v70 = v141
								v77 = v101
								v78 = v100
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v141
						v155 = v152
						v159 = v68
					} else {
						v57 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v57
						if v31|v32 == v57 {
							v155 = v39
							v159 = int32(0)
						} else {
							v62 = v31
							v63 = v32
							v68 = v39
							v70 = v34 + int32(22)
							v77 = v62
							v78 = v63
							for {
								v82 = v14 + int32(-40)
								v85 = m.G0
								v86 = int32(16)
								v87 = v85 - v86
								m.G0 = v87
								F___udivmodti4(m, v87, v78, v77, int64(10000), int64(0))
								mBase = m.M
								v91 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
								v92 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v82)+8)) = v92
								*(*int64)(unsafe.Add(mBase, uint32(v82))) = v91
								m.G0 = v87 + v86
								v99 = v14 + int32(-56)
								v100 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
								v101 = *(*int64)(unsafe.Add(mBase, uint32(v14+int32(-32))))
								v102 = int64(55536)
								v103 = int64(0)
								v108 = int64(32)
								v111 = int64(base.Ui64(v100) >> (uint(v108) % 64))
								v114 = int64(4294967295)
								v117 = v100 & v114
								v118 = v102 * v117
								v122 = int64(base.Ui64(v118)>>(uint(v108)%64)) + v102*v111
								v129 = v117*v103 + v122&v114
								*(*int64)(unsafe.Add(mBase, uint32(v99)+8)) = v100*v103 + v101*v102 + v103*v111 + int64(base.Ui64(v122)>>(uint(v108)%64)) + int64(base.Ui64(v129)>>(uint(v108)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v99))) = v118&v114 | v129<<(uint(v108)%64)
								v141 = v70 - int32(2)
								v142 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
								v143 = v142 + v78
								*(*uint16)(unsafe.Add(mBase, uint32(v141))) = uint16(v143)
								v147 = int64(0)
								v152 = v68 + int32(1)
								if v77 == v147 {
									v153 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v78))
								} else {
									v153 = base.B2i32(v77 != v147)
								}
								if v153 != 0 {
									v68 = v152
									v70 = v141
									v77 = v101
									v78 = v100
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v141
							v155 = v152
							v159 = v68
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v159
					*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v155
					v173 = F_make_result_opt_error(m, v14+int32(-24), int32(0))
					mBase = m.M
					v174 = m.ExcPending
					if v174 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v34)
						mBase = m.M
						v176 = m.ExcPending
						if v176 != 0 {
							return int32(0)
						} else {
							v177 = v173
							m.G0 = v16 - int32(-64)
							return v177
						}
					}
				}
			} else {
				v26 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
				v177 = int32(0)
				m.G0 = v16 - int32(-64)
				return v177
			}
		}
	}
}
func F_numeric_poly_var_samp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v8 != 0 {
		v10 = int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v10 = v9
	}
	v11 = int32(1)
	v15 = F_numeric_poly_stddev_internal(m, v10, v11, v11, v6+int32(15))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
		if v19 == int32(1) {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
			v25 = int32(0)
		} else {
			v25 = v15
		}
		m.G0 = v6 + int32(16)
		return v25
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
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
	v25 = v20 & int32(65535)
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
	v48 = v45 & int32(65535)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v48
	if v48 == int32(61440) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L55
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L51
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L47
	}
L9:
	;
	v61 = F_pq_getmsgint(m, v18, int32(2))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L13
	}
L10:
	;
	if v45&int32(49151) == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if v45&int32(61439) != int32(49152) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	v64 = v61 & int32(65535)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v64
	if v61&int32(49152) != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	if v25 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v69 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	if v45&int32(45056) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	v82 = F_pq_getmsgint(m, v18, int32(2))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L17
L20:
	;
	if base.Ui32(int32(10000)) <= base.Ui32(v82&int32(65535)) {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v88 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v37+v69<<(uint(v88)%32)))) = uint16(v82)
	v93 = v69 + v88
	if v93 != v25 {
		v69 = v93
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	m.G0 = v15 + int32(32)
	return v210
L24:
	;
	F_pfree(m, v206)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L46
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v64
	v116 = v40<<(uint(int32(16))%32)>>(uint(int32(14))%32) + v64
	if v116+int32(4) <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	goto L27
L27:
	;
	v165 = F_make_result_opt_error(m, v15+int32(8), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L37
	}
L28:
	;
	v154 = F_apply_typmod(m, v15+int32(8), v17, int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L34
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = int64(0)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v128 = int32(base.Ui32(v116+int32(7)) >> (uint(int32(2)) % 32))
	if base.Ui32(v25) < base.Ui32(v128) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v128
	v132 = v61 & int32(3)
	if v132 == int32(0) {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v138 = int32(2)
	v139 = v37 + v128<<(uint(int32(1))%32) - v138
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(v139))))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v132<<(uint(v138)%32))+uint32(_consts[1123])))
	v146 = base.I32_rem_s(v140, v145)
	v147 = v140 - v146
	*(*uint16)(unsafe.Add(mBase, uint32(v139))) = uint16(v147)
	goto L28
L34:
	;
	v159 = F_make_result_opt_error(m, v15+int32(8), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	if v161 != 0 {
		v205 = v159
		v206 = v161
		goto L24
	} else {
		goto L36
	}
L36:
	;
	v210 = v159
	goto L23
L37:
	;
	if v17 < int32(4) {
		v205 = v165
		v206 = v30
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+4)))
	if v169 == int32(49152) {
		v205 = v165
		v206 = v30
		goto L24
	} else {
		goto L39
	}
L39:
	;
	v173 = F_errsave_start(m, int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v173 == int32(0) {
		v205 = v165
		v206 = v30
		goto L24
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(32555), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(base.Ui32(v17-int32(4)) >> (uint(int32(16)) % 32))
	v189 = int32(21)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = (v17<<(uint(v189)%32) - int32(8388608)) >> (uint(v189) % 32)
	F_errdetail(m, int32(651984), v15)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errsave_finish(m, int32(0), int32(520410), int32(8138), int32(326347))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v205 = v165
	v206 = v30
	goto L24
L46:
	;
	v210 = v205
	goto L23
L47:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(361563), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(520410), int32(1108), int32(37833))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errmsg(m, int32(361604), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(520410), int32(1114), int32(37833))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(361521), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(520410), int32(1123), int32(37833))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
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
		v9 = int32(49152)
		if v8&v9 == v9 {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
		} else {
			if base.I32_extend16_s(v8) < int32(0) {
				v26 = int32(base.Ui32(v8)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v26 = v8 & int32(16383)
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
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v157 int64
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v173 int64
	_ = v173
	var v175 int64
	_ = v175
	var v177 int64
	_ = v177
	var v179 int64
	_ = v179
	var v202 int32
	_ = v202
	var v205 int64
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v221 int64
	_ = v221
	var v223 int64
	_ = v223
	var v225 int64
	_ = v225
	var v227 int64
	_ = v227
	var v250 int32
	_ = v250
	var v253 int64
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int64
	_ = v262
	var v264 int64
	_ = v264
	var v266 int64
	_ = v266
	var v269 int64
	_ = v269
	var v271 int64
	_ = v271
	var v273 int64
	_ = v273
	var v275 int64
	_ = v275
	var v298 int32
	_ = v298
	var v301 int64
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int64
	_ = v310
	var v312 int64
	_ = v312
	var v314 int64
	_ = v314
	var v317 int64
	_ = v317
	var v319 int64
	_ = v319
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
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
		v44 = v9 + int32(24)
		v45 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v44))) = v45
		*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v45
		*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v45
		F_pq_begintypsend(m, v9+int32(32))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int32(0)
		} else {
			v57 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
			F_enlargeStringInfo(m, v9+int32(32), int32(8))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
				v66 = int64(56)
				v68 = int64(65280)
				v70 = int64(40)
				v73 = int64(16711680)
				v75 = int64(24)
				v77 = int64(4278190080)
				v79 = int64(8)
				*(*int64)(unsafe.Add(mBase, uint32(v63+v64))) = v57<<(uint(v66)%64) | v57&v68<<(uint(v70)%64) | (v57&v73<<(uint(v75)%64) | v57&v77<<(uint(v79)%64)) | (int64(base.Ui64(v57)>>(uint(v79)%64))&v77 | int64(base.Ui64(v57)>>(uint(v75)%64))&v73 | (int64(base.Ui64(v57)>>(uint(v70)%64))&v68 | int64(base.Ui64(v57)>>(uint(v66)%64))))
				v102 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v63 + v102
				F_accum_sum_final(m, v42+int32(16), v9+v102)
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					F_numericvar_serialize(m, v9+int32(32), v9+int32(8))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						F_accum_sum_final(m, v42+int32(44), v9+int32(8))
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return int32(0)
						} else {
							F_numericvar_serialize(m, v9+int32(32), v9+int32(8))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return int32(0)
							} else {
								v129 = *(*int32)(unsafe.Add(mBase, uint32(v42)+72))
								F_enlargeStringInfo(m, v9+int32(32), int32(4))
								mBase = m.M
								v134 = m.ExcPending
								if v134 != 0 {
									return int32(0)
								} else {
									v135 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
									v136 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
									v138 = int32(24)
									v140 = int32(65280)
									v142 = int32(8)
									*(*int32)(unsafe.Add(mBase, uint32(v135+v136))) = v129<<(uint(v138)%32) | v129&v140<<(uint(v142)%32) | (int32(base.Ui32(v129)>>(uint(v142)%32))&v140 | int32(base.Ui32(v129)>>(uint(v138)%32)))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v135 + int32(4)
									v157 = *(*int64)(unsafe.Add(mBase, uint32(v42)+80))
									F_enlargeStringInfo(m, v9+int32(32), v142)
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return int32(0)
									} else {
										v163 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
										v164 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
										v166 = int64(56)
										v168 = int64(65280)
										v170 = int64(40)
										v173 = int64(16711680)
										v175 = int64(24)
										v177 = int64(4278190080)
										v179 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v163+v164))) = v157<<(uint(v166)%64) | v157&v168<<(uint(v170)%64) | (v157&v173<<(uint(v175)%64) | v157&v177<<(uint(v179)%64)) | (int64(base.Ui64(v157)>>(uint(v179)%64))&v177 | int64(base.Ui64(v157)>>(uint(v175)%64))&v173 | (int64(base.Ui64(v157)>>(uint(v170)%64))&v168 | int64(base.Ui64(v157)>>(uint(v166)%64))))
										v202 = int32(8)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v163 + v202
										v205 = *(*int64)(unsafe.Add(mBase, uint32(v42)+88))
										F_enlargeStringInfo(m, v9+int32(32), v202)
										mBase = m.M
										v210 = m.ExcPending
										if v210 != 0 {
											return int32(0)
										} else {
											v211 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
											v212 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
											v214 = int64(56)
											v216 = int64(65280)
											v218 = int64(40)
											v221 = int64(16711680)
											v223 = int64(24)
											v225 = int64(4278190080)
											v227 = int64(8)
											*(*int64)(unsafe.Add(mBase, uint32(v211+v212))) = v205<<(uint(v214)%64) | v205&v216<<(uint(v218)%64) | (v205&v221<<(uint(v223)%64) | v205&v225<<(uint(v227)%64)) | (int64(base.Ui64(v205)>>(uint(v227)%64))&v225 | int64(base.Ui64(v205)>>(uint(v223)%64))&v221 | (int64(base.Ui64(v205)>>(uint(v218)%64))&v216 | int64(base.Ui64(v205)>>(uint(v214)%64))))
											v250 = int32(8)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v211 + v250
											v253 = *(*int64)(unsafe.Add(mBase, uint32(v42)+96))
											F_enlargeStringInfo(m, v9+int32(32), v250)
											mBase = m.M
											v258 = m.ExcPending
											if v258 != 0 {
												return int32(0)
											} else {
												v259 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
												v260 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
												v262 = int64(56)
												v264 = int64(65280)
												v266 = int64(40)
												v269 = int64(16711680)
												v271 = int64(24)
												v273 = int64(4278190080)
												v275 = int64(8)
												*(*int64)(unsafe.Add(mBase, uint32(v259+v260))) = v253<<(uint(v262)%64) | v253&v264<<(uint(v266)%64) | (v253&v269<<(uint(v271)%64) | v253&v273<<(uint(v275)%64)) | (int64(base.Ui64(v253)>>(uint(v275)%64))&v273 | int64(base.Ui64(v253)>>(uint(v271)%64))&v269 | (int64(base.Ui64(v253)>>(uint(v266)%64))&v264 | int64(base.Ui64(v253)>>(uint(v262)%64))))
												v298 = int32(8)
												*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v259 + v298
												v301 = *(*int64)(unsafe.Add(mBase, uint32(v42)+104))
												F_enlargeStringInfo(m, v9+int32(32), v298)
												mBase = m.M
												v306 = m.ExcPending
												if v306 != 0 {
													return int32(0)
												} else {
													v307 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
													v308 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
													v310 = int64(56)
													v312 = int64(65280)
													v314 = int64(40)
													v317 = int64(16711680)
													v319 = int64(24)
													v321 = int64(4278190080)
													v323 = int64(8)
													*(*int64)(unsafe.Add(mBase, uint32(v307+v308))) = v301<<(uint(v310)%64) | v301&v312<<(uint(v314)%64) | (v301&v317<<(uint(v319)%64) | v301&v321<<(uint(v323)%64)) | (int64(base.Ui64(v301)>>(uint(v323)%64))&v321 | int64(base.Ui64(v301)>>(uint(v319)%64))&v317 | (int64(base.Ui64(v301)>>(uint(v314)%64))&v312 | int64(base.Ui64(v301)>>(uint(v310)%64))))
													*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v307 + int32(8)
													v350 = v9 + int32(32)
													v352 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
													v353 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v352))) = v353 << (uint(int32(2)) % 32)
													v357 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
													if v357 != 0 {
														F_pfree(m, v357)
														mBase = m.M
														v359 = m.ExcPending
														if v359 != 0 {
															return int32(0)
														} else {
															m.G0 = v9 + int32(48)
															return v352
														}
													} else {
														m.G0 = v9 + int32(48)
														return v352
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
		v367 = m.ExcPending
		if v367 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(66285), int32(0))
			mBase = m.M
			v371 = m.ExcPending
			if v371 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(520410), int32(5442), int32(356551))
				mBase = m.M
				v376 = m.ExcPending
				if v376 != 0 {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
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
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+4)))
			v25 = int32(49152)
			v26 = v24 & v25
			if v26 == v25 {
				if v24 != int32(53248) {
					if v24 != int32(49152) {
						if v23 != int32(61440) {
							v45 = int32(-1)
						} else {
							v45 = int32(0)
						}
						v164 = v45
					} else {
						v164 = base.B2i32(v23 != int32(49152))
					}
				} else {
					if v23 == int32(49152) {
						v40 = int32(-1)
					} else {
						v40 = base.B2i32(v23 != int32(53248))
					}
					v164 = v40
				}
			} else {
				if base.Ui32(int32(49152)) <= base.Ui32(v23) {
					if v23 == int32(61440) {
						v52 = int32(1)
					} else {
						v52 = int32(-1)
					}
					v164 = v52
				} else {
					v54 = v4 + int32(6)
					v59 = base.B2i32(int32(0) <= base.I32_extend16_s(v24))
					if int32(0) <= base.I32_extend16_s(v24) {
						v60 = int32(-8)
					} else {
						v60 = int32(-6)
					}
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
					if int32(0) <= base.I32_extend16_s(v24) {
						v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54))))
						v74 = v64
					} else {
						v74 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					v75 = v60 + int32(base.Ui32(v61)>>(uint(int32(2))%32))
					v77 = v9 + int32(6)
					v82 = base.B2i32(int32(0) <= base.I32_extend16_s(v23))
					if int32(0) <= base.I32_extend16_s(v23) {
						v83 = int32(-8)
					} else {
						v83 = int32(-6)
					}
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					if int32(0) <= base.I32_extend16_s(v23) {
						v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v77))))
						v97 = v87
					} else {
						v97 = v23<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v23&int32(63)
					}
					v98 = v83 + int32(base.Ui32(v84)>>(uint(int32(2))%32))
					v104 = v23 & int32(49152)
					if v104 == int32(32768) {
						v107 = v23 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v107 = v104
					}
					if base.Ui32(v75) <= base.Ui32(int32(1)) {
						if base.Ui32(v98) < base.Ui32(int32(2)) {
							v164 = int32(0)
						} else {
							if v107 == int32(16384) {
								v117 = int32(1)
							} else {
								v117 = int32(-1)
							}
							v164 = v117
						}
					} else {
						if v26 == int32(32768) {
							v124 = v24 << (uint(int32(1)) % 32) & int32(16384)
						} else {
							v124 = v26
						}
						if base.Ui32(v98) <= base.Ui32(int32(1)) {
							if v124 != 0 {
								v129 = int32(-1)
							} else {
								v129 = int32(1)
							}
							v164 = v129
						} else {
							if int32(0) <= base.I32_extend16_s(v24) {
								v132 = v4 + int32(8)
							} else {
								v132 = v54
							}
							v134 = int32(base.Ui32(v75) >> (uint(int32(1)) % 32))
							if int32(0) <= base.I32_extend16_s(v23) {
								v137 = v9 + int32(8)
							} else {
								v137 = v77
							}
							v139 = int32(base.Ui32(v98) >> (uint(int32(1)) % 32))
							if v124 == int32(0) {
								if v107 == int32(16384) {
									v164 = int32(1)
								} else {
									v145 = F_cmp_abs_common(m, v132, v134, v74, v137, v139, v97)
									mBase = m.M
									v164 = v145
								}
							} else {
								if v107 == int32(0) {
									v164 = int32(-1)
								} else {
									v149 = F_cmp_abs_common(m, v137, v139, v97, v132, v134, v74)
									mBase = m.M
									v164 = v149
								}
							}
						}
					}
				}
			}
			if v164 < int32(0) {
				v167 = v4
			} else {
				v167 = v9
			}
			return v167
		}
	}
}
func F_numeric_stddev_samp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v8 != 0 {
		v10 = int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v10 = v9
	}
	v15 = F_numeric_stddev_internal(m, v10, int32(0), int32(1), v6+int32(15))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
		if v19 == int32(1) {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
			v25 = int32(0)
		} else {
			v25 = v15
		}
		m.G0 = v6 + int32(16)
		return v25
	}
}
