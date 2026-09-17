package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_read_stream_begin_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	v5 = l4
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[0]))
	if v18 == int32(0) {
		v32 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[1]))
		v50 = v32
		v52 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
		if l1 == int32(0) {
			v56 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[3]))
			v65 = v56
		} else {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v59 = base.I32_div_s(v57, int32(2))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v60 != int32(1) {
				v63 = v59
			} else {
				v63 = v57
			}
			v65 = v63
		}
		v66 = int32(_a_F_read_stream_begin_impl_0)
		if v66 <= v50 {
			v69 = v66
		} else {
			v69 = v50
		}
		v72 = v52 * (v69 + int32(1))
		v74 = int32(16)
		v79 = (v52<<(uint(v74)%32) - int32(_a_F_read_stream_begin_impl_1)) >> (uint(v74) % 32)
		v80 = int32(_a_F_read_stream_begin_impl_2) - v79
		if base.Ui32(v72) < base.Ui32(v80) {
			v82 = v72
		} else {
			v82 = v80
		}
		if base.Ui32(v65) < base.Ui32(v82) {
			v84 = v65
		} else {
			v84 = v82
		}
		v85 = int32(1)
		if v69 <= v85 {
			v88 = v85
		} else {
			v88 = v69
		}
		v90 = v88 * int32(84)
		v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
		if v92 != int32(-1) {
			v96 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[4]))
			v99 = v96
		} else {
			v98 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[5]))
			v99 = v98
		}
		if base.Ui32(v84) < base.Ui32(v99) {
			v101 = v84
		} else {
			v101 = v99
		}
		if base.Ui32(v101) <= base.Ui32(int32(1)) {
			v104 = int32(1)
		} else {
			v104 = v101
		}
		v106 = v104 + int32(1)
		v107 = base.I32_extend16_s(v106)
		v112 = (v107 + v79) << (uint(int32(2)) % 32)
		v116 = F_palloc(m, v90+l8*v107+v112+int32(96))
		mBase = m.M
		v117 = m.ExcPending
		if v117 != 0 {
			return int32(0)
		} else {
			base.MemoryFill(m, v116, int32(0), int32(80))
			v125 = (v116 + v112 + int32(87)) & int32(-8)
			*(*int32)(unsafe.Add(mBase, uint32(v116)+64)) = v125
			if l8 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v116)+60)) = (v90 + v125 + int32(7)) & int32(-8)
			} else {
			}
			v134 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[6]))
			v138 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v116)+25)) = uint8(v138)
			*(*uint8)(unsafe.Add(mBase, uint32(v116)+24)) = uint8(base.B2i32(v134 == int32(0)))
			if v134 != 0 {
				if v50 != 0 {
					v158 = v69
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = int32(8)
					v158 = int32(1)
				}
			} else {
				v144 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[7])))
				if v144&int32(1)|l0&int32(2)|base.B2i32(v50 <= int32(0)) != 0 {
					if v50 != 0 {
						v158 = v69
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = int32(8)
						v158 = int32(1)
					}
				} else {
					v153 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v116)+26)) = uint8(v153)
					v158 = v69
				}
			}
			*(*uint16)(unsafe.Add(mBase, uint32(v116))) = uint16(v158)
			v162 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
			*(*int32)(unsafe.Add(mBase, uint32(v116)+56)) = l8
			*(*uint16)(unsafe.Add(mBase, uint32(v116)+2)) = uint16(v162)
			*(*uint16)(unsafe.Add(mBase, uint32(v116)+8)) = uint16(v104)
			*(*int32)(unsafe.Add(mBase, uint32(v116)+36)) = l7
			*(*int32)(unsafe.Add(mBase, uint32(v116)+32)) = l6
			*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)) = uint16(v106)
			*(*int64)(unsafe.Add(mBase, uint32(v116)+40)) = int64(-1)
			v171 = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v116)+28)) = v171
			v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			*(*uint8)(unsafe.Add(mBase, uint32(v116)+27)) = uint8(base.B2i32(v173 != v171))
			if base.Ui32(v104) < base.Ui32(base.I32_extend16_s(v162)) {
				v179 = v104
			} else {
				v179 = v162
			}
			if l0&int32(4) != 0 {
				v183 = v179
			} else {
				v183 = int32(1)
			}
			*(*uint16)(unsafe.Add(mBase, uint32(v116)+14)) = uint16(v183)
			if int32(0) < v158 {
				v198 = int32(0)
				for {
					v204 = v198 * int32(84)
					v205 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
					*(*int32)(unsafe.Add(mBase, uint32(v204+v205)+4)) = l2
					v208 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
					*(*int32)(unsafe.Add(mBase, uint32(v208+v204)+8)) = l3
					v211 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
					*(*uint8)(unsafe.Add(mBase, uint32(v211+v204)+12)) = uint8(v5)
					v214 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
					*(*int32)(unsafe.Add(mBase, uint32(v214+v204)+16)) = l5
					v217 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
					*(*int32)(unsafe.Add(mBase, uint32(v217+v204)+20)) = l1
					v221 = v198 + int32(1)
					if v221 != v158 {
						v198 = v221
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			return v116
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		if l2 != 0 {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
			if base.Ui32(v22) < base.Ui32(int32(_a_F_read_stream_begin_impl_3)) {
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[1]))
				v50 = v32
				v52 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
				if l1 == int32(0) {
					v56 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[3]))
					v65 = v56
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v59 = base.I32_div_s(v57, int32(2))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v60 != int32(1) {
						v63 = v59
					} else {
						v63 = v57
					}
					v65 = v63
				}
				v66 = int32(_a_F_read_stream_begin_impl_0)
				if v66 <= v50 {
					v69 = v66
				} else {
					v69 = v50
				}
				v72 = v52 * (v69 + int32(1))
				v74 = int32(16)
				v79 = (v52<<(uint(v74)%32) - int32(_a_F_read_stream_begin_impl_1)) >> (uint(v74) % 32)
				v80 = int32(_a_F_read_stream_begin_impl_2) - v79
				if base.Ui32(v72) < base.Ui32(v80) {
					v82 = v72
				} else {
					v82 = v80
				}
				if base.Ui32(v65) < base.Ui32(v82) {
					v84 = v65
				} else {
					v84 = v82
				}
				v85 = int32(1)
				if v69 <= v85 {
					v88 = v85
				} else {
					v88 = v69
				}
				v90 = v88 * int32(84)
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
				if v92 != int32(-1) {
					v96 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[4]))
					v99 = v96
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[5]))
					v99 = v98
				}
				if base.Ui32(v84) < base.Ui32(v99) {
					v101 = v84
				} else {
					v101 = v99
				}
				if base.Ui32(v101) <= base.Ui32(int32(1)) {
					v104 = int32(1)
				} else {
					v104 = v101
				}
				v106 = v104 + int32(1)
				v107 = base.I32_extend16_s(v106)
				v112 = (v107 + v79) << (uint(int32(2)) % 32)
				v116 = F_palloc(m, v90+l8*v107+v112+int32(96))
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return int32(0)
				} else {
					base.MemoryFill(m, v116, int32(0), int32(80))
					v125 = (v116 + v112 + int32(87)) & int32(-8)
					*(*int32)(unsafe.Add(mBase, uint32(v116)+64)) = v125
					if l8 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v116)+60)) = (v90 + v125 + int32(7)) & int32(-8)
					} else {
					}
					v134 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[6]))
					v138 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v116)+25)) = uint8(v138)
					*(*uint8)(unsafe.Add(mBase, uint32(v116)+24)) = uint8(base.B2i32(v134 == int32(0)))
					if v134 != 0 {
						if v50 != 0 {
							v158 = v69
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = int32(8)
							v158 = int32(1)
						}
					} else {
						v144 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[7])))
						if v144&int32(1)|l0&int32(2)|base.B2i32(v50 <= int32(0)) != 0 {
							if v50 != 0 {
								v158 = v69
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = int32(8)
								v158 = int32(1)
							}
						} else {
							v153 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v116)+26)) = uint8(v153)
							v158 = v69
						}
					}
					*(*uint16)(unsafe.Add(mBase, uint32(v116))) = uint16(v158)
					v162 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v116)+56)) = l8
					*(*uint16)(unsafe.Add(mBase, uint32(v116)+2)) = uint16(v162)
					*(*uint16)(unsafe.Add(mBase, uint32(v116)+8)) = uint16(v104)
					*(*int32)(unsafe.Add(mBase, uint32(v116)+36)) = l7
					*(*int32)(unsafe.Add(mBase, uint32(v116)+32)) = l6
					*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)) = uint16(v106)
					*(*int64)(unsafe.Add(mBase, uint32(v116)+40)) = int64(-1)
					v171 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v116)+28)) = v171
					v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					*(*uint8)(unsafe.Add(mBase, uint32(v116)+27)) = uint8(base.B2i32(v173 != v171))
					if base.Ui32(v104) < base.Ui32(base.I32_extend16_s(v162)) {
						v179 = v104
					} else {
						v179 = v162
					}
					if l0&int32(4) != 0 {
						v183 = v179
					} else {
						v183 = int32(1)
					}
					*(*uint16)(unsafe.Add(mBase, uint32(v116)+14)) = uint16(v183)
					if int32(0) < v158 {
						v198 = int32(0)
						for {
							v204 = v198 * int32(84)
							v205 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v204+v205)+4)) = l2
							v208 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v208+v204)+8)) = l3
							v211 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
							*(*uint8)(unsafe.Add(mBase, uint32(v211+v204)+12)) = uint8(v5)
							v214 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v214+v204)+16)) = l5
							v217 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v217+v204)+20)) = l1
							v221 = v198 + int32(1)
							if v221 != v158 {
								v198 = v221
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					return v116
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				if base.B2i32(base.Ui32(v25) < base.Ui32(int32(_a_F_read_stream_begin_impl_3))) == int32(0) {
					if l0&int32(1) != 0 {
						v35 = F_get_tablespace_maintenance_io_concurrency(m, v21)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v50 = v35
							v52 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
							if l1 == int32(0) {
								v56 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[3]))
								v65 = v56
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v59 = base.I32_div_s(v57, int32(2))
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								if v60 != int32(1) {
									v63 = v59
								} else {
									v63 = v57
								}
								v65 = v63
							}
							v66 = int32(_a_F_read_stream_begin_impl_0)
							if v66 <= v50 {
								v69 = v66
							} else {
								v69 = v50
							}
							v72 = v52 * (v69 + int32(1))
							v74 = int32(16)
							v79 = (v52<<(uint(v74)%32) - int32(_a_F_read_stream_begin_impl_1)) >> (uint(v74) % 32)
							v80 = int32(_a_F_read_stream_begin_impl_2) - v79
							if base.Ui32(v72) < base.Ui32(v80) {
								v82 = v72
							} else {
								v82 = v80
							}
							if base.Ui32(v65) < base.Ui32(v82) {
								v84 = v65
							} else {
								v84 = v82
							}
							v85 = int32(1)
							if v69 <= v85 {
								v88 = v85
							} else {
								v88 = v69
							}
							v90 = v88 * int32(84)
							v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
							if v92 != int32(-1) {
								v96 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[4]))
								v99 = v96
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[5]))
								v99 = v98
							}
							if base.Ui32(v84) < base.Ui32(v99) {
								v101 = v84
							} else {
								v101 = v99
							}
							if base.Ui32(v101) <= base.Ui32(int32(1)) {
								v104 = int32(1)
							} else {
								v104 = v101
							}
							v106 = v104 + int32(1)
							v107 = base.I32_extend16_s(v106)
							v112 = (v107 + v79) << (uint(int32(2)) % 32)
							v116 = F_palloc(m, v90+l8*v107+v112+int32(96))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int32(0)
							} else {
								base.MemoryFill(m, v116, int32(0), int32(80))
								v125 = (v116 + v112 + int32(87)) & int32(-8)
								*(*int32)(unsafe.Add(mBase, uint32(v116)+64)) = v125
								if l8 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v116)+60)) = (v90 + v125 + int32(7)) & int32(-8)
								} else {
								}
								v134 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[6]))
								v138 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v116)+25)) = uint8(v138)
								*(*uint8)(unsafe.Add(mBase, uint32(v116)+24)) = uint8(base.B2i32(v134 == int32(0)))
								if v134 != 0 {
									if v50 != 0 {
										v158 = v69
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = int32(8)
										v158 = int32(1)
									}
								} else {
									v144 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[7])))
									if v144&int32(1)|l0&int32(2)|base.B2i32(v50 <= int32(0)) != 0 {
										if v50 != 0 {
											v158 = v69
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = int32(8)
											v158 = int32(1)
										}
									} else {
										v153 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v116)+26)) = uint8(v153)
										v158 = v69
									}
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v116))) = uint16(v158)
								v162 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v116)+56)) = l8
								*(*uint16)(unsafe.Add(mBase, uint32(v116)+2)) = uint16(v162)
								*(*uint16)(unsafe.Add(mBase, uint32(v116)+8)) = uint16(v104)
								*(*int32)(unsafe.Add(mBase, uint32(v116)+36)) = l7
								*(*int32)(unsafe.Add(mBase, uint32(v116)+32)) = l6
								*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)) = uint16(v106)
								*(*int64)(unsafe.Add(mBase, uint32(v116)+40)) = int64(-1)
								v171 = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v116)+28)) = v171
								v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
								*(*uint8)(unsafe.Add(mBase, uint32(v116)+27)) = uint8(base.B2i32(v173 != v171))
								if base.Ui32(v104) < base.Ui32(base.I32_extend16_s(v162)) {
									v179 = v104
								} else {
									v179 = v162
								}
								if l0&int32(4) != 0 {
									v183 = v179
								} else {
									v183 = int32(1)
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v116)+14)) = uint16(v183)
								if int32(0) < v158 {
									v198 = int32(0)
									for {
										v204 = v198 * int32(84)
										v205 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v204+v205)+4)) = l2
										v208 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v208+v204)+8)) = l3
										v211 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
										*(*uint8)(unsafe.Add(mBase, uint32(v211+v204)+12)) = uint8(v5)
										v214 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v214+v204)+16)) = l5
										v217 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v217+v204)+20)) = l1
										v221 = v198 + int32(1)
										if v221 != v158 {
											v198 = v221
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								return v116
							}
						}
					} else {
						v39 = F_get_tablespace(m, v21)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							if v41 != 0 {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
								if int32(0) <= v42 {
									v48 = v42
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[1]))
									v48 = v47
								}
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[1]))
								v48 = v47
							}
							v50 = v48
							v52 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
							if l1 == int32(0) {
								v56 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[3]))
								v65 = v56
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v59 = base.I32_div_s(v57, int32(2))
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								if v60 != int32(1) {
									v63 = v59
								} else {
									v63 = v57
								}
								v65 = v63
							}
							v66 = int32(_a_F_read_stream_begin_impl_0)
							if v66 <= v50 {
								v69 = v66
							} else {
								v69 = v50
							}
							v72 = v52 * (v69 + int32(1))
							v74 = int32(16)
							v79 = (v52<<(uint(v74)%32) - int32(_a_F_read_stream_begin_impl_1)) >> (uint(v74) % 32)
							v80 = int32(_a_F_read_stream_begin_impl_2) - v79
							if base.Ui32(v72) < base.Ui32(v80) {
								v82 = v72
							} else {
								v82 = v80
							}
							if base.Ui32(v65) < base.Ui32(v82) {
								v84 = v65
							} else {
								v84 = v82
							}
							v85 = int32(1)
							if v69 <= v85 {
								v88 = v85
							} else {
								v88 = v69
							}
							v90 = v88 * int32(84)
							v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
							if v92 != int32(-1) {
								v96 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[4]))
								v99 = v96
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[5]))
								v99 = v98
							}
							if base.Ui32(v84) < base.Ui32(v99) {
								v101 = v84
							} else {
								v101 = v99
							}
							if base.Ui32(v101) <= base.Ui32(int32(1)) {
								v104 = int32(1)
							} else {
								v104 = v101
							}
							v106 = v104 + int32(1)
							v107 = base.I32_extend16_s(v106)
							v112 = (v107 + v79) << (uint(int32(2)) % 32)
							v116 = F_palloc(m, v90+l8*v107+v112+int32(96))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int32(0)
							} else {
								base.MemoryFill(m, v116, int32(0), int32(80))
								v125 = (v116 + v112 + int32(87)) & int32(-8)
								*(*int32)(unsafe.Add(mBase, uint32(v116)+64)) = v125
								if l8 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v116)+60)) = (v90 + v125 + int32(7)) & int32(-8)
								} else {
								}
								v134 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[6]))
								v138 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v116)+25)) = uint8(v138)
								*(*uint8)(unsafe.Add(mBase, uint32(v116)+24)) = uint8(base.B2i32(v134 == int32(0)))
								if v134 != 0 {
									if v50 != 0 {
										v158 = v69
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = int32(8)
										v158 = int32(1)
									}
								} else {
									v144 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[7])))
									if v144&int32(1)|l0&int32(2)|base.B2i32(v50 <= int32(0)) != 0 {
										if v50 != 0 {
											v158 = v69
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = int32(8)
											v158 = int32(1)
										}
									} else {
										v153 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v116)+26)) = uint8(v153)
										v158 = v69
									}
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v116))) = uint16(v158)
								v162 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v116)+56)) = l8
								*(*uint16)(unsafe.Add(mBase, uint32(v116)+2)) = uint16(v162)
								*(*uint16)(unsafe.Add(mBase, uint32(v116)+8)) = uint16(v104)
								*(*int32)(unsafe.Add(mBase, uint32(v116)+36)) = l7
								*(*int32)(unsafe.Add(mBase, uint32(v116)+32)) = l6
								*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)) = uint16(v106)
								*(*int64)(unsafe.Add(mBase, uint32(v116)+40)) = int64(-1)
								v171 = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v116)+28)) = v171
								v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
								*(*uint8)(unsafe.Add(mBase, uint32(v116)+27)) = uint8(base.B2i32(v173 != v171))
								if base.Ui32(v104) < base.Ui32(base.I32_extend16_s(v162)) {
									v179 = v104
								} else {
									v179 = v162
								}
								if l0&int32(4) != 0 {
									v183 = v179
								} else {
									v183 = int32(1)
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v116)+14)) = uint16(v183)
								if int32(0) < v158 {
									v198 = int32(0)
									for {
										v204 = v198 * int32(84)
										v205 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v204+v205)+4)) = l2
										v208 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v208+v204)+8)) = l3
										v211 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
										*(*uint8)(unsafe.Add(mBase, uint32(v211+v204)+12)) = uint8(v5)
										v214 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v214+v204)+16)) = l5
										v217 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v217+v204)+20)) = l1
										v221 = v198 + int32(1)
										if v221 != v158 {
											v198 = v221
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								return v116
							}
						}
					}
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[1]))
					v50 = v32
					v52 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
					if l1 == int32(0) {
						v56 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[3]))
						v65 = v56
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v59 = base.I32_div_s(v57, int32(2))
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if v60 != int32(1) {
							v63 = v59
						} else {
							v63 = v57
						}
						v65 = v63
					}
					v66 = int32(_a_F_read_stream_begin_impl_0)
					if v66 <= v50 {
						v69 = v66
					} else {
						v69 = v50
					}
					v72 = v52 * (v69 + int32(1))
					v74 = int32(16)
					v79 = (v52<<(uint(v74)%32) - int32(_a_F_read_stream_begin_impl_1)) >> (uint(v74) % 32)
					v80 = int32(_a_F_read_stream_begin_impl_2) - v79
					if base.Ui32(v72) < base.Ui32(v80) {
						v82 = v72
					} else {
						v82 = v80
					}
					if base.Ui32(v65) < base.Ui32(v82) {
						v84 = v65
					} else {
						v84 = v82
					}
					v85 = int32(1)
					if v69 <= v85 {
						v88 = v85
					} else {
						v88 = v69
					}
					v90 = v88 * int32(84)
					v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					if v92 != int32(-1) {
						v96 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[4]))
						v99 = v96
					} else {
						v98 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[5]))
						v99 = v98
					}
					if base.Ui32(v84) < base.Ui32(v99) {
						v101 = v84
					} else {
						v101 = v99
					}
					if base.Ui32(v101) <= base.Ui32(int32(1)) {
						v104 = int32(1)
					} else {
						v104 = v101
					}
					v106 = v104 + int32(1)
					v107 = base.I32_extend16_s(v106)
					v112 = (v107 + v79) << (uint(int32(2)) % 32)
					v116 = F_palloc(m, v90+l8*v107+v112+int32(96))
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return int32(0)
					} else {
						base.MemoryFill(m, v116, int32(0), int32(80))
						v125 = (v116 + v112 + int32(87)) & int32(-8)
						*(*int32)(unsafe.Add(mBase, uint32(v116)+64)) = v125
						if l8 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v116)+60)) = (v90 + v125 + int32(7)) & int32(-8)
						} else {
						}
						v134 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[6]))
						v138 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v116)+25)) = uint8(v138)
						*(*uint8)(unsafe.Add(mBase, uint32(v116)+24)) = uint8(base.B2i32(v134 == int32(0)))
						if v134 != 0 {
							if v50 != 0 {
								v158 = v69
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = int32(8)
								v158 = int32(1)
							}
						} else {
							v144 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[7])))
							if v144&int32(1)|l0&int32(2)|base.B2i32(v50 <= int32(0)) != 0 {
								if v50 != 0 {
									v158 = v69
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = int32(8)
									v158 = int32(1)
								}
							} else {
								v153 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v116)+26)) = uint8(v153)
								v158 = v69
							}
						}
						*(*uint16)(unsafe.Add(mBase, uint32(v116))) = uint16(v158)
						v162 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v116)+56)) = l8
						*(*uint16)(unsafe.Add(mBase, uint32(v116)+2)) = uint16(v162)
						*(*uint16)(unsafe.Add(mBase, uint32(v116)+8)) = uint16(v104)
						*(*int32)(unsafe.Add(mBase, uint32(v116)+36)) = l7
						*(*int32)(unsafe.Add(mBase, uint32(v116)+32)) = l6
						*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)) = uint16(v106)
						*(*int64)(unsafe.Add(mBase, uint32(v116)+40)) = int64(-1)
						v171 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v116)+28)) = v171
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						*(*uint8)(unsafe.Add(mBase, uint32(v116)+27)) = uint8(base.B2i32(v173 != v171))
						if base.Ui32(v104) < base.Ui32(base.I32_extend16_s(v162)) {
							v179 = v104
						} else {
							v179 = v162
						}
						if l0&int32(4) != 0 {
							v183 = v179
						} else {
							v183 = int32(1)
						}
						*(*uint16)(unsafe.Add(mBase, uint32(v116)+14)) = uint16(v183)
						if int32(0) < v158 {
							v198 = int32(0)
							for {
								v204 = v198 * int32(84)
								v205 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
								*(*int32)(unsafe.Add(mBase, uint32(v204+v205)+4)) = l2
								v208 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
								*(*int32)(unsafe.Add(mBase, uint32(v208+v204)+8)) = l3
								v211 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
								*(*uint8)(unsafe.Add(mBase, uint32(v211+v204)+12)) = uint8(v5)
								v214 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
								*(*int32)(unsafe.Add(mBase, uint32(v214+v204)+16)) = l5
								v217 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
								*(*int32)(unsafe.Add(mBase, uint32(v217+v204)+20)) = l1
								v221 = v198 + int32(1)
								if v221 != v158 {
									v198 = v221
									continue
								} else {
									break
								}
								break
							}
						} else {
						}
						return v116
					}
				}
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			if base.B2i32(base.Ui32(v25) < base.Ui32(int32(_a_F_read_stream_begin_impl_3))) == int32(0) {
				if l0&int32(1) != 0 {
					v35 = F_get_tablespace_maintenance_io_concurrency(m, v21)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v50 = v35
						v52 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
						if l1 == int32(0) {
							v56 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[3]))
							v65 = v56
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v59 = base.I32_div_s(v57, int32(2))
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if v60 != int32(1) {
								v63 = v59
							} else {
								v63 = v57
							}
							v65 = v63
						}
						v66 = int32(_a_F_read_stream_begin_impl_0)
						if v66 <= v50 {
							v69 = v66
						} else {
							v69 = v50
						}
						v72 = v52 * (v69 + int32(1))
						v74 = int32(16)
						v79 = (v52<<(uint(v74)%32) - int32(_a_F_read_stream_begin_impl_1)) >> (uint(v74) % 32)
						v80 = int32(_a_F_read_stream_begin_impl_2) - v79
						if base.Ui32(v72) < base.Ui32(v80) {
							v82 = v72
						} else {
							v82 = v80
						}
						if base.Ui32(v65) < base.Ui32(v82) {
							v84 = v65
						} else {
							v84 = v82
						}
						v85 = int32(1)
						if v69 <= v85 {
							v88 = v85
						} else {
							v88 = v69
						}
						v90 = v88 * int32(84)
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						if v92 != int32(-1) {
							v96 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[4]))
							v99 = v96
						} else {
							v98 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[5]))
							v99 = v98
						}
						if base.Ui32(v84) < base.Ui32(v99) {
							v101 = v84
						} else {
							v101 = v99
						}
						if base.Ui32(v101) <= base.Ui32(int32(1)) {
							v104 = int32(1)
						} else {
							v104 = v101
						}
						v106 = v104 + int32(1)
						v107 = base.I32_extend16_s(v106)
						v112 = (v107 + v79) << (uint(int32(2)) % 32)
						v116 = F_palloc(m, v90+l8*v107+v112+int32(96))
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return int32(0)
						} else {
							base.MemoryFill(m, v116, int32(0), int32(80))
							v125 = (v116 + v112 + int32(87)) & int32(-8)
							*(*int32)(unsafe.Add(mBase, uint32(v116)+64)) = v125
							if l8 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v116)+60)) = (v90 + v125 + int32(7)) & int32(-8)
							} else {
							}
							v134 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[6]))
							v138 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v116)+25)) = uint8(v138)
							*(*uint8)(unsafe.Add(mBase, uint32(v116)+24)) = uint8(base.B2i32(v134 == int32(0)))
							if v134 != 0 {
								if v50 != 0 {
									v158 = v69
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = int32(8)
									v158 = int32(1)
								}
							} else {
								v144 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[7])))
								if v144&int32(1)|l0&int32(2)|base.B2i32(v50 <= int32(0)) != 0 {
									if v50 != 0 {
										v158 = v69
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = int32(8)
										v158 = int32(1)
									}
								} else {
									v153 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v116)+26)) = uint8(v153)
									v158 = v69
								}
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v116))) = uint16(v158)
							v162 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v116)+56)) = l8
							*(*uint16)(unsafe.Add(mBase, uint32(v116)+2)) = uint16(v162)
							*(*uint16)(unsafe.Add(mBase, uint32(v116)+8)) = uint16(v104)
							*(*int32)(unsafe.Add(mBase, uint32(v116)+36)) = l7
							*(*int32)(unsafe.Add(mBase, uint32(v116)+32)) = l6
							*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)) = uint16(v106)
							*(*int64)(unsafe.Add(mBase, uint32(v116)+40)) = int64(-1)
							v171 = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v116)+28)) = v171
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
							*(*uint8)(unsafe.Add(mBase, uint32(v116)+27)) = uint8(base.B2i32(v173 != v171))
							if base.Ui32(v104) < base.Ui32(base.I32_extend16_s(v162)) {
								v179 = v104
							} else {
								v179 = v162
							}
							if l0&int32(4) != 0 {
								v183 = v179
							} else {
								v183 = int32(1)
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v116)+14)) = uint16(v183)
							if int32(0) < v158 {
								v198 = int32(0)
								for {
									v204 = v198 * int32(84)
									v205 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v204+v205)+4)) = l2
									v208 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v208+v204)+8)) = l3
									v211 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
									*(*uint8)(unsafe.Add(mBase, uint32(v211+v204)+12)) = uint8(v5)
									v214 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v214+v204)+16)) = l5
									v217 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v217+v204)+20)) = l1
									v221 = v198 + int32(1)
									if v221 != v158 {
										v198 = v221
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							return v116
						}
					}
				} else {
					v39 = F_get_tablespace(m, v21)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
						if v41 != 0 {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
							if int32(0) <= v42 {
								v48 = v42
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[1]))
								v48 = v47
							}
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[1]))
							v48 = v47
						}
						v50 = v48
						v52 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
						if l1 == int32(0) {
							v56 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[3]))
							v65 = v56
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v59 = base.I32_div_s(v57, int32(2))
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if v60 != int32(1) {
								v63 = v59
							} else {
								v63 = v57
							}
							v65 = v63
						}
						v66 = int32(_a_F_read_stream_begin_impl_0)
						if v66 <= v50 {
							v69 = v66
						} else {
							v69 = v50
						}
						v72 = v52 * (v69 + int32(1))
						v74 = int32(16)
						v79 = (v52<<(uint(v74)%32) - int32(_a_F_read_stream_begin_impl_1)) >> (uint(v74) % 32)
						v80 = int32(_a_F_read_stream_begin_impl_2) - v79
						if base.Ui32(v72) < base.Ui32(v80) {
							v82 = v72
						} else {
							v82 = v80
						}
						if base.Ui32(v65) < base.Ui32(v82) {
							v84 = v65
						} else {
							v84 = v82
						}
						v85 = int32(1)
						if v69 <= v85 {
							v88 = v85
						} else {
							v88 = v69
						}
						v90 = v88 * int32(84)
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						if v92 != int32(-1) {
							v96 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[4]))
							v99 = v96
						} else {
							v98 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[5]))
							v99 = v98
						}
						if base.Ui32(v84) < base.Ui32(v99) {
							v101 = v84
						} else {
							v101 = v99
						}
						if base.Ui32(v101) <= base.Ui32(int32(1)) {
							v104 = int32(1)
						} else {
							v104 = v101
						}
						v106 = v104 + int32(1)
						v107 = base.I32_extend16_s(v106)
						v112 = (v107 + v79) << (uint(int32(2)) % 32)
						v116 = F_palloc(m, v90+l8*v107+v112+int32(96))
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return int32(0)
						} else {
							base.MemoryFill(m, v116, int32(0), int32(80))
							v125 = (v116 + v112 + int32(87)) & int32(-8)
							*(*int32)(unsafe.Add(mBase, uint32(v116)+64)) = v125
							if l8 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v116)+60)) = (v90 + v125 + int32(7)) & int32(-8)
							} else {
							}
							v134 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[6]))
							v138 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v116)+25)) = uint8(v138)
							*(*uint8)(unsafe.Add(mBase, uint32(v116)+24)) = uint8(base.B2i32(v134 == int32(0)))
							if v134 != 0 {
								if v50 != 0 {
									v158 = v69
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = int32(8)
									v158 = int32(1)
								}
							} else {
								v144 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[7])))
								if v144&int32(1)|l0&int32(2)|base.B2i32(v50 <= int32(0)) != 0 {
									if v50 != 0 {
										v158 = v69
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = int32(8)
										v158 = int32(1)
									}
								} else {
									v153 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v116)+26)) = uint8(v153)
									v158 = v69
								}
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v116))) = uint16(v158)
							v162 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v116)+56)) = l8
							*(*uint16)(unsafe.Add(mBase, uint32(v116)+2)) = uint16(v162)
							*(*uint16)(unsafe.Add(mBase, uint32(v116)+8)) = uint16(v104)
							*(*int32)(unsafe.Add(mBase, uint32(v116)+36)) = l7
							*(*int32)(unsafe.Add(mBase, uint32(v116)+32)) = l6
							*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)) = uint16(v106)
							*(*int64)(unsafe.Add(mBase, uint32(v116)+40)) = int64(-1)
							v171 = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v116)+28)) = v171
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
							*(*uint8)(unsafe.Add(mBase, uint32(v116)+27)) = uint8(base.B2i32(v173 != v171))
							if base.Ui32(v104) < base.Ui32(base.I32_extend16_s(v162)) {
								v179 = v104
							} else {
								v179 = v162
							}
							if l0&int32(4) != 0 {
								v183 = v179
							} else {
								v183 = int32(1)
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v116)+14)) = uint16(v183)
							if int32(0) < v158 {
								v198 = int32(0)
								for {
									v204 = v198 * int32(84)
									v205 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v204+v205)+4)) = l2
									v208 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v208+v204)+8)) = l3
									v211 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
									*(*uint8)(unsafe.Add(mBase, uint32(v211+v204)+12)) = uint8(v5)
									v214 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v214+v204)+16)) = l5
									v217 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v217+v204)+20)) = l1
									v221 = v198 + int32(1)
									if v221 != v158 {
										v198 = v221
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							return v116
						}
					}
				}
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[1]))
				v50 = v32
				v52 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
				if l1 == int32(0) {
					v56 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[3]))
					v65 = v56
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v59 = base.I32_div_s(v57, int32(2))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v60 != int32(1) {
						v63 = v59
					} else {
						v63 = v57
					}
					v65 = v63
				}
				v66 = int32(_a_F_read_stream_begin_impl_0)
				if v66 <= v50 {
					v69 = v66
				} else {
					v69 = v50
				}
				v72 = v52 * (v69 + int32(1))
				v74 = int32(16)
				v79 = (v52<<(uint(v74)%32) - int32(_a_F_read_stream_begin_impl_1)) >> (uint(v74) % 32)
				v80 = int32(_a_F_read_stream_begin_impl_2) - v79
				if base.Ui32(v72) < base.Ui32(v80) {
					v82 = v72
				} else {
					v82 = v80
				}
				if base.Ui32(v65) < base.Ui32(v82) {
					v84 = v65
				} else {
					v84 = v82
				}
				v85 = int32(1)
				if v69 <= v85 {
					v88 = v85
				} else {
					v88 = v69
				}
				v90 = v88 * int32(84)
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
				if v92 != int32(-1) {
					v96 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[4]))
					v99 = v96
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[5]))
					v99 = v98
				}
				if base.Ui32(v84) < base.Ui32(v99) {
					v101 = v84
				} else {
					v101 = v99
				}
				if base.Ui32(v101) <= base.Ui32(int32(1)) {
					v104 = int32(1)
				} else {
					v104 = v101
				}
				v106 = v104 + int32(1)
				v107 = base.I32_extend16_s(v106)
				v112 = (v107 + v79) << (uint(int32(2)) % 32)
				v116 = F_palloc(m, v90+l8*v107+v112+int32(96))
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return int32(0)
				} else {
					base.MemoryFill(m, v116, int32(0), int32(80))
					v125 = (v116 + v112 + int32(87)) & int32(-8)
					*(*int32)(unsafe.Add(mBase, uint32(v116)+64)) = v125
					if l8 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v116)+60)) = (v90 + v125 + int32(7)) & int32(-8)
					} else {
					}
					v134 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[6]))
					v138 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v116)+25)) = uint8(v138)
					*(*uint8)(unsafe.Add(mBase, uint32(v116)+24)) = uint8(base.B2i32(v134 == int32(0)))
					if v134 != 0 {
						if v50 != 0 {
							v158 = v69
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = int32(8)
							v158 = int32(1)
						}
					} else {
						v144 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[7])))
						if v144&int32(1)|l0&int32(2)|base.B2i32(v50 <= int32(0)) != 0 {
							if v50 != 0 {
								v158 = v69
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = int32(8)
								v158 = int32(1)
							}
						} else {
							v153 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v116)+26)) = uint8(v153)
							v158 = v69
						}
					}
					*(*uint16)(unsafe.Add(mBase, uint32(v116))) = uint16(v158)
					v162 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_begin_impl[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v116)+56)) = l8
					*(*uint16)(unsafe.Add(mBase, uint32(v116)+2)) = uint16(v162)
					*(*uint16)(unsafe.Add(mBase, uint32(v116)+8)) = uint16(v104)
					*(*int32)(unsafe.Add(mBase, uint32(v116)+36)) = l7
					*(*int32)(unsafe.Add(mBase, uint32(v116)+32)) = l6
					*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)) = uint16(v106)
					*(*int64)(unsafe.Add(mBase, uint32(v116)+40)) = int64(-1)
					v171 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v116)+28)) = v171
					v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					*(*uint8)(unsafe.Add(mBase, uint32(v116)+27)) = uint8(base.B2i32(v173 != v171))
					if base.Ui32(v104) < base.Ui32(base.I32_extend16_s(v162)) {
						v179 = v104
					} else {
						v179 = v162
					}
					if l0&int32(4) != 0 {
						v183 = v179
					} else {
						v183 = int32(1)
					}
					*(*uint16)(unsafe.Add(mBase, uint32(v116)+14)) = uint16(v183)
					if int32(0) < v158 {
						v198 = int32(0)
						for {
							v204 = v198 * int32(84)
							v205 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v204+v205)+4)) = l2
							v208 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v208+v204)+8)) = l3
							v211 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
							*(*uint8)(unsafe.Add(mBase, uint32(v211+v204)+12)) = uint8(v5)
							v214 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v214+v204)+16)) = l5
							v217 = *(*int32)(unsafe.Add(mBase, uint32(v116)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v217+v204)+20)) = l1
							v221 = v198 + int32(1)
							if v221 != v158 {
								v198 = v221
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					return v116
				}
			}
		}
	}
}
